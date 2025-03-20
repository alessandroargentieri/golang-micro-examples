package main

import (
	"bytes"
	"context"
	"encoding/base64"
	"fmt"
	"log"
	"net/url"
	"os"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	smithyendpoints "github.com/aws/smithy-go/endpoints"
)

// CivoS3EndpointResolverV2 is the implementation of the aws-sdk-go-v2 EndpointResolverV2 interface to customise the S3 Presigned URL
// matching the Civo pattern: https://objectstore.<region>.civo.com/<bucket-name>/...
// instead of the default AWS one: https://<bucket-name>.objectstore.<region>.civo.com/...
type CivoS3EndpointResolverV2 struct {
	Host string
}

// ResolveEndpoint is the implementation of the ResolveEndpoint function of the aws-sdk-go-v2 EndpointResolverV2 interface
func (r *CivoS3EndpointResolverV2) ResolveEndpoint(ctx context.Context, params s3.EndpointParameters) (smithyendpoints.Endpoint, error) {

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

func main() {
	// Get environment variables
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
	base64SVG := os.Getenv("BASE64_SVG")
	if base64SVG == "" {
		log.Fatalf("BASE64_SVG env var missing")
	}

	// Example base64 SVG
	// In reality, this would come from your application
	//base64SVG := "PHN2ZyB4bWxucz0iaHR0cDovL3d3dy53My5vcmcvMjAwMC9zdmciIHdpZHRoPSIxMDAiIGhlaWdodD0iMTAwIj48Y2lyY2xlIGN4PSI1MCIgY3k9IjUwIiByPSI0MCIgc3Ryb2tlPSJibGFjayIgc3Ryb2tlLXdpZHRoPSIzIiBmaWxsPSJyZWQiIC8+PC9zdmc+"

	// Create a context
	ctx := context.Background()

	/////
	cfg, err := config.LoadDefaultConfig(ctx,
		config.WithRegion("lon1"),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(
			accessKeyID,
			secretAccessKey,
			"",
		)),
	)
	if err != nil {
		log.Fatalf("failed to load s3 bucket config %s", err)
	}

	// Create S3 client with UsePathStyle option and custom endpoint resolver to match Civo Endpoint Style (for presigned URLs only)
	client, err := s3.NewFromConfig(cfg, func(o *s3.Options) {
		o.UsePathStyle = true
	}, s3.WithEndpointResolverV2(&CivoS3EndpointResolverV2{
		Host: "https://objectstore.lon1.civo.com",
	})), nil

	/////

	// Step 1: Remove any data URL prefix if it exists
	// For example, "data:image/svg+xml;base64," should be removed
	cleanBase64 := base64SVG
	if strings.Contains(cleanBase64, ",") {
		parts := strings.Split(cleanBase64, ",")
		cleanBase64 = parts[len(parts)-1]
	}

	// Step 2: Decode the base64 string
	decodedSVG, err := base64.StdEncoding.DecodeString(cleanBase64)
	if err != nil {
		log.Fatalf("Failed to decode base64 string: %v", err)
	}
	_ = decodedSVG

	// Step 3: Create a reader from the decoded data
	//reader := bytes.NewReader(decodedSVG)

	// Step 4: Upload to S3
	_, err = client.PutObject(ctx, &s3.PutObjectInput{
		Bucket:          aws.String(bucketName),
		Key:             aws.String(bucketFilePath),
		Body:            strings.NewReader(base64SVG),
		ContentType:     aws.String("image/svg+xml"),
		ContentEncoding: aws.String("base64"),
	})
	if err != nil {
		log.Fatalf("Failed to upload file to S3: %v", err)
	}

	fmt.Printf("Successfully uploaded SVG to s3://%s/%s\n", bucketName, bucketFilePath)

	// Generate the public URL (if bucket has public access)
	publicURL := fmt.Sprintf("https://%s/%s/%s", host, bucketName, bucketFilePath)
	fmt.Printf("Public URL: %s\n", publicURL)
}

// UploadBase64SVG is a reusable function to upload a base64-encoded SVG to S3
func UploadBase64SVG(ctx context.Context, client *s3.Client, bucketName, objectKey, base64SVG string) error {
	// Step 1: Remove any data URL prefix if it exists
	cleanBase64 := base64SVG
	if strings.Contains(cleanBase64, ",") {
		parts := strings.Split(cleanBase64, ",")
		cleanBase64 = parts[len(parts)-1]
	}

	// Step 2: Decode the base64 string
	decodedSVG, err := base64.StdEncoding.DecodeString(cleanBase64)
	if err != nil {
		return fmt.Errorf("failed to decode base64 string: %w", err)
	}

	// Step 3: Create a reader from the decoded data
	reader := bytes.NewReader(decodedSVG)

	// Step 4: Upload to S3
	_, err = client.PutObject(ctx, &s3.PutObjectInput{
		Bucket:      aws.String(bucketName),
		Key:         aws.String(objectKey),
		Body:        reader,
		ContentType: aws.String("image/svg+xml"),
	})
	if err != nil {
		return fmt.Errorf("failed to upload file to S3: %w", err)
	}

	return nil
}
