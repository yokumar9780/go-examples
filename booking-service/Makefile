# Project Variables
APP_NAME := booking-api
MAIN := ./cmd/main.go
BIN := bin/$(APP_NAME)

# Commands
GO := go
SWAG := $(GOPATH)/bin/swag

.PHONY: all build run swag clean docker fmt test tidy

# Build the binary
build:
	$(GO) build -o $(BIN) $(MAIN)

# Run the application
run: build
	./$(BIN)

# Generate Swagger docs
swag:
	@if [ ! -f $(SWAG) ]; then \
		$(GO) install github.com/swaggo/swag/cmd/swag@latest; \
	fi
	$(SWAG) init --generalInfo $(MAIN) --output ./docs

# Format the code
fmt:
	$(GO) fmt ./...

# Clean up binaries and docs
clean:
	rm -rf bin/
	rm -rf docs/

# Tidy up modules
tidy:
	$(GO) mod tidy

# Build Docker image
docker:
	docker build -t $(APP_NAME):latest .

# Run tests
test:
	$(GO) test ./...
