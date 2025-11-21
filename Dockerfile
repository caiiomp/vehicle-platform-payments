FROM golang:1.25.2-alpine AS builder

RUN apk add --no-cache git

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o vehicle-platform-payments ./src/main.go

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/vehicle-platform-payments .

CMD ["./vehicle-platform-payments"]
