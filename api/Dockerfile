# Build Stage
FROM golang:alpine AS builder

WORKDIR /build

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o main .

# Runtime Stage
FROM alpine

RUN adduser -S -D -H -h /app appuser

WORKDIR /app

COPY --from=builder /build/main .
COPY --from=builder /build/.env .

USER appuser

EXPOSE 3000

CMD ["./main"]