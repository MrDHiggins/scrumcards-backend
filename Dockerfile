# ----------------------
# 1. Build Stage
# ----------------------
FROM golang:1.23 AS builder
WORKDIR /app

# Cache Go modules
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the application
COPY . .

# Build the Go binary (static build)
RUN CGO_ENABLED=0 GOOS=linux go build -o scrumdcards-backend ./cmd/server

# ----------------------
# 2. Runtime Stage (Distroless)
# ----------------------
FROM gcr.io/distroless/static-debian12

WORKDIR /app

COPY --from=builder /app/scrumdcards-backend .

EXPOSE 8080
USER nonroot:nonroot

CMD ["./scrumdcards-backend"]