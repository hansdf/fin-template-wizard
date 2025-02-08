# Stage 1: Build the Go application
FROM golang:latest AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o template-wizard-cli .

# Stage 2: Create a lightweight runtime image
FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/template-wizard-cli .
COPY --from=builder /app/templates ./templates
CMD ["./template-wizard-cli"]