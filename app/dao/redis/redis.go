package redis

import (
	"app/configs"
	"fmt"

	"github.com/go-redis/redis"
)

var rdb *redis.Client

func Init(config *configs.RedisConf) (err error) {
	rdb = redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%d",
			config.Host,
			config.Port,
		),
		Password: config.Password,
		DB:       config.Db,
		PoolSize: config.PollSize,
	})

	_, err = rdb.Ping().Result()
	return
}

func Close() {
	_ = rdb.Close()
}
