# Open Telemetry and Prometheus example

Go HTTP server that uses OpenTelemetry to collect metrics and export them to Prometheus. This example will track request counts by path and HTTP status code.

# Key Features

## OpenTelemetry Integration:

Creates a Prometheus exporter for metrics
Sets up a meter provider with OpenTelemetry
Defines a counter metric for HTTP requests
Records metrics with labels for path and status code

## HTTP Server:

Multiple endpoints with different behaviors
`/` - Always returns 200
`/api/data` - Randomly returns 200, 400, or 500 status codes
`/health` - Health check endpoint
`/metrics` - Prometheus metrics endpoint (auto-generated)

## Metric Collection:

Tracks request count by HTTP path and status code
Uses OpenTelemetry attributes to add dimensions to metrics
Automatically exports metrics in Prometheus format

## Run it locally

Start the server:
```bash
go run main.go
```
Into another terminal, perform some calls:
```bash
# Make multiple requests to different endpoints
for i in {1..10}; do curl http://localhost:8080/; done
for i in {1..10}; do curl http://localhost:8080/api/data; done
for i in {1..5}; do curl http://localhost:8080/health; done
```
Check the Prometheus metrics:
```bash
curl http://localhost:9090/metrics
```
You'll get them into the Prometheus format:
```bash
# HELP http_requests_total Total number of HTTP requests
# TYPE http_requests_total counter
http_requests_total{path="/",status_code="200"} 10
http_requests_total{path="/api/data",status_code="200"} 6
http_requests_total{path="/api/data",status_code="400"} 2
http_requests_total{path="/api/data",status_code="500"} 2
http_requests_total{path="/health",status_code="200"} 5
```

## Docker build

```bash
# Build the Docker image
docker build -t otel-prometheus-server .

# Run the container
docker run -p 8080:8080 otel-prometheus-server

# Or run with name for easier management
docker run -d --name otel-server -p 8080:8080 otel-prometheus-server
```

# Full Demonstration with Prometheus server

To do a full demonstration (docker build included) using docker compose with the http server, the prometheus server, grafana with a pre-configured dashboard you can run:

```bash
./run-example.sh
```
When the service docker image is built, the containers are up and running and some traffic has been generated, you can visit the pre-configured dashboard:

http://localhost:3000/d/otel-demo/opentelemetry-demo-dashboard

where you can see in a Grafana dashboard the data collected from the Prometheus server which is periodically scraping data from the go server at the https://localhost:9090/metrics endpoint.

If you want to interact with the Prometheus server directly instead of Grafana, you can use the Prometheus API: at:
```bash
# basic query
curl http://localhost:9091/api/v1/query?query=http_requests_total | jq

# all requests by path:
curl http://localhost:9091/api/v1/query?query=sum(http_requests_total)%20by%20(path) | jq

# all requests by status code:
curl http://localhost:9091/api/v1/query?query=sum(http_requests_total)%20by%20(status_code) | jq

# request rate per minute:
curl http://localhost:9091/api/v1/query?query=rate(http_requests_total[1m]) | jq

# total number of requests:
curl http://localhost:9091/api/v1/query?query=sum(http_requests_total) | jq
```
While using Prometheus Web UI:
You can also use the Prometheus web interface at http://localhost:9091/query and enter these queries in the query box:
```
http_requests_total
sum(http_requests_total) by (path)
sum(http_requests_total) by (status_code)
rate(http_requests_total[5m])
```

