package dao

import (
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"time"
)

func (d *Dao) RDB() *redis.Client {
	return d.rdb
}

// RSet redis set
func (d *Dao) RSet(c *gin.Context, key string, value string, expiration time.Duration) (err error) {
	err = d.rdb.Set(c, key, value, expiration).Err()
	if err != nil {
		return err
	}
	return
}

// RGet redis get
func (d *Dao) RGet(c *gin.Context, key string) (data string, err error) {
	return d.rdb.Get(c, key).Result()
}

// RDelete redis delete
func (d *Dao) RDelete(c *gin.Context, key string) (int64, error) {
	return d.rdb.Del(c, key).Result()
}

// RPub redis publish
func (d *Dao) RPub(c *gin.Context, key string, value string) (err error) {
	return d.rdb.Publish(c, key, value).Err()
}
