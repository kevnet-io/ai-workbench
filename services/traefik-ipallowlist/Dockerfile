# Build stage
FROM golang:1.21-alpine AS builder

WORKDIR /build

# Copy go mod files
COPY go.mod ./

# Copy source code
COPY main.go ./

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags="-w -s" -o ipallowlist .

# Final stage
FROM scratch

# Copy the binary from builder
COPY --from=builder /build/ipallowlist /ipallowlist

# Expose port
EXPOSE 8080

# Run the binary
ENTRYPOINT ["/ipallowlist"]
