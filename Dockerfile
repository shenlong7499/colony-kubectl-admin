FROM golang:alpine as builder
WORKDIR /app/colony-service
COPY . .

RUN go env -w GO111MODULE=on \
    && go env -w CGO_ENABLED=0 \
    && go env \
    && go mod tidy \
    && go build -o server .

FROM alpine:latest

LABEL MAINTAINER="shenlong7499@163.com"

WORKDIR /app/colony-service
COPY --from=0 /app/colony-service/config.yaml ./config.yaml
COPY --from=0 /app/colony-service/.kube/config ./.kube/config
COPY --from=0 /app/colony-service/server ./
EXPOSE 8082
ENTRYPOINT ./server
