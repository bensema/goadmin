package service

import (
	"github.com/bensema/gcurd"
	"github.com/bensema/goadmin/dao"
	"github.com/bensema/goadmin/model"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"library/ecode"
)

func (s *Service) PermitAPI(c *gin.Context, adminId int) error {
	url := c.FullPath()
	api, err := s.dao.GetApiByUrl(c, url)
	if err != nil {
		log.Errorf(err.Error())
		return ecode.AccessDenied
	}
	var wvs []*gcurd.WhereValue
	wvs = append(wvs, gcurd.EQ("admin_id", adminId))
	obj := &model.AdminRole{}
	admRoles, err := dao.Find(c, s.dao.DB(), obj, wvs, obj.New)
	if err != nil {
		return ecode.AccessDenied
	}

	for _, admRole := range admRoles {
		var wvs []*gcurd.WhereValue
		wvs = append(wvs, gcurd.EQ("role_id", admRole.RoleId))
		wvs = append(wvs, gcurd.EQ("api_id", api.Id))
		obj := &model.RoleApi{}
		var roleApis []*model.RoleApi
		roleApis, err = dao.Find(c, s.dao.DB(), obj, wvs, obj.New)
		if err != nil {
			continue
		}
		if len(roleApis) > 0 {
			return nil
		}
	}

	return ecode.AccessDenied
}

func (s *Service) PermitWeb(c *gin.Context, adminId int) error {
	url := c.FullPath()
	menu, err := s.dao.GetMenuByUrl(c, url)
	if err != nil {
		log.Errorf(err.Error())
		return ecode.AccessDenied
	}

	var farr []*gcurd.WhereValue
	farr = append(farr, gcurd.EQ("admin_id", adminId))

	obj := &model.AdminRole{}
	admRoles, err := dao.Find(c, s.dao.DB(), obj, farr, obj.New)
	if err != nil {
		log.Errorf(err.Error())
		return ecode.AccessDenied
	}
	for _, admRole := range admRoles {
		var wvs []*gcurd.WhereValue
		wvs = append(wvs, gcurd.EQ("role_id", admRole.RoleId))
		wvs = append(wvs, gcurd.EQ("menu_id", menu.Id))
		obj := &model.RoleMenu{}
		var roleMenus []*model.RoleMenu
		roleMenus, err = dao.Find(c, s.dao.DB(), obj, wvs, obj.New)
		if err != nil {
			continue
		}
		if len(roleMenus) > 0 {
			return nil
		}
	}

	return ecode.AccessDenied
}
