FROM golang:1.22-alpine AS builder

WORKDIR /src

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o /netbox-cli .


FROM alpine:3.21

COPY --from=builder /netbox-cli /usr/local/bin/netbox-cli

ENTRYPOINT ["netbox-cli"]
