#!/bin/bash

if [ $# -lt 2 ]; then
    echo "usage [ backend|frontend ] [ tag ]"
    exit 1
fi

baseServer=$1
serverTag=$2
echo "build server, based on "$baseServer" tag "$serverTag

# 根目录下构建
cd ..
if [ $baseServer = "backend" ]; then
    # 构建服务镜像
    docker build -f ./build/Dockerfile -t 18509518245/config-server:$serverTag .
elif [ $baseServer = "frontend" ]; then
    # 构建前端镜像
    docker build -f ./web/build/Dockerfile -t 18509518245/config-web:$serverTag .
    docker build -f ./web/run/Dockerfile -t 18509518245/config-nginx:$serverTag .
fi
