#!/bin/bash

# Complete setup and demonstration script

echo "🚀 Setting up OpenTelemetry + Prometheus + Grafana demonstration"
echo

# Create directory structure
echo "📁 Creating directory structure..."
mkdir -p grafana/provisioning/datasources
mkdir -p grafana/provisioning/dashboards
mkdir -p grafana/dashboards

echo "📋 Required directory structure:"
echo "otel-prometheus-example/"
echo "├── main.go"
echo "├── go.mod"
echo "├── go.sum" 
echo "├── Dockerfile"
echo "├── docker-compose.yml"
echo "├── prometheus.yml"
echo "├── .gitignore"
echo "└── grafana/"
echo "    ├── provisioning/"
echo "    │   ├── datasources/"
echo "    │   │   └── datasource.yml"
echo "    │   └── dashboards/"
echo "    │       └── dashboard.yml"
echo "    └── dashboards/"
echo "        └── otel-demo-dashboard.json"
echo

echo "⚙️  Make sure all configuration files are created from the artifacts..."
echo "✅ Files needed:"
echo "   - prometheus.yml (Prometheus config)"
echo "   - grafana/provisioning/datasources/datasource.yml (Grafana datasource)"
echo "   - grafana/provisioning/dashboards/dashboard.yml (Dashboard provisioning)"
echo "   - grafana/dashboards/otel-demo-dashboard.json (Actual dashboard)"
echo

# Build and start all services
echo "🐳 Building and starting services..."
docker compose up --build -d

echo
echo "✅ Services are starting up..."
echo
echo "📋 Available endpoints:"
echo "   🔹 Go App:        http://localhost:8080"
echo "   🔹 Go Metrics:    http://localhost:9090/metrics"
echo "   🔹 Prometheus:    http://localhost:9091"
echo "   🔹 Grafana:       http://localhost:3000 (admin/admin)"
echo
echo "🎯 Direct dashboard link (after startup):"
echo "   📊 Dashboard:     http://localhost:3000/d/otel-demo/opentelemetry-demo-dashboard"
echo
echo "🧪 Generate some test traffic:"
echo "   curl http://localhost:8080/"
echo "   curl http://localhost:8080/api/data"
echo "   curl http://localhost:8080/health"
echo

# Wait for services to be ready
echo "⏳ Waiting for services to be ready..."
sleep 30

# Generate some test traffic
echo "🔄 Generating test traffic..."
for i in {1..20}; do
    curl -s http://localhost:8080/ > /dev/null
    curl -s http://localhost:8080/api/data > /dev/null
    curl -s http://localhost:8080/health > /dev/null
    sleep 1
done

echo
echo "✨ Setup complete! The dashboard should be automatically available at:"
echo "   🎯 http://localhost:3000/d/otel-demo/opentelemetry-demo-dashboard"
echo
echo "📊 Dashboard includes:"
echo "   - Total HTTP requests counter"
echo "   - Request rate over time"
echo "   - Pie charts for requests by path and status code"
echo "   - Detailed table view"
echo

# Show current metrics
echo "📊 Current metrics sample:"
curl -s http://localhost:9090/metrics | grep http_requests_total | head -5