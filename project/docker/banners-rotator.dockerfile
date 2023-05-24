
FROM golang:1.19 as build

ENV BIN_FILE /opt/banners-rotator/banners-app
ENV CODE_DIR /go/src/

WORKDIR ${CODE_DIR}

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . ${CODE_DIR}

ARG LDFLAGS
RUN CGO_ENABLED=0 go build \
        -ldflags "$LDFLAGS" \
        -o ${BIN_FILE} cmd/*

FROM alpine:3.9

LABEL ORGANIZATION="OTUS"
LABEL SERVICE="banners-rotator"
LABEL MAINTAINERS="a.polikarpov@ficto.ru"

ENV BIN_FILE "/opt/banners-rotator/banners-app"
COPY --from=build ${BIN_FILE} ${BIN_FILE}

CMD ${BIN_FILE} -config ${CONFIG_FILE}