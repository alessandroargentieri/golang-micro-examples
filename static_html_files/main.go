package main

import (
    "net/http"
    "fmt"
)

func main() {
    // Serve static files from the "public" directory
    fs := http.FileServer(http.Dir("public"))
    http.Handle("/", fs)

    // Start the HTTP server on port 8080
    fmt.Println("Server started on port 8081...")
    http.ListenAndServe(":8081", nil)
}

