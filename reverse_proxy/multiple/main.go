package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

// Map of paths to backend servers
var backendServers = map[string]string{
	"/echo/":   "http://localhost:8585",
	"/api-go/": "https://jsonplaceholder.typicode.com/users",
}

func main() {
	// Handle incoming requests and route them to the correct backend
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		for prefix, backend := range backendServers {
			if len(r.URL.Path) >= len(prefix) && r.URL.Path[:len(prefix)] == prefix {
				targetURL, _ := url.Parse(backend)
				proxy := httputil.NewSingleHostReverseProxy(targetURL)

				// Modify request before forwarding
				r.URL.Path = r.URL.Path[len(prefix)-1:] // Preserve endpoint structure
				proxy.ServeHTTP(w, r)
				return
			}
		}
		http.Error(w, "Not Found", http.StatusNotFound)
	})

	// Start the reverse proxy
	log.Println("Reverse proxy running on :8082...")
	log.Fatal(http.ListenAndServe(":8082", nil))
}
