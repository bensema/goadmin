package dao

import (
	"github.com/bensema/goadmin/dao/internal"
	"github.com/bensema/goadmin/model"
	"github.com/gin-gonic/gin"
)

func (d *Dao) GetAdminByName(c *gin.Context, name string) (*model.Admin, error) {
	return internal.GetAdminByName(c, d.db, name)
}

func (d *Dao) DeleteAdminRoleByAdminId(c *gin.Context, adminId int) error {
	return internal.DeleteAdminRoleByAdminId(c, d.db, adminId)
}

func (d *Dao) DeleteRolePermissionByRoleId(c *gin.Context, id int) error {
	return internal.DeleteRolePermissionByRoleId(c, d.db, id)
}

func (d *Dao) GetRoleByName(c *gin.Context, name string) (*model.Role, error) {
	return internal.GetRoleByName(c, d.db, name)
}

func (d *Dao) GetPermissionByName(c *gin.Context, name string) (*model.Permission, error) {
	return internal.GetPermissionByName(c, d.db, name)
}

func (d *Dao) DeletePermissionMenuByPermissionId(c *gin.Context, id int) error {
	return internal.DeletePermissionMenuByPermissionId(c, d.db, id)
}

func (d *Dao) DeletePermissionOperationByPermissionId(c *gin.Context, id int) error {
	return internal.DeletePermissionOperationByPermissionId(c, d.db, id)
}

func (d *Dao) GetMenuByName(c *gin.Context, name string) (*model.Menu, error) {
	return internal.GetMenuByName(c, d.db, name)
}

func (d *Dao) GetOperationByName(c *gin.Context, name string) (*model.Operation, error) {
	return internal.GetOperationByName(c, d.db, name)
}
