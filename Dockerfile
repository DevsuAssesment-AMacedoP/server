# syntax=docker/dockerfile:1

## Build
FROM golang:1.19-alpine AS builder

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY *.go ./
COPY Makefile ./

RUN make build

## Deploy
FROM gcr.io/distroless/base-debian10

WORKDIR /

COPY --from=builder /app/bin/server /server

EXPOSE 5000

USER nonroot:nonroot

ENTRYPOINT ["/server"]