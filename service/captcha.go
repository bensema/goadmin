package service

import (
	"github.com/bensema/goadmin/utils"
	"github.com/bensema/library/image/captcha"
	"github.com/gin-gonic/gin"
	"image/gif"
	"strings"
)

func (s *Service) CaptchaGif(c *gin.Context) (*gif.GIF, string) {
	return s.captcha.GifCaptcha.Create(4, captcha.ALL)
}

func (s *Service) CaptchaImg(c *gin.Context) (*captcha.Image, string) {
	return s.captcha.ImgCaptcha.Create(4, captcha.ALL)
}

func (s *Service) SetCaptchaCache(c *gin.Context, code string) (string, error) {
	key := utils.RandomString(30)
	err := s.dao.SetCaptchaCache(c, key, code)
	return key, err
}

func (s *Service) CaptchaVerify(c *gin.Context, key, code string) (bool, error) {
	planCode, err := s.dao.GetCaptchaCache(c, key)
	return strings.ToLower(planCode) == strings.ToLower(code), err
}
