package main

import (
	"context"
	"fmt"
	"log"

	//"net/http"
	"net/url"
	"os"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	smithyendpoints "github.com/aws/smithy-go/endpoints"
)

type CivoEndpointResolverV2 struct {
	Host string
}

func (r *CivoEndpointResolverV2) ResolveEndpoint(ctx context.Context, params s3.EndpointParameters) (smithyendpoints.Endpoint, error) {

	params.Prefix = aws.String("")
	params.ForcePathStyle = aws.Bool(true)
	params.Endpoint = aws.String(r.Host)

	desiredURL := fmt.Sprintf("%s/%s/%s", *params.Endpoint, *params.Bucket, *params.Key)

	finalURL, err := url.Parse(desiredURL)
	if err != nil {
		return smithyendpoints.Endpoint{}, fmt.Errorf("failed to parse target URL: %w", err)
	}

	return smithyendpoints.Endpoint{
		URI: *finalURL,
	}, nil
}

// GeneratePresignedUploadURL creates a presigned URL for uploading a file to Civo S3 with the URL format: https://objectstore.<region>.civo.com/<bucketname>
func GeneratePresignedUploadURL(ctx context.Context, bucketName, objectKey, host, accessKeyID, secretAccessKey string, expireInMinutes int) (string, error) {

	// Create AWS config with custom credentials and endpoint resolver
	cfg, err := config.LoadDefaultConfig(ctx,
		config.WithRegion("lon1"),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(
			accessKeyID,
			secretAccessKey,
			"",
		)),
	)
	if err != nil {
		return "", fmt.Errorf("failed to load AWS config: %w", err)
	}

	// Create S3 client with UsePathStyle option and custom endpoint resolver to match Civo Endpoint Style
	client := s3.NewFromConfig(cfg, func(o *s3.Options) {
		o.UsePathStyle = true
	}, s3.WithEndpointResolverV2(&CivoEndpointResolverV2{
		Host: "https://objectstore.lon1.civo.com",
	}))

	// Create presigner
	presigner := s3.NewPresignClient(client)

	// Generate presigned request for PutObject
	presignedReq, err := presigner.PresignPutObject(ctx, &s3.PutObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(objectKey),
	}, func(opts *s3.PresignOptions) {
		opts.Expires = time.Duration(expireInMinutes) * time.Minute
	})
	if err != nil {
		return "", fmt.Errorf("failed to generate presigned URL: %w", err)
	}

	return presignedReq.URL, nil
}

func main() {
	// Get required environment variables
	host := os.Getenv("AWS_HOST")
	if host == "" {
		log.Fatalf("AWS_HOST env var missing")
	}
	bucketName := os.Getenv("BUCKET_NAME")
	if bucketName == "" {
		log.Fatalf("BUCKET_NAME env var missing")
	}
	accessKeyID := os.Getenv("AWS_ACCESS_KEY_ID")
	if accessKeyID == "" {
		log.Fatalf("AWS_ACCESS_KEY_ID env var missing")
	}
	secretAccessKey := os.Getenv("AWS_SECRET_ACCESS_KEY")
	if secretAccessKey == "" {
		log.Fatalf("AWS_SECRET_ACCESS_KEY env var missing")
	}
	bucketFilePath := os.Getenv("BUCKET_FILE_PATH")
	if bucketFilePath == "" {
		log.Fatalf("BUCKET_FILE_PATH env var missing")
	}

	// Generate a presigned URL valid for 60 minutes
	ctx := context.Background()
	presignedURL, err := GeneratePresignedUploadURL(ctx, bucketName, bucketFilePath, host, accessKeyID, secretAccessKey, 60)
	if err != nil {
		log.Fatalf("Failed to generate presigned URL: %v", err)
	}

	fmt.Println("Presigned Upload URL:")
	fmt.Println(presignedURL)
}
