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
	@echo "  make tidy  - Run go mod tidy"
