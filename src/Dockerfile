FROM golang:1.23-alpine AS builder

WORKDIR /app
COPY . .
RUN go mod tidy

RUN CGO_ENABLED=0 GOOS=linux go build -o ./worker ./workers/main.go
RUN CGO_ENABLED=0 GOOS=linux go build -o ./starter ./starter/main.go

FROM alpine:latest
WORKDIR /app

COPY --from=builder /app/worker ./
COPY --from=builder /app/starter ./

CMD ["/bin/sh", "-c", "./worker && ./starter"][1][3]
