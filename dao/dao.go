package dao

import (
	"context"
	"database/sql"
	"github.com/bensema/goadmin/config"
	"github.com/go-redis/redis/v8"
	xRedis "library/cache/redis"
	xSql "library/database/sql"
)

type Dao struct {
	c   *config.Config
	db  *sql.DB
	rdb *redis.Client
}

func New(c *config.Config) (d *Dao) {
	d = &Dao{
		c:   c,
		db:  xSql.New(c.MySQL),
		rdb: xRedis.New(c.Redis),
	}
	return
}

func (d *Dao) DB() *sql.DB {
	return d.db
}

func (d *Dao) Ping(c context.Context) (err error) {
	err = d.db.Ping()
	return
}

// Close close the resource.
func (d *Dao) Close() {
	if d.db != nil {
		_ = d.db.Close()
	}
	if d.rdb != nil {
		_ = d.rdb.Close()
	}
}
