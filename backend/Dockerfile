FROM golang:1.21-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . ./

COPY config/config.yaml /app/config/config.yaml

RUN go build -o booklib

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/booklib .

EXPOSE 8080

CMD ["./booklib"]
