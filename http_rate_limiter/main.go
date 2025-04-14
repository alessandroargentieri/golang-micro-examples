package main

import (
    "log"
    "net/http"
    "time"

    "github.com/ulule/limiter"
    "github.com/ulule/limiter/drivers/middleware/stdlib"
    "github.com/ulule/limiter/drivers/store/memory"
)

// rateLimitMiddleware is a middleware that limits the number of requests per IP address.
func rateLimitMiddleware(next http.Handler) http.Handler {
    // Create a new limiter store using in-memory storage.
    rate := limiter.Rate{
        Period: 1 * time.Second,
        Limit:  6, // 6 requests per second.
    }
    store := memory.NewStore()
    limitMiddleware := stdlib.NewMiddleware(limiter.New(store, rate))

    return limitMiddleware.Handler(next)
}

func main() {
    // Create a new HTTP handler.
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("Hello, World!"))
    })

    // Wrap the HTTP handler with the rate limiting middleware.
    handler := rateLimitMiddleware(http.DefaultServeMux)

    // Start the HTTP server.
    log.Println("Server is listening on port 8080")
    if err := http.ListenAndServe(":8080", handler); err != nil {
        log.Fatal(err)
    }
}
