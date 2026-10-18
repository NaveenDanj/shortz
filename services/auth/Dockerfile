FROM golang:1.25-alpine AS builder

WORKDIR /app

COPY go.work go.work.sum ./
COPY config/go.mod config/go.sum ./config/
COPY contract/go.mod ./contract/
COPY authclient/go.mod authclient/go.sum ./authclient/
COPY services/auth/go.mod services/auth/go.sum ./services/auth/
COPY services/users/go.mod services/users/go.sum ./services/users/
RUN go mod download

COPY config ./config
COPY contract ./contract
COPY authclient ./authclient
COPY services/auth ./services/auth

RUN CGO_ENABLED=0 GOOS=linux \
    go build -o /bin/auth ./services/auth/cmd/server

FROM gcr.io/distroless/static-debian12

COPY --from=builder /bin/auth /auth

EXPOSE 8080

ENTRYPOINT ["/auth"]
