FROM golang:1.20-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod tidy

COPY . ./

RUN go build -o billing-ingestor cmd/api/main.go

FROM alpine:latest

COPY --from=builder /app/billing-ingestor /billing-ingestor

EXPOSE 8081

CMD ["/billing-ingestor"]
