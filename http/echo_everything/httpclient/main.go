package main

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

func main() {

	type BookRequest struct {
		Title     string `json:"title,omitempty"`
		Author    string `json:"author,omitempty"`
		Publisher string `json:"-"`
	}
	type BookResponse struct {
		Title  string `json:"title,omitempty"`
		Author string `json:"author,omitempty"`
	}

	reqBody := BookRequest{Title: "Relic", Author: "Lincoln & Child", Publisher: "Grand Central"}

	resp, err := NewPostRequest[BookRequest, BookResponse]("http://localhost:8585/hello/path?confirm=yes", reqBody).
		WithBasic("user", "pA$$word").
		WithHeaders(map[string]string{
			"greets": "hello",
			"color":  "blue",
		}).Execute()
	if resp != nil {
		fmt.Printf("Response: %+v\n", *resp)
	} else {
		fmt.Printf("Error %s\n", err)
	}
}

type HTTPClient[T any, V any] struct {
	Url     string
	Verb    string
	Headers map[string]string
	ReqBody *T
}

func NewGetRequest[V any](url string) *HTTPClient[interface{}, V] {
	return &HTTPClient[interface{}, V]{Url: url, Verb: "GET"}
}

func NewDeleteRequest[V any](url string) *HTTPClient[interface{}, V] {
	return &HTTPClient[interface{}, V]{Url: url, Verb: "DELETE"}
}

func NewPostRequest[T any, V any](url string, reqBody T) *HTTPClient[T, V] {
	return &HTTPClient[T, V]{Url: url, Verb: "POST", ReqBody: &reqBody}
}

func NewPatchRequest[T any, V any](url string, reqBody T) *HTTPClient[T, V] {
	return &HTTPClient[T, V]{Url: url, Verb: "PATCH", ReqBody: &reqBody}
}

func NewPutRequest[T any, V any](url string, reqBody T) *HTTPClient[T, V] {
	return &HTTPClient[T, V]{Url: url, Verb: "PUT", ReqBody: &reqBody}
}

func (client *HTTPClient[T, V]) WithBearer(token string) *HTTPClient[T, V] {
	if client.Headers == nil {
		client.Headers = make(map[string]string)
	}
	client.Headers["Authorization"] = fmt.Sprintf("Bearer %s", token)
	return client
}

func (client *HTTPClient[T, V]) WithHeaders(headers map[string]string) *HTTPClient[T, V] {
	if client.Headers == nil {
		client.Headers = make(map[string]string)
	}
	for k, v := range headers {
		client.Headers[k] = v
	}
	return client
}

func (client *HTTPClient[T, V]) WithBasic(user, password string) *HTTPClient[T, V] {
	if client.Headers == nil {
		client.Headers = make(map[string]string)
	}
	client.Headers["Authorization"] = fmt.Sprintf("Basic %s", base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s", user,
		password))))
	return client
}

func (client *HTTPClient[T, V]) Execute() (*V, error) {
	return RestCall[T, V](client.Verb, client.Url, client.ReqBody, client.Headers)
}

func RestCall[T any, V any](verb, url string, reqBody *T, headers map[string]string) (*V, error) {
	var ioReader io.Reader
	if reqBody != nil {
		// Marshal the request body to JSON
		reqBodyBytes, err := json.Marshal(reqBody)
		if err != nil {
			return nil, err
		}
		ioReader = bytes.NewBuffer(reqBodyBytes)

	}

	// Create an HTTP request
	req, err := http.NewRequest(verb, url, ioReader)
	if err != nil {
		return nil, err
	}

	// Set headers
	for key, value := range headers {
		req.Header.Set(key, value)
	}

	// Set content type header
	req.Header.Set("Content-Type", "application/json")

	// Perform the HTTP request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Check the response status code
	if resp.StatusCode >= 400 {
		return nil, errors.New("HTTP request failed with status: " + resp.Status)
	}

	// Decode the response JSON
	var responseBody V
	err = json.NewDecoder(resp.Body).Decode(&responseBody)
	if err != nil {
		return nil, err
	}

	return &responseBody, nil
}
