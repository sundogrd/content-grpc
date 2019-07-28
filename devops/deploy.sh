#!/usr/bin/env bash

if docker images | grep -q  sundogrd/content-grpc; then
    docker rmi -f `docker images | grep sundogrd/content-grpc | awk '{print $3}'`
fi
if docker ps -a | grep -q sundogrd-content-grpc; then
    docker rm -f $(docker ps -a | grep sundogrd-content-grpc | awk '{print $1}')
fi

docker pull sundogrd/content-grpc:$1-$2
docker run -d --name sundogrd-content-grpc -p 25555:25555 sundogrd/content-grpc:$1-$2
