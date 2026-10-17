FROM golang:1.25-alpine AS builder

WORKDIR /app

COPY go.work go.work.sum ./
COPY config/go.mod config/go.sum ./config/
COPY contract/go.mod ./contract/
COPY authclient/go.mod authclient/go.sum ./authclient/
COPY services/users/go.mod services/users/go.sum ./services/users/
RUN go mod download

COPY config ./config
COPY contract ./contract
COPY authclient ./authclient
COPY services/users ./services/users

RUN CGO_ENABLED=0 GOOS=linux \
    go build -o /bin/users ./services/users/cmd/server

FROM gcr.io/distroless/static-debian12

COPY --from=builder /bin/users /users
EXPOSE 8080

ENTRYPOINT ["/users"]
