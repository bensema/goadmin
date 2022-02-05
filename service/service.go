package service

import (
	"fmt"
	"github.com/bensema/goadmin/config"
	"github.com/bensema/goadmin/dao"
	"github.com/bensema/goadmin/model"
	"github.com/gin-gonic/gin"
	"library/image/captcha"
	xip "library/net/ip"
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

func GetAdminFromContext(c *gin.Context) (*model.Admin, error) {
	id, _ := c.Get("admin_id")
	name, _ := c.Get("admin_name")
	fmt.Println(id, name)
	var a = model.Admin{}
	a.Id = id.(int)
	a.Name = name.(string)
	return &a, nil
}
