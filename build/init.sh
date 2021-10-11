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

    # 构建服务端镜像
    if [ $# -ge 3 ] && [ $3 = "https" ]; then
        docker build -f ./build/https/Dockerfile -t 18509518245/config-server-https:$serverTag .
    else
        docker build -f ./build/http/Dockerfile -t 18509518245/config-server-http:$serverTag .
    fi

elif [ $baseServer = "frontend" ]; then

    # 构建前端镜像
    docker build -f ./web/build/Dockerfile -t 18509518245/config-web:$serverTag .
    docker build -f ./web/run/Dockerfile -t 18509518245/config-nginx:$serverTag .
    
 else
    echo "参数错误"
fi
