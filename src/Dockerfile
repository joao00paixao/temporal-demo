FROM golang:1.23-alpine AS builder

WORKDIR /app
COPY . .
RUN go mod tidy
RUN CGO_ENABLED=0 GOOS=linux go build -o ./worker ./workers/main.go 

FROM alpine:latest
WORKDIR /app
RUN ls
COPY --from=builder /app/worker ./

CMD ["/app/worker"]
