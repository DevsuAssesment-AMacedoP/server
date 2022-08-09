# syntax=docker/dockerfile:1

## Build
FROM golang:1.19-alpine AS builder

RUN apk add --no-cache make

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY . .

RUN make build

## Deploy
FROM golang:1.19-alpine

WORKDIR /

COPY --from=builder /app/bin/server /

USER 1000:1000

ENTRYPOINT ["/server"]