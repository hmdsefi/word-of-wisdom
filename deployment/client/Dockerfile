ARG TARGET_DIR=/app

# Build
FROM golang:1.19 as builder
ARG TARGET_DIR
ARG GOBIN

ENV GO111MODULE=on
ENV GOPROXY="https://goproxy.io,direct"

WORKDIR /go/src/client
# Copy go mod and sum files
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN TARGET_DIR=$TARGET_DIR make build-client

# Dev image
FROM builder AS dev
EXPOSE 8080
ENTRYPOINT /app/client

# Run
FROM alpine:3.12 as prod
ARG TARGET_DIR
ARG GOBIN
RUN apk add --no-cache ca-certificates make git \
    && rm -rf /var/cache/apk/*

WORKDIR /app
COPY --from=builder ${TARGET_DIR}/client client

EXPOSE 8080