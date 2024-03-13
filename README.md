## init

```
gin
go get -u github.com/gin-gonic/gin@v1.8.1
import "github.com/gin-gonic/gin"

viper
go get github.com/spf13/viper@v1.13.0
```

## k8s

```
client-go
go get k8s.io/client-go@v0.20.4
```

## 集成

1. 编写 Dockerfile: 构建 + 运行，两步

```
FROM golang:alpine as builder
WORKDIR /Users/longshen/GolandProjects/awesomeProject
COPY . .

RUN go env -w GO111MODULE=on \
    && go env -w CGO_ENABLED=0 \
    && go env \
    && go mod tidy \
    && go build -o server .

FROM alpine:latest

LABEL MAINTAINER="shenlong7499@163.com"

WORKDIR /Users/longshen/GolandProjects/awesomeProject
COPY --from=0 /Users/longshen/GolandProjects/awesomeProject/config.yaml ./config.yaml
COPY --from=0 /Users/longshen/GolandProjects/awesomeProject/.kube/config ./.kube/config
COPY --from=0 /Users/longshen/GolandProjects/awesomeProject/server ./
EXPOSE 8080
ENTRYPOINT ./server
```

2. 打包、推送远端仓库

```
docker build -t harbor.quest/privacy/colony-service:v1.0
docker login harbor.quest
docker push colony-service:v1.0
```