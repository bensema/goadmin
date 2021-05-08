package dao

import (
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"time"
)

func (d *Dao) RDB() *redis.Client {
	return d.rdb
}

func (d *Dao) Set(c *gin.Context, key string, value string, expiration time.Duration) (err error) {
	err = d.rdb.Set(c, key, value, expiration).Err()
	if err != nil {
		return err
	}
	return
}

func (d *Dao) Get(c *gin.Context, key string) (data string, err error) {
	return d.rdb.Get(c, key).Result()
}

func (d *Dao) Del(c *gin.Context, key string) (int64, error) {
	return d.rdb.Del(c, key).Result()
}

func (d *Dao) Pub(c *gin.Context, key string, value string) (err error) {
	return d.rdb.Publish(c, key, value).Err()
}
