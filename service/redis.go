package service

import (
	"github.com/gin-gonic/gin"
	jsoniter "github.com/json-iterator/go"
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
	return s.dao.RSet(c, key, val, expiration)
}

func (s *Service) GetCacheObj(c *gin.Context, key string, obj interface{}) error {
	objStr, err := s.dao.RGet(c, key)
	if err != nil {
		return err
	}
	return json.UnmarshalFromString(objStr, obj)
}

func (s *Service) DelCache(c *gin.Context, key string) error {
	_, err := s.dao.RDelete(c, key)
	return err
}
