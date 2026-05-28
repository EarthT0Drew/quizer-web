# Builder
FROM golang:1.26.3-alpine3.23 AS builder
# RUN apk add --no-cache git # Only needed if we need a dependency that lives on git
WORKDIR /src
COPY go.mod ./
# COPY go.sum ./ # Only needed if we have a go.sum file, used if we're using outside modules
RUN go mod download
COPY ./ ./
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/quizer ./cmd/server

# Runtime
FROM alpine:3.23
# RUN apk add --no-cache ca-certificates # Only needed for HTTPS
WORKDIR /app
COPY --from=builder /app/quizer ./
# EXPOSE 8080 # Uncomment when API is ready
ENTRYPOINT ["/app/quizer"]