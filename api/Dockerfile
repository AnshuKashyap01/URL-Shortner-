# Build Stage
FROM golang:alpine AS builder

WORKDIR /build

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o main .

# Runtime Stage
FROM alpine

RUN adduser -S -D -H -h /app appuser

WORKDIR /app

COPY --from=builder /build/main .
COPY --from=builder /build/static ./static

USER appuser

EXPOSE 3000

CMD ["./main"]