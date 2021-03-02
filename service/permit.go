package service

import (
	"github.com/bensema/goadmin/model"
	"github.com/bensema/library/ecode"
	"github.com/bensema/library/log"
	"github.com/gin-gonic/gin"
)

func (s *Service) PermitAPI(c *gin.Context, uid int) error {
	url := c.FullPath()
	operation, err := s.dao.GetOperationByUrl(c, url)
	if err != nil {
		log.Errorf(err.Error())
		return ecode.AccessDenied
	}
	farr := &model.FindAdminRoleReq{}
	farr.AdminId = uid
	admRoles, err := s.dao.FindAdminRole(c, farr)
	if err != nil {
		log.Errorf(err.Error())
		return ecode.AccessDenied
	}

	for _, admRole := range admRoles {
		frpr := &model.FindRolePermissionReq{}
		frpr.RoleId = admRole.RoleId
		rolePermissions, err := s.dao.FindRolePermission(c, frpr)
		if err != nil {
			continue
		}
		for _, rolePermission := range rolePermissions {
			fpor := &model.FindPermissionOperationReq{}
			fpor.PermissionId = rolePermission.PermissionId
			fpor.OperationId = operation.Id
			po, err := s.dao.FindPermissionOperation(c, fpor)
			if err != nil {
				continue
			}
			if len(po) > 0 {
				return nil
			}
		}
	}

	return ecode.AccessDenied
}

func (s *Service) PermitWeb(c *gin.Context, uid int) error {
	url := c.FullPath()
	menu, err := s.dao.GetMenuByUrl(c, url)
	if err != nil {
		log.Errorf(err.Error())
		return ecode.AccessDenied
	}
	farr := &model.FindAdminRoleReq{}
	farr.AdminId = uid
	admRoles, err := s.dao.FindAdminRole(c, farr)
	if err != nil {
		log.Errorf(err.Error())
		return ecode.AccessDenied
	}

	for _, admRole := range admRoles {
		frpr := &model.FindRolePermissionReq{}
		frpr.RoleId = admRole.RoleId
		rolePermissions, err := s.dao.FindRolePermission(c, frpr)
		if err != nil {
			continue
		}
		for _, rolePermission := range rolePermissions {
			fpor := &model.FindPermissionMenuReq{}
			fpor.PermissionId = rolePermission.PermissionId
			fpor.MenuId = menu.Id
			po, err := s.dao.FindPermissionMenu(c, fpor)
			if err != nil {
				continue
			}
			if len(po) > 0 {
				return nil
			}
		}
	}

	return ecode.AccessDenied
}
