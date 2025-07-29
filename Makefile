# Configuration variables
BINARY_NAME=product-api
MAIN_PATH=./cmd/api
DOCS_PATH=./docs

.PHONY: install-tools up build down test swagger run help lint gen

# Docker compose commands
up:
	docker-compose up -d
	@echo "🚀 Services started with sample data!"
	@echo "📚 API: http://localhost:8080"
	@echo "📊 Swagger: http://localhost:8080/swagger/index.html"

build:
	docker-compose up -d --build
	@echo "🚀 Services started with sample data!"
	@echo "📚 API: http://localhost:8080"
	@echo "📊 Swagger: http://localhost:8080/swagger/index.html"

down:
	docker-compose down

# Development commands
install-tools:
	go install github.com/swaggo/swag/cmd/swag@v1.8.12
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	go install github.com/golang/mock/mockgen@latest

swagger:
	@rm -rf $(DOCS_PATH)
	swag init -g $(MAIN_PATH)/main.go -o $(DOCS_PATH) --parseInternal --parseDependency

run: swagger
	go run $(MAIN_PATH)

test:
	go test ./... -v

lint:
	golangci-lint run

gen:
	go generate ./...

help:
	@echo "Available commands:"
	@echo "  up            - Start all services with sample data"
	@echo "  down          - Stop all services"
	@echo "  build         - Build and start services"
	@echo "  run           - Run API locally"
	@echo "  test          - Run tests"
	@echo "  swagger       - Generate Swagger docs"
	@echo "  gen           - Generate mocks"
	@echo "  lint          - Run linter"
	@echo "  install-tools - Install dev tools"
