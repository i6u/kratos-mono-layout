FROM golang:1.16.6  AS builder
LABEL stage=builder

ARG APP_RELATIVE_PATH
ARG VERSION

COPY ./ /src
WORKDIR /src/app/${APP_RELATIVE_PATH}

ENV GO111MODULE=on
ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64
ENV GOPROXY=https://goproxy.cn,direct

RUN mkdir -p bin/ && go build -mod=vendor -ldflags "-X main.Version=$VERSION" -o ./bin/ ./...

FROM debian:stable-slim as alpine-basic

ARG APP_RELATIVE_PATH

ENV TZ='Asia/Shanghai'
#RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.ustc.edu.cn/g' /etc/apk/repositories &&\
#    apk update && apk upgrade && apk add --update --no-cache tzdata && apk add --update --no-cache ca-certificates && \
#    rm -rf /var/cache/apk/* \
RUN sed -i 's/deb.debian.org/mirrors.ustc.edu.cn/g' /etc/apt/sources.list \
    && apt-get update && apt-get install -y --no-install-recommends \
		ca-certificates  \
        netbase \
        && rm -rf /var/lib/apt/lists/ \
        && apt-get autoremove -y && apt-get autoclean -y

COPY --from=builder /src/app/${APP_RELATIVE_PATH}/bin /app

WORKDIR /app

EXPOSE 8000
EXPOSE 9000
VOLUME /data/conf

CMD ["./server", "-conf", "/data/conf"]