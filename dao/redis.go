package dao

import (
	"encoding/json"
	"github.com/bensema/goadmin/model"
	"github.com/bensema/library/cache/redis"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

func (d *Dao) SetCaptchaCache(c *gin.Context, key string, code string) (err error) {
	conn, err := d.rds.GetContext(c.Request.Context())
	if err != nil {
		return
	}
	defer conn.Close()
	if err = conn.Send("SET", key, code); err != nil {
		return
	}
	if err = conn.Send("EXPIRE", key, 60*5); err != nil {
		return
	}

	if err != nil {
		return
	}
	err = conn.Flush()
	if err != nil {
		return
	}
	return
}

func (d *Dao) GetCaptchaCache(c *gin.Context, key string) (code string, err error) {
	conn, err := d.rds.GetContext(c.Request.Context())
	if err != nil {
		return
	}
	defer conn.Close()
	code, err = redis.String(conn.Do("GET", key))
	return
}

func (d *Dao) DelCaptchaCache(c *gin.Context, key string) (err error) {
	conn, err := d.rds.GetContext(c.Request.Context())
	if err != nil {
		return
	}
	defer conn.Close()
	_, err = conn.Do("DEL", key)
	return
}

func (d *Dao) SetAdminSessionCache(c *gin.Context, key string, adminSession *model.AdminSession) (err error) {
	var (
		bs []byte
	)
	conn, err := d.rds.GetContext(c.Request.Context())
	if err != nil {
		return
	}
	defer conn.Close()
	if bs, err = json.Marshal(adminSession); err != nil {
		return errors.WithStack(err)
	}
	if err = conn.Send("SET", key, bs); err != nil {
		return
	}
	if err = conn.Send("EXPIRE", key, d.adminSessionExpire); err != nil {
		return
	}
	if err = conn.Flush(); err != nil {
		return
	}
	if _, err = conn.Receive(); err != nil {
		return
	}
	return
}

func (d *Dao) GetAdminSessionCache(c *gin.Context, key string) (adminSession *model.AdminSession, err error) {
	var (
		data []byte
	)
	conn, err := d.rds.GetContext(c.Request.Context())
	if err != nil {
		return
	}
	defer conn.Close()
	if data, err = redis.Bytes(conn.Do("GET", key)); err != nil {
		if err == redis.ErrNil {
			err = nil
		}
		return
	}
	adminSession = new(model.AdminSession)
	if err = json.Unmarshal(data, &adminSession); err != nil {
		err = errors.WithStack(err)
		return nil, err
	}
	return
}
