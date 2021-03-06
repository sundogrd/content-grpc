#!/usr/bin/env bash

cp ./config/service.template.config.json ./config/service.config.json

sed -i "s/Your Username/$1/g" "./config/service.config.json"
sed -i "s/Your Password/$2/g" "./config/service.config.json"


ip -4 route list match 0/0 | awk '{print $3 " host.docker.internal"}' >> /etc/hosts