# This dockerfile uses extends image https://hub.docker.com/bridgewwater/gin-api-swagger-temple
# VERSION 1
# Author: bridgewwater
# dockerfile official document https://docs.docker.com/engine/reference/builder/
# https://hub.docker.com/_/golang
FROM golang:1.23.8 AS golang-builder

ARG TARGETPLATFORM
ARG BUILDPLATFORM
ARG ENV_BUILD_DIST_CODE_MARK=unknown
ARG GO_ENV_PACKAGE_NAME=github.com/bridgewwater/gin-api-swagger-temple
ARG GO_ENV_ROOT_BUILD_BIN_NAME=gin-api-swagger-temple
ARG GO_ENV_ROOT_BUILD_BIN_PATH=build/${GO_ENV_ROOT_BUILD_BIN_NAME}
ARG GO_ENV_ROOT_BUILD_ENTRANCE="cmd/gin-api-swagger-temple/main.go"

ARG GO_PATH_SOURCE_DIR=/go/src
WORKDIR ${GO_PATH_SOURCE_DIR}

RUN mkdir -p ${GO_PATH_SOURCE_DIR}/${GO_ENV_PACKAGE_NAME}
COPY . ${GO_PATH_SOURCE_DIR}/${GO_ENV_PACKAGE_NAME}

# proxy golang
#RUN go env -w "GOPROXY=https://goproxy.cn,direct"
RUN go env -w "GOPRIVATE='*.gitlab.com,*.gitee.com"

RUN echo "build running on $BUILDPLATFORM, building for $TARGETPLATFORM"

RUN cd ${GO_PATH_SOURCE_DIR}/${GO_ENV_PACKAGE_NAME} && \
    go mod download -x

RUN  cd ${GO_PATH_SOURCE_DIR}/${GO_ENV_PACKAGE_NAME} && \
  CGO_ENABLED=0 \
  go build \
  -a \
  -v \
  -installsuffix cgo \
  -ldflags '-X main.buildID=${ENV_BUILD_DIST_CODE_MARK} -w -s --extldflags "-static -fpic"' \
  -tags netgo \
  -o ${GO_ENV_ROOT_BUILD_BIN_PATH} \
  ${GO_ENV_ROOT_BUILD_ENTRANCE}

# https://hub.docker.com/_/alpine
FROM alpine:3.17

#ARG DOCKER_CLI_VERSION=${DOCKER_CLI_VERSION}
ARG GO_ENV_PACKAGE_NAME=github.com/bridgewwater/gin-api-swagger-temple
ARG GO_ENV_ROOT_BUILD_BIN_NAME=gin-api-swagger-temple
ARG GO_ENV_ROOT_BUILD_BIN_PATH=build/${GO_ENV_ROOT_BUILD_BIN_NAME}

ARG GO_PATH_SOURCE_DIR=/go/src

#RUN apk --no-cache add \
#  ca-certificates mailcap curl \
#  && rm -rf /var/cache/apk/* /tmp/*

RUN mkdir /app
WORKDIR /app

COPY --from=golang-builder ${GO_PATH_SOURCE_DIR}/${GO_ENV_PACKAGE_NAME}/${GO_ENV_ROOT_BUILD_BIN_PATH} .
COPY --from=golang-builder ${GO_PATH_SOURCE_DIR}/${GO_ENV_PACKAGE_NAME}/conf/config.yaml ./conf/config.yaml
COPY --from=golang-builder ${GO_PATH_SOURCE_DIR}/${GO_ENV_PACKAGE_NAME}/docs/*.json ./docs/
COPY --from=golang-builder ${GO_PATH_SOURCE_DIR}/${GO_ENV_PACKAGE_NAME}/docs/*.yaml ./docs/
ENTRYPOINT [ "/app/gin-api-swagger-temple" ]
# CMD ["/app/gin-api-swagger-temple", "--help"]