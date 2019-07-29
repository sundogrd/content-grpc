#!/usr/bin/env bash
git pull && git checkout master && git branch | grep -v "master" | xargs git branch -D 

git checkout -b $1 origin/$1

if docker ps -a | grep -q sundogrd-content-grpc; then
    docker rm -f $(docker ps -a | grep sundogrd-content-grpc | awk '{print $1}')
fi

if docker images | grep -q  sundogrd/content-grpc; then
    docker rmi -f `docker images | grep sundogrd/content-grpc | awk '{print $3}'`
fi

docker pull sundogrd/content-grpc:$1-$2
docker run -d --name sundogrd-content-grpc -p 25555:25555 sundogrd/content-grpc:$1-$2
