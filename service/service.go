package service

import (
	"github.com/bensema/goadmin/config"
	"github.com/bensema/goadmin/dao"
	"github.com/bensema/library/image/captcha"
	xip "github.com/bensema/library/net/ip"
)

type Service struct {
	conf      *config.Config
	dao       *dao.Dao
	captcha   *captcha.Captcha
	Ip2Region *xip.Ip2Region
}

func New(c *config.Config) (s *Service) {
	ip2Region, _ := xip.New(c.Ip2Region.Path)
	s = &Service{
		conf:      c,
		dao:       dao.New(c),
		captcha:   captcha.New(c.Captcha),
		Ip2Region: ip2Region,
	}
	return
}

// Close close all dao.
func (s *Service) Close() {
	// log.Info("Close Dao physically!")
	s.dao.Close()
	// log.Info("Service Closed!")
}
