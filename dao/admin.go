package dao

import (
	"github.com/bensema/goadmin/dao/internal"
	"github.com/bensema/goadmin/model"
	"github.com/gin-gonic/gin"
)

func (d *Dao) GetAdminById(c *gin.Context, id int) (au *model.Admin, err error) {
	return internal.GetAdminById(c, d.db, id)
}

func (d *Dao) GetAdminByName(c *gin.Context, name string) (au *model.Admin, err error) {
	return internal.GetAdminByName(c, d.db, name)
}

func (d *Dao) GetMenuById(c *gin.Context, id int) (au *model.Menu, err error) {
	return internal.GetMenuById(c, d.db, id)
}

func (d *Dao) LogAdminLogin(c *gin.Context, m *model.LogAdminLogin) error {
	return internal.CreateLogAdminLogin(c, d.db, m)
}

func (d *Dao) FindAdmin(c *gin.Context, req *model.FindAdminReq) (objs []*model.Admin, err error) {
	return internal.FindAdmin(c, d.db, req)
}

func (d *Dao) FindAdminRole(c *gin.Context, req *model.FindAdminRoleReq) (objs []*model.AdminRole, err error) {
	return internal.FindAdminRole(c, d.db, req)
}

func (d *Dao) FindRole(c *gin.Context, req *model.FindRoleReq) (objs []*model.Role, err error) {
	return internal.FindRole(c, d.db, req)
}

func (d *Dao) FindRolePermission(c *gin.Context, req *model.FindRolePermissionReq) (objs []*model.RolePermission, err error) {
	return internal.FindRolePermission(c, d.db, req)
}

func (d *Dao) FindPermission(c *gin.Context, req *model.FindPermissionReq) (objs []*model.Permission, err error) {
	return internal.FindPermission(c, d.db, req)
}

func (d *Dao) FindPermissionMenu(c *gin.Context, req *model.FindPermissionMenuReq) (objs []*model.PermissionMenu, err error) {
	return internal.FindPermissionMenu(c, d.db, req)
}

func (d *Dao) FindMenu(c *gin.Context, req *model.FindMenuReq) (objs []*model.Menu, err error) {
	return internal.FindMenu(c, d.db, req)
}

func (d *Dao) FindPermissionOperation(c *gin.Context, req *model.FindPermissionOperationReq) (objs []*model.PermissionOperation, err error) {
	return internal.FindPermissionOperation(c, d.db, req)
}

func (d *Dao) FindOperation(c *gin.Context, req *model.FindOperationReq) (objs []*model.Operation, err error) {
	return internal.FindOperation(c, d.db, req)
}
