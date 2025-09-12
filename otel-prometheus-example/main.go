package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/prometheus"
	"go.opentelemetry.io/otel/metric"
	sdkmetric "go.opentelemetry.io/otel/sdk/metric"
)

// HTTP Server with open telemetry operating to track requests
type Server struct {
	requestCounter metric.Int64Counter
}

// NewServer is the constructor of a new HTTP Server
func NewServer() (*Server, error) {
	// Create Prometheus exporter
	exporter, err := prometheus.New()
	if err != nil {
		return nil, fmt.Errorf("failed to create prometheus exporter: %w", err)
	}

	// Create a metric provider with the Prometheus exporter
	provider := sdkmetric.NewMeterProvider(
		sdkmetric.WithReader(exporter),
	)
	otel.SetMeterProvider(provider)

	// Start metrics server on port 9090
	go func() {
		fmt.Println("Metrics server starting on :9090")
		http.Handle("/metrics", promhttp.Handler())
		if err := http.ListenAndServe(":9090", nil); err != nil {
			log.Fatal("Metrics server failed to start:", err)
		}
	}()

	// Create a meter
	meter := otel.Meter("http-server-example")

	// Create a counter for HTTP requests
	requestCounter, err := meter.Int64Counter(
		"http_requests_total", // the metric we're going to export
		metric.WithDescription("Total number of HTTP requests"),
		metric.WithUnit("1"),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create request counter: %w", err)
	}

	return &Server{
		requestCounter: requestCounter,
	}, nil
}

// used to record the HTTP requests with OTEL
func (s *Server) recordRequest(ctx context.Context, path string, statusCode int) {
	s.requestCounter.Add(ctx, 1,
		metric.WithAttributes(
			attribute.String("path", path),
			attribute.Int("status_code", statusCode),
		),
	)
	// Debug output
	fmt.Printf("Recorded request: path=%s, status_code=%d\n", path, statusCode)
}

// handler for the HTTP endpoints on port 8080
func (s *Server) homeHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	// Simulate some processing time
	time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)

	statusCode := http.StatusOK
	w.WriteHeader(statusCode)
	w.Write([]byte("Hello from home!"))

	s.recordRequest(ctx, r.URL.Path, statusCode)
}

// handler for the HTTP endpoints on port 8080
func (s *Server) apiHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	// Simulate some processing time
	time.Sleep(time.Duration(rand.Intn(50)) * time.Millisecond)

	// Randomly return different status codes to show varied metrics
	statusCodes := []int{
		http.StatusOK,
		http.StatusOK,
		http.StatusOK,
		http.StatusBadRequest,
		http.StatusInternalServerError,
	}
	statusCode := statusCodes[rand.Intn(len(statusCodes))]

	w.WriteHeader(statusCode)

	switch statusCode {
	case http.StatusOK:
		w.Write([]byte(`{"message": "API response successful", "timestamp": "` + time.Now().Format(time.RFC3339) + `"}`))
	case http.StatusBadRequest:
		w.Write([]byte(`{"error": "Bad request"}`))
	case http.StatusInternalServerError:
		w.Write([]byte(`{"error": "Internal server error"}`))
	}

	s.recordRequest(ctx, r.URL.Path, statusCode)
}

// handler for the HTTP endpoints on port 8080
func (s *Server) healthHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	statusCode := http.StatusOK
	w.WriteHeader(statusCode)
	w.Write([]byte("OK"))

	s.recordRequest(ctx, r.URL.Path, statusCode)
}

// handler for the HTTP endpoints on port 8080
func (s *Server) notFoundHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	statusCode := http.StatusNotFound
	w.WriteHeader(statusCode)
	w.Write([]byte("Page not found"))

	s.recordRequest(ctx, r.URL.Path, statusCode)
}

func main() {
	// Initialize the server with OpenTelemetry
	server, err := NewServer()
	if err != nil {
		log.Fatal("Failed to initialize server:", err)
	}

	// Create a new ServeMux to avoid conflicts
	mux := http.NewServeMux()

	// Set up HTTP routes
	mux.HandleFunc("/", server.homeHandler)
	mux.HandleFunc("/api/data", server.apiHandler)
	mux.HandleFunc("/health", server.healthHandler)
	mux.HandleFunc("/favicon.ico", server.notFoundHandler)

	// Start the main server
	fmt.Println("Server starting on :8080")
	fmt.Println("Available endpoints:")
	fmt.Println("  GET  /          - Home page")
	fmt.Println("  GET  /api/data  - API endpoint (returns random status codes)")
	fmt.Println("  GET  /health    - Health check")
	fmt.Println()
	fmt.Println("Metrics server starting on :9090")
	fmt.Println("  GET  /metrics   - Prometheus metrics (auto-generated)")
	fmt.Println()
	fmt.Println("Try making requests to different endpoints, then check /metrics")
	fmt.Println("Example commands:")
	fmt.Println("  curl http://localhost:8080/")
	fmt.Println("  curl http://localhost:8080/api/data")
	fmt.Println("  curl http://localhost:8080/health")
	fmt.Println("  curl http://localhost:9090/metrics")

	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal("Server failed to start:", err)
	}
}
