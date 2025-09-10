FROM golang:1.24.4-alpine AS builder

RUN apk add --no-cache git

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -o /migrate ./cmd/migrate/main.go

RUN addgroup -g 10001 -S appgroup && adduser -u 10001 -S -D -G appgroup appuser

FROM alpine:3.22
WORKDIR /app

COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /etc/group /etc/group

COPY --from=builder /migrate /migrate
COPY migrations/ /migrations

USER appuser:appgroup

ENTRYPOINT ["/migrate"]
