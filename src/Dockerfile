FROM golang:1.23-alpine AS builder

WORKDIR /app
COPY . .
RUN go mod tidy
RUN CGO_ENABLED=0 GOOS=linux go build -o workers ./workers/main.go
RUN ls -la workers

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/workers ./
RUN ls -la workers && chmod +x workers
CMD ["/app/workers"]
