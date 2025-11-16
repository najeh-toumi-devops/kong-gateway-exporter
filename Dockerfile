# Build stage
FROM golang:1.21-alpine AS builder
WORKDIR /app
COPY go.mod main.go ./
RUN go mod tidy
RUN go build -o kong-gateway-exporter main.go

# Final image
FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/kong-gateway-exporter ./
EXPOSE 9542
ENTRYPOINT ["./kong-gateway-exporter"]
CMD ["--kong-url=http://kong:8001","--port=9542"]
