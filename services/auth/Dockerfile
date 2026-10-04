FROM golang:1.25-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux \
    go build -o /bin/auth ./cmd/server

FROM gcr.io/distroless/static-debian12

COPY --from=builder /bin/auth /auth

EXPOSE 3001

ENTRYPOINT ["/auth"]