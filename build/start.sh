#!/bin/bash

# 这是服务启动脚本

if [ $# -lt 2 ]; then
    echo "usage [ mysql | redis | backend | frontend ] [ tag ]"
    exit 1
fi

# 前后端服务所使用配置项
HttpPort=0
HttpsPort=0
ServerPort=0
MysqlHost=""
MysqlPort=0
MysqlPwd=""

function readIni()
{ 
    ConfigFile=$(pwd)/../conf/conf.ini
    if [ ! -e $ConfigFile ]; then
        echo "配置文件不存在"
        exit 1
    fi

    Section="nginx"
    Key="HttpsPort"
    HttpsPort=`awk -F '=' '/\['$Section'\]/{a=1}a==1&&$1~/'$Key'/{print $2;exit}' $ConfigFile`

    Key="HttpPort"
    HttpPort=`awk -F '=' '/\['$Section'\]/{a=1}a==1&&$1~/'$Key'/{print $2;exit}' $ConfigFile`

    Section="server"
    Key="Port"
    ServerPort=`awk -F '=' '/\['$Section'\]/{a=1}a==1&&$1~/'$Key'/{print $2;exit}' $ConfigFile`

    Section="database"
    Key="Host"
    MysqlHost=`awk -F '=' '/\['$Section'\]/{a=1}a==1&&$1~/'$Key'/{print $2;exit}' $ConfigFile`

    Key="Port"
    MysqlPort=`awk -F '=' '/\['$Section'\]/{a=1}a==1&&$1~/'$Key'/{print $2;exit}' $ConfigFile`

    Key="Password"
    MysqlPwd=`awk -F '=' '/\['$Section'\]/{a=1}a==1&&$1~/'$Key'/{print $2;exit}' $ConfigFile`
}

readIni

if [ $HttpPort = 0 ] || [ $HttpsPort = 0 ] || [ $ServerPort = 0 ] \
    || [ $MysqlPort = 0 ] || [ $MysqlPwd = "" ] || [ $MysqlHost == "" ]; then
    echo "配置文件解析出错"
    exit 1
fi

BaseServer=$1
ServerTag=$2
echo "run server, based on "$BaseServer" tag "$ServerTag

# 有的参数暂时不能用shell变量代替

if [ $BaseServer = "mysql" ]; then
    # 构建数据库
    docker run -d --rm -p $MysqlPort:3306 --name config-db --env MYSQL_ROOT_PASSWORD=123456 mysql:$ServerTag
    # 等待数据库初始化
    sleep 30
    # 初始化table
    mysql -h $MysqlHost -u root -p123456 < $(pwd)/../docs/sql/config.sql
elif [$BaseServer = "redis"]; then 
    docker run -d --rm -p 6379:6379 -v $(pwd)/../conf/redis/redis.conf:/etc/redis/redis.conf --name config-redis redis:$ServerTag --requirepass 123456
elif [ $BaseServer = "backend" ]; then
    # 运行服务端容器
    docker run -d --rm -p $ServerPort:8000 -p 8001:8001 --name config-server 18509518245/config-server:$ServerTag
elif [ $BaseServer = "frontend" ]; then
    # 运行前端容器
    container_id=$(docker create 18509518245/config-web:$ServerTag)
    docker cp $container_id:/home/dist $(pwd)/../web
    docker rm $container_id

    # -d 后台运行 --rm 如果有同名的容器，自动删除
    docker run -d --rm -p $HttpPort:9528 -p $HttpsPort:9529 -v $(pwd)/../web/dist:/usr/share/nginx/html:ro \
        --name config-web 18509518245/config-nginx:$ServerTag
else 
    echo "参数错误"
fi
