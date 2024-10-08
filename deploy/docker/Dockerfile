# Build stage
FROM golang:1.22-alpine3.20 AS builder

# The latest alpine images don't have some tools like (`git` and `bash`).
# Adding git, bash and openssh to the image
RUN apk add --no-cache git make bash ca-certificates tzdata \
    --repository http://mirrors.aliyun.com/alpine/v3.11/community \
    --repository http://mirrors.aliyun.com/alpine/v3.11/main

RUN GRPC_HEALTH_PROBE_VERSION=v0.3.0 && \
    wget -qO/go/bin/grpc_health_probe \
    https://github.com/grpc-ecosystem/grpc-health-probe/releases/download/${GRPC_HEALTH_PROBE_VERSION}/grpc_health_probe-linux-amd64 && \
    chmod +x /go/bin/grpc_health_probe

# 镜像设置必要的环境变量
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 \
    GOPROXY="https://goproxy.cn,direct" \
    TZ=Asia/Shanghai \
    APP_ENV=docker

# 移动到工作目录
WORKDIR /go/src/chatgram

# 复制项目中的 go.mod 和 go.sum文件并下载依赖信息
COPY go.mod .
COPY go.sum .

# 将代码复制到容器中
COPY . .
COPY config ./config

# Build the Go app
#RUN make build
RUN make build

# Runtime stage
FROM alpine:3.20

WORKDIR /bin

ENV APP_ENV prod

# 从builder镜像中把 /build 拷贝到当前目录
COPY --from=builder /go/src/chatgram/build/chatgram   /bin/chatgram
COPY --from=builder /go/src/chatgram/config           /data/conf/chatgram/config
COPY --from=builder /go/bin/grpc_health_probe         /bin/grpc_health_probe


RUN apk update \
 && apk add --no-cache curl jq \
 && rm -rf /var/cache/apk/* \
 && mkdir -p  /data/logs/

# Expose port 8080 to the outside world
EXPOSE 8080
EXPOSE 9090

# 需要运行的命令
ENTRYPOINT ["/bin/chatgram", "-c", "/data/conf/chatgram/config"]

# 1. build image: docker build -t qloog/chatgram:v1.0.0 -f deploy/docker/Dockerfile .
# 2. start: docker run --rm -it -p 8080:8080 chatgram:v1.0.0
# 3. test: curl -i http://localhost:8080/health
# 4. tag: docker tag chatgram:v1.0.0 qloog/chatgram:v1.0.0
# 5. push tag: docker push qloog/chatgram:v1.0.0
