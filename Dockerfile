# Build Golang
FROM docker.io/bitnami/golang:1.20-debian-11 AS builder
WORKDIR /opt
COPY . .
RUN go version
RUN echo "Build Golang"
RUN go env -w GO111MODULE=on
RUN go env -w GOPROXY=https://goproxy.cn,https://goproxy.io,direct
RUN CGO_ENABLED=0 GOARCH=amd64 go build -o /opt/app

FROM docker.io/debian:10.10-slim
WORKDIR /opt
ENV TZ=Asia/Shanghai
COPY --from=builder /opt/app /opt/app
COPY --from=builder /opt/configs/dev/config.yaml configs/dev/
EXPOSE 8080
CMD ["/opt/app"]


