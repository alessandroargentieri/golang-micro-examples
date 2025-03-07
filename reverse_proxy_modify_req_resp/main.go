package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

// ReverseProxyHandler with ModifyResponse to add a header
func ReverseProxyHandler(target string) http.Handler {
	targetURL, err := url.Parse(target)
	if err != nil {
		log.Fatalf("Error parsing target URL: %v", err)
	}

	proxy := httputil.NewSingleHostReverseProxy(targetURL)

	// Modify response to add a header
	proxy.ModifyResponse = func(resp *http.Response) error {
		resp.Header.Add("X-Proxy-Modified", "true")
		return nil
	}

	return proxy
}

func main() {
	target := "https://echo.free.beeceptor.com" // Replace with actual backend URL

	http.Handle("/", ReverseProxyHandler(target))

	log.Println("Starting reverse proxy on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
