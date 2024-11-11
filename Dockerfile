FROM golang:1.22.0-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o main cmd/currency-converter/main.go

FROM alpine:latest as runner

ENV GO_ENV=production

WORKDIR /app

COPY --from=builder /app/main .

# Expose the port on which the application runs (if applicable)
EXPOSE 8080

# Command to run the application
CMD ["./main"]
