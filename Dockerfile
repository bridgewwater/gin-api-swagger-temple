# This dockerfile uses extends image https://hub.docker.com/bridgewwater/golang-project-temple-base
# VERSION 1
# Author: sinlov
# dockerfile offical document https://docs.docker.com/engine/reference/builder/
# https://hub.docker.com/_/golang
FROM golang:1.17.13-buster as builder

ARG GO_PATH_SOURCE_DIR=/go/src/
WORKDIR ${GO_PATH_SOURCE_DIR}

RUN mkdir -p ${GO_PATH_SOURCE_DIR}/github.com/bridgewwater/golang-project-temple-base
COPY $PWD ${GO_PATH_SOURCE_DIR}/github.com/bridgewwater/golang-project-temple-base

RUN cd ${GO_PATH_SOURCE_DIR}/github.com/bridgewwater/golang-project-temple-base && \
    go mod download -x

RUN  cd ${GO_PATH_SOURCE_DIR}/github.com/bridgewwater/golang-project-temple-base && \
  CGO_ENABLED=0 \
  go build \
  -a \
  -installsuffix cgo \
  -ldflags '-w -s --extldflags "-static -fpic"' \
  -tags netgo \
  -o golang-project-temple-base \
  main.go

# https://hub.docker.com/_/alpine
FROM alpine:3.17

ARG DOCKER_CLI_VERSION=${DOCKER_CLI_VERSION}

#RUN apk --no-cache add \
#  ca-certificates mailcap curl \
#  && rm -rf /var/cache/apk/* /tmp/*

RUN mkdir /app
WORKDIR /app

COPY --from=builder /go/src/github.com/bridgewwater/golang-project-temple-base/golang-project-temple-base .
ENTRYPOINT ["/app/golang-project-temple-base"]
# CMD ["/app/golang-project-temple-base", "--help"]