# ========== Build stage ==========
FROM golang:1.25.1 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o restaurant-api ./cmd/api

# ========== Runtime stage ==========
FROM alpine:3.20

WORKDIR /app
COPY --from=builder /app/restaurant-api .

EXPOSE 8080
CMD ["./restaurant-api"]