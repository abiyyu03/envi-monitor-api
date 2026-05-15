FROM golang:1.26-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o envi-monitor ./app

FROM alpine:3.20

WORKDIR /app

COPY --from=builder /app/envi-monitor .
COPY --from=builder /app/.env .

EXPOSE 8080

CMD ["./envi-monitor"]
