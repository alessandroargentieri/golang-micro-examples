package main

import (
	"net/http"
	"net/http/httputil"
	"net/url"
	"log"
)
// modified
func main() {
	// Target backend server (change this to your actual backend URL)
	targetURL, _ := url.Parse("https://echo.free.beeceptor.com")

	// Create the reverse proxy
	proxy := httputil.NewSingleHostReverseProxy(targetURL)

	// Handle incoming requests and proxy them
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		proxy.ServeHTTP(w, r)
	})

	// Start the server
	log.Println("Reverse proxy running on :8082...")
	log.Fatal(http.ListenAndServe(":8082", nil))
}

