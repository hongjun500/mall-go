FROM golang:1.19.9-alpine3.18 AS builder

# 设置工作目录
WORKDIR /mall-go

COPY . .

# 设置环境变量并且使用国内代理禁用镜像源的缓存
ENV GOPROXY=https://goproxy.cn,direct

# 下载依赖并整理
RUN go mod download && go mod tidy

# 构建项目
RUN go build -o /mall-go/app /mall-go/cmd/main.go




# 使用alpine镜像作为基础镜像
FROM golang:1.19.9-alpine3.18 AS runner

# 设置工作目录
WORKDIR /mall-go

# 复制项目文件到镜像中
COPY --from=builder /mall-go/app .
# 这里使用 docker 相关的配置，目前主要是可以通过几个环境变量来配置项目
COPY --from=builder /mall-go/configs.docker ./configs
#COPY configs ./configs

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
ENTRYPOINT ["./app"]

ARG IMAGE_NAME=mall-go

LABEL name=$IMAGE_NAME \
      description="mall-go项目镜像" \
      version="1.0.0" \
      author="hongjun500" \
      email="" \
# Path: Dockerfile