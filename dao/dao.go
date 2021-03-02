package dao

import (
	"context"
	"database/sql"
	"github.com/bensema/goadmin/config"
	"github.com/bensema/library/cache/redis"
	bSql "github.com/bensema/library/database/sql"
	"time"
)

type Dao struct {
	c *config.Config

	db *sql.DB

	rds *redis.Pool

	adminSessionExpire int64
}

func New(c *config.Config) (d *Dao) {
	d = &Dao{
		c:                  c,
		db:                 bSql.New(c.MySQL),
		rds:                redis.NewPool(c.Redis.Config),
		adminSessionExpire: int64(time.Duration(c.Redis.AdminSessionExpire) / time.Second),
	}
	return
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
}
