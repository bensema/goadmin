package service

import (
	"github.com/bensema/goadmin/model"
	"github.com/bensema/goadmin/utils"
	"github.com/gin-gonic/gin"
	"image/gif"
	"library/image/captcha"
	"strings"
	"time"
)

func (s *Service) CaptchaGif(c *gin.Context) (*gif.GIF, string) {
	return s.captcha.GifCaptcha.Create(4, captcha.ALL)
}

func (s *Service) CaptchaImg(c *gin.Context) (*captcha.Image, string) {
	return s.captcha.ImgCaptcha.Create(4, captcha.ALL)
}

func (s *Service) SetCaptchaCache(c *gin.Context, code string) (string, error) {
	key := utils.RandomString(30)
	_key := strings.Join([]string{model.RedisPrefixCaptcha, key}, ":")
	err := s.SetCacheObj(c, _key, code, 60*time.Second)
	return key, err
}

func (s *Service) GetCaptchaCache(c *gin.Context, key string) (code string, err error) {
	_key := strings.Join([]string{model.RedisPrefixCaptcha, key}, ":")
	err = s.GetCacheObj(c, _key, &code)
	return
}

func (s *Service) DelCaptchaCache(c *gin.Context, key string) error {
	_key := strings.Join([]string{model.RedisPrefixCaptcha, key}, ":")
	return s.DelCache(c, _key)
}

func (s *Service) CaptchaVerify(c *gin.Context, key, code string) (bool, error) {
	_key := strings.Join([]string{model.RedisPrefixCaptcha, key}, ":")
	var planCode string
	err := s.GetCacheObj(c, _key, &planCode)
	if planCode == "" {
		return false, model.ErrNil
	}
	return strings.ToLower(planCode) == strings.ToLower(code), err
}
