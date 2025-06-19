FROM golang:1.24-alpine AS builder

# Install ca-certificates.
RUN apk --no-cache add ca-certificates

# Move to working directory (/build).
WORKDIR /build

# Copy and download dependency using go mod.
COPY go.mod go.sum ./
RUN go mod download

# Copy your code into the container.
COPY . .

# Set necessary environment variables and build your project.
ENV CGO_ENABLED=0
RUN go build -ldflags="-s -w" -o app_analytics

FROM scratch

# Copy certificates from the builder image.
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

# Copy project's binary and templates from /build to the scratch container.
COPY --from=builder /build/app_analytics /

# Expose port.
EXPOSE 8080

# Set entry point.
ENTRYPOINT ["/app_analytics"]