# Use official Golang image for build stage
FROM golang:1.24.2 AS builder

WORKDIR /app

# Copy go.mod and go.sum first for caching
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code
COPY . .

# Generate Swagger docs
RUN go install github.com/swaggo/swag/cmd/swag@latest && swag init --generalInfo cmd/main.go --output ./docs

# Build the Go application
RUN go build -o main ./cmd/main.go

# Runtime stage with a minimal image
FROM gcr.io/distroless/base-debian11

WORKDIR /app
COPY --from=builder /app/main .
COPY --from=builder /app/docs ./docs

# Application port
EXPOSE 8080

ENTRYPOINT ["./main"]
