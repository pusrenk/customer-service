# Go parameters
BINARY_NAME=customer-service
MAIN_PATH=./cmd/customer-service/main.go
GO=go
GOFMT=gofmt
GOLINT=golangci-lint
MOCKERY=mockery

# Build flags
LDFLAGS=-ldflags "-s -w"

.PHONY: all build clean test coverage lint fmt help mocks upgrade-deps proto

all: clean build

build:
	@echo "Building..."
	$(GO) build $(LDFLAGS) -o $(BINARY_NAME) $(MAIN_PATH)

run:
	@echo "Running..."
	$(GO) run $(MAIN_PATH)

clean:
	@echo "Cleaning..."
	rm -f $(BINARY_NAME)
	rm -f coverage.out

test:
	@echo "Running tests..."
	$(GO) test -v ./...

coverage:
	@echo "Running tests with coverage..."
	$(GO) test -coverprofile=coverage.out ./...
	$(GO) tool cover -html=coverage.out

lint:
	@echo "Running linter..."
	$(GOLINT) run

fmt:
	@echo "Formatting code..."
	$(GOFMT) -w .

deps:
	@echo "Downloading dependencies..."
	$(GO) mod download

tidy:
	@echo "Tidying dependencies..."
	$(GO) mod tidy

mocks:
	@echo "Generating mocks..."
	$(MOCKERY) --all --output ./test/mocks --outpkg mocks

upgrade-deps:
	@echo "Upgrading dependencies..."
	$(GO) get -u ./...
	$(GO) mod tidy

proto:
	@echo "Generating protobuf code..."
	@for proto_file in internal/protobuf/*.proto; do \
		protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative $$proto_file; \
	done

help:
	@echo "Available targets:"
	@echo "  all        	- Clean and build the project"
	@echo "  build      	- Build the project"
	@echo "  run        	- Run the project"
	@echo "  clean      	- Remove build artifacts"
	@echo "  test       	- Run tests"
	@echo "  coverage   	- Run tests with coverage report"
	@echo "  lint       	- Run linter"
	@echo "  fmt        	- Format code"
	@echo "  deps       	- Download dependencies"
	@echo "  tidy       	- Tidy dependencies"
	@echo "  mocks      	- Generate mocks for testing"
	@echo "  upgrade-deps 	- Upgrade dependencies"
	@echo "  proto      	- Generate protobuf code"
	@echo "  help       	- Show this help message"