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
		obj := &model.RolePermission{}
		var rolePermissions []*model.RolePermission
		rolePermissions, err = dao.Find(c, s.dao.DB(), obj, wvs, obj.New)
		if err != nil {
			continue
		}
		for _, rolePermission := range rolePermissions {
			var wvs []*gcurd.WhereValue
			wvs = append(wvs, gcurd.EQ("permission_id", rolePermission.PermissionId))
			wvs = append(wvs, gcurd.EQ("api_id", api.Id))
			obj := &model.PermissionApi{}
			var permissionApis []*model.PermissionApi
			permissionApis, err = dao.Find(c, s.dao.DB(), obj, wvs, obj.New)
			if err != nil {
				continue
			}
			if len(permissionApis) > 0 {
				return nil
			}

		}

	}

	return ecode.AccessDenied
}
