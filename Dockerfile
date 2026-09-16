FROM golang:1.23-alpine AS builder

RUN apk add --no-cache gcc musl-dev

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=1 GOOS=linux go build -o forum-app ./cmd/web

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/forum-app .

COPY --from=builder /app/schema.sql .
COPY --from=builder /app/ui ./ui

EXPOSE 8080

CMD ["./forum-app"]
