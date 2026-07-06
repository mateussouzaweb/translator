# syntax=docker/dockerfile:1
# Build stage
FROM golang:1.26.4-alpine AS builder

WORKDIR /src
COPY go.mod go.sum ./
RUN go mod download

COPY . .
ENV CGO_ENABLED=0 GOOS=linux
RUN go build -buildvcs=false -o /usr/local/bin/translator ./cmd/translator

# Runtime stage
FROM alpine:latest AS runtime

# Install dependencies
RUN apk --no-cache add gcompat

# Copy binary
COPY --from=builder /usr/local/bin/translator /usr/local/bin/

# Run program
WORKDIR /app
ENTRYPOINT ["translator"]
CMD ["--help"]