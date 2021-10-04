#!/bin/bash
docker run --name configDb -p 3306:3306 -e MYSQL_ROOT_PASSWORD=123456 -d mysql:latest

# 执行sql文件
mysql -h 127.0.0.1 -u root -p123456 < $(pwd)/../docs/sql/config.sql