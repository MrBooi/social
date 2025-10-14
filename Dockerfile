
# ===========================
# 🧱 Build Stage
# ===========================
FROM golang:1.25 AS builder

# Set working directory
WORKDIR /app

# Copy and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Install swag CLI and generate Swagger docs
RUN go install github.com/swaggo/swag/cmd/swag@latest
RUN swag init -g cmd/api/main.go -o docs

# Build the Go binary (static)
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o api cmd/api/*.go

# ===========================
# 🚀 Run Stage
# ===========================
FROM scratch

# Metadata
LABEL maintainer="Ayabonga Booi <you@example.com>" \
      description="Social API built with Go and Swagger docs" \
      version="1.0.0"

# Set working directory
WORKDIR /app

# Copy CA certificates and binary
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /app/api .
COPY --from=builder /app/docs ./docs

# Non-root user for safety
USER 1000

# Expose app port
EXPOSE 8080

# Run the app
CMD ["./api"]
