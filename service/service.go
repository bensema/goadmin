package service

import (
	"github.com/bensema/goadmin/config"
	"github.com/bensema/goadmin/dao"
	"github.com/bensema/goadmin/model"
	"github.com/gin-gonic/gin"
	"library/image/captcha"
	xip "library/net/ip"
)

var Srv *Service

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

func (s *Service) getAdminFromContext(c *gin.Context) (*model.Admin, error) {
	sid, err := c.Cookie("admin-session")
	if err != nil {
		return nil, err
	}
	adminSession := &model.AdminSession{}

	err = s.GetAdminSessionCache(c, sid, adminSession)
	if err != nil {
		return nil, err
	}
	return s.dao.GetAdminByAdminId(c, adminSession.AdminId)
}

func (s *Service) GetAdminFromContext(c *gin.Context) (*model.Admin, error) {
	return s.getAdminFromContext(c)
}
