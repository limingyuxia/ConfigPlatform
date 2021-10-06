#!/bin/bash

if [ $# -lt 2 ]; then
    echo "usage [ mysql|backend|frontend ] [ tag ]"
    exit 1
fi

baseServer=$1
serverTag=$2
echo "run server, based on "$baseServer" tag "$serverTag

if [ $baseServer = "mysql" ]; then
    # 构建数据库
    docker run --name configDb -p 3306:3306 -e MYSQL_ROOT_PASSWORD=123456 -d mysql:$serverTag
    # 等待数据库初始化
    sleep 30
    # 初始化table
    mysql -h limingyuxia.asia -u root -p123456 < $(pwd)/../docs/sql/config.sql
elif [ $baseServer = "backend" ]; then
    # 运行服务端容器
    docker run -d --rm -p 8000:8000 --name config-server 18509518245/config-server:$serverTag
elif [ $baseServer = "frontend" ]; then
    # 运行前端容器
    container_id=$(docker create 18509518245/config-web:$serverTag)
    docker cp $container_id:/home/dist $(pwd)/../web
    docker rm $container_id

    # -d 后台运行 -rm 如果有同名的容器，自动删除
    docker run -d --rm -p 9528:9528 -p 9529:9529 -v $(pwd)/../web/dist:/usr/share/nginx/html:ro \
        --name config-nginx 18509518245/config-nginx:$serverTag
fi
