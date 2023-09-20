FROM golang:1.19.9-alpine3.18 AS builder

# 作者信息
LABEL maintainer="hongjun500 <"
LABEL version="1.0"
LABEL description="This is a docker image for golang project."

# 设置工作目录
WORKDIR /mall-go

COPY . .

# 设置环境变量并且使用国内代理禁用镜像源的缓存
ENV GOPROXY=https://goproxy.cn,direct

# 下载依赖、整理模块、构建项目以及清理构建缓存
RUN go mod download && \
    go mod tidy && \
    rm -f /docs && rm -f /logs && rm -f /scripts && \
    go build -o /mall-go/app /mall-go/cmd/main.go && \
    go clean -cache



# 使用 gcr.io/distroless/base 镜像作为基础镜像，由于镜像非常精简，因此减小了潜在的攻击面，提高了安全性。这对于安全敏感的应用程序很有用。
#FROM gcr.io/distroless/base AS runner

# 使用 alpine:latest 镜像作为基础镜像，对操作系统工具和库的依赖性，如果需要一些特定的工具或库，Alpine可能更适合
FROM alpine:latest AS runner

# 设置工作目录
WORKDIR /mall-go

# 复制项目文件到镜像中
COPY --from=builder ./mall-go/app .
# 这里使用 docker 相关的配置，目前主要是可以通过几个环境变量来配置项目
COPY --from=builder ./mall-go/configs.docker ./configs

# 将时区设置为东八区
RUN echo "https://mirrors.aliyun.com/alpine/v3.8/main/" > /etc/apk/repositories \
    && echo "https://mirrors.aliyun.com/alpine/v3.8/community/" >> /etc/apk/repositories \
    && apk update \
    && apk add --no-cache tzdata \
    && cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime  \
    && apk del tzdata \
    && echo "Asia/Shanghai" > /etc/timezone \

# 暴露三个端口
EXPOSE 8080 8081 8082


# 启动项目
CMD ["./app"]
#ENTRYPOINT ["./app"]