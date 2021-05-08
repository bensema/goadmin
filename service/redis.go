package service

import (
	"github.com/bensema/goadmin/model"
	"github.com/gin-gonic/gin"
	jsoniter "github.com/json-iterator/go"
	"strings"
	"time"
)

var (
	json = jsoniter.ConfigCompatibleWithStandardLibrary
)

func (s *Service) SetCacheObj(c *gin.Context, key string, obj interface{}, expiration time.Duration) error {
	val, err := json.MarshalToString(obj)
	if err != nil {
		return err
	}
	return s.dao.Set(c, key, val, expiration)
}

func (s *Service) GetCacheObj(c *gin.Context, key string, obj interface{}) error {
	objStr, err := s.dao.Get(c, key)
	if err != nil {
		return err
	}
	return json.UnmarshalFromString(objStr, obj)
}

func (s *Service) DelCache(c *gin.Context, key string) error {
	_, err := s.dao.Del(c, key)
	return err
}

func (s *Service) GetAdminSessionCache(c *gin.Context, key string, obj interface{}) error {
	_key := strings.Join([]string{model.RedisPrefixAdminSession, key}, ":")
	return s.GetCacheObj(c, _key, obj)
}

func (s *Service) DeleteAdminSessionCache(c *gin.Context, key string) error {
	_key := strings.Join([]string{model.RedisPrefixAdminSession, key}, ":")
	return s.DelCache(c, _key)
}

func (s *Service) SetAdminSessionCache(c *gin.Context, key string, adminSession *model.AdminSession) (err error) {
	_key := strings.Join([]string{model.RedisPrefixAdminSession, key}, ":")
	return s.SetCacheObj(c, _key, adminSession, 60*60*24*time.Second)
}
