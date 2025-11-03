# --- BUILD STAGE ---
FROM golang:1.25-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o property-service .

# --- FINAL STAGE ---
FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/property-service .
ENTRYPOINT ["./property-service"]
