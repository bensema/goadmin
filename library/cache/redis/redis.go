package redis

import (
	"github.com/go-redis/redis/v8"
)

func New(conf *Config) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     conf.Addr,
		Password: conf.Password,
		DB:       conf.Db,
	})
	return client
}
