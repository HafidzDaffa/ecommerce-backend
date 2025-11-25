# Build stage
FROM golang:1.23-alpine AS builder

WORKDIR /app

# Set Go proxy and timeout
ARG GOPROXY=https://proxy.golang.org,direct
ENV GOPROXY=${GOPROXY}
ENV GOPRIVATE=""
ENV GOSUMDB=sum.golang.org
ENV GOTOOLCHAIN=auto

# Install build dependencies
RUN apk add --no-cache git ca-certificates

# Copy go mod files
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main ./cmd/api

# Runtime stage
FROM alpine:latest

RUN apk --no-cache add ca-certificates tzdata

WORKDIR /root/

# Copy binary from builder
COPY --from=builder /app/main .

# Copy .env file if needed (optional, better to use docker-compose environment)
# COPY --from=builder /app/.env .

EXPOSE 8080

CMD ["./main"]
