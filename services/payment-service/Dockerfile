# --- BUILD STAGE ---
FROM golang:1.25-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN swag init --parseDependency --parseInternal --parseDepth 1
RUN go build -o payment-service .

# --- FINAL STAGE ---
FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/payment-service .
ENTRYPOINT ["./payment-service"]
