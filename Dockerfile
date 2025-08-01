# etapa de build
FROM golang:1.24.4 AS builder
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

# muda para onde está o main.go
WORKDIR /app/cmd
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -o rest-api .

# etapa final
FROM alpine:latest
RUN apk add --no-cache ca-certificates && addgroup -S appgroup && adduser -S appuser -G appgroup

WORKDIR /app
COPY --from=builder /app/cmd/rest-api .

USER appuser
EXPOSE 8080
ENTRYPOINT ["./rest-api"]
