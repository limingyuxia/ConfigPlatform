package redis

import (
	"ConfigPlatform/conf"
	"errors"
	"log"
	"strconv"

	"github.com/go-redis/redis"
)

var (
	RedisConn *redis.Client
)

func InitRedis() error {
	log.Print("redis: ", conf.RedisSetting.Host+":"+strconv.Itoa(conf.RedisSetting.Port))

	RedisConn = redis.NewClient(&redis.Options{
		Addr:     conf.RedisSetting.Host + ":" + strconv.Itoa(conf.RedisSetting.Port),
		Password: conf.DatabaseSetting.Password,
		DB:       0, // use default DB
	})

	if RedisConn == nil {
		log.Print("connect to redis failed!")
		return errors.New("connect to redis failed")
	}

	return nil
}
