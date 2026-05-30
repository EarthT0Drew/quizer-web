# Builder
FROM golang:1.26.3-alpine3.23 AS builder
# There are packages that live on github
RUN apk add git
WORKDIR /src
COPY go.mod go.sum ./
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