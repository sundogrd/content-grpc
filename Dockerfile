FROM golang:latest
MAINTAINER Breakinferno <1972952841@qq.com>

COPY . $GOPATH/src/github.com/sundogrd/content-grpc
WORKDIR $GOPATH/src/github.com/sundogrd/content-grpc

ENV GO111MODULE=on
ARG DB_USER
ARG DB_PWD
RUN go build .

RUN "./devops/build_docker.sh" $DB_USER $DB_PWD

EXPOSE 25555
  
ENTRYPOINT ["./devops/entrypoint.sh"]