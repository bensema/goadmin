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
