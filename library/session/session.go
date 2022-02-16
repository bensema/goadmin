package session

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	jsoniter "github.com/json-iterator/go"
	xRedis "library/cache/redis"
	"math/rand"
	"time"
)

const (
	JSESSIONID = "JSESSIONID"
	PREFIX     = "JSESSIONID"
)

var (
	json = jsoniter.ConfigCompatibleWithStandardLibrary
)

type Session struct {
	rdb *redis.Client
}

func New(c *xRedis.Config) (d *Session) {
	d = &Session{
		rdb: xRedis.New(c),
	}
	return
}

func GenerateSID() string {
	str := "0123456789abcdefghijklmnopqrstuvwxyz"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < 32; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}

func (s *Session) getSession(c context.Context, sid string, obj any) error {
	sid = PREFIX + ":" + sid
	objStr, err := s.rdb.Get(c, sid).Result()
	if err != nil {
		return err
	}
	return json.UnmarshalFromString(objStr, obj)
}

func (s *Session) setSession(c context.Context, sid string, obj any) error {
	sid = PREFIX + ":" + sid
	val, err := json.MarshalToString(obj)
	if err != nil {
		return err
	}
	return s.rdb.Set(c, sid, val, 0*60*24*time.Second).Err()

}

func (s *Session) deleteSession(c context.Context, sid string) error {
	sid = PREFIX + ":" + sid
	_, err := s.rdb.Del(c, sid).Result()
	return err
}

func (s *Session) GinSetSession(c *gin.Context, obj any) error {
	sid, _ := c.Cookie(JSESSIONID)
	_ = s.deleteSession(c, sid)
	sid = GenerateSID()
	c.SetCookie(JSESSIONID, sid, 60*60*8, "/", "", false, true)
	return s.setSession(c, sid, obj)
}

func (s *Session) GinRefreshSession(c *gin.Context) error {
	sid, _ := c.Cookie(JSESSIONID)
	sid = PREFIX + ":" + sid
	objStr, err := s.rdb.Get(c, sid).Result()
	if err != nil {
		return err
	}
	sid = GenerateSID()

	c.SetCookie(JSESSIONID, sid, 60*60*8, "/", "", false, true)
	return s.rdb.Set(c, PREFIX+":"+sid, objStr, 0*60*24*time.Second).Err()
}

func (s *Session) GinLoadSession(c *gin.Context, obj any) error {
	sid, err := c.Cookie(JSESSIONID)
	if err != nil {
		return err
	}

	return s.getSession(c, sid, obj)
}

func (s *Session) GinClearSession(c *gin.Context) error {
	sid, err := c.Cookie(JSESSIONID)
	if err != nil {
		return err
	}
	err = s.deleteSession(c, sid)
	c.SetCookie(JSESSIONID, "", -1, "/", "", false, true)

	return err
}
