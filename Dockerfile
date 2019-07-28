FROM golang:latest
MAINTAINER Breakinferno <1972952841@qq.com>

COPY . $GOPATH/src/github.com/sundogrd/content-grpc
WORKDIR $GOPATH/src/github.com/sundogrd/content-grpc

ENV GO111MODULE=on

RUN go build .

RUN "./devops/build_docker.sh"

EXPOSE 25555
  
ENTRYPOINT ["./devops/entrypoint.sh"]