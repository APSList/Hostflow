# --- BUILD STAGE ---
FROM golang:1.25-alpine AS builder
WORKDIR /app

# Kopiraj dependencies
COPY go.mod go.sum ./
RUN go mod download

# Kopiraj kodo
COPY . .

# Generiraj swagger dokumentacijo
RUN go install github.com/swaggo/swag/cmd/swag@latest
RUN swag init --parseDependency --parseInternal --parseDepth 1

# Zgradi binarko
RUN go build -o payment-service .

# --- FINAL STAGE ---
FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/payment-service .
ENTRYPOINT ["./payment-service"]
