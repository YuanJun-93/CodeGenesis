.PHONY: run gen tidy help

# Start backend server
run:
	@echo "Starting server..."
	go run cmd/server/main.go -f configs/code-genesis-api.yaml

# Generate API code
gen:
	@echo "Generating API code..."
	goctl api go -api api/code_genesis.api -dir . -style gozero
	go mod tidy

# Run go mod tidy
tidy:
	go mod tidy

# Show help
help:
	@echo "Usage:"
	@echo "  make run   - Start the backend server"
	@echo "  make gen   - Generate API code from .api file"
	@echo "  make swagger - Generate OpenAPI/Swagger spec"
	@echo "  make test  - Run unit tests"
	@echo "  make tidy  - Run go mod tidy"

# Run unit tests (skips integration tests)
test:
	@echo "Running unit tests..."
	go test -v ./...

# Run API integration tests (costs money)
test-api:
	@echo "Running API integration tests..."
	go test -v -tags integration ./internal/logic/generator/...

# Generate Swagger file
swagger:
	@echo "Generating Swagger..."
	goctl api plugin -plugin goctl-swagger="swagger -filename swagger.json" -api api/code_genesis.api -dir api/doc
