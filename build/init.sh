#!/bin/bash
# 运行mysql
docker run --name configDb -p 3306:3306 -e MYSQL_ROOT_PASSWORD=123456 -d mysql:latest

# 执行sql文件
mysql -h 127.0.0.1 -u root -p123456 < $(pwd)/../docs/sql/config.sql

# 构建服务镜像
docker build -f ./build/Dockerfile -t 18509518245/config-server:v0.1 .

# 运行容器
docker run -it --rm -p 8000:8000 -p 9527:9527 --name config-platform 18509518245/config-server:v0.1

# 构建前端镜像
cd web
docker build -f ./Dockerfile -t 18509518245/config-web:v0.1 .

# 运行容器
docker run -it --rm -p 9528:9528 --name config-web 18509518245/config-web:v0.1
