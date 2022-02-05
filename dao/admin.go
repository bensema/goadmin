package dao

import (
	"github.com/bensema/gcurd"
	"github.com/bensema/goadmin/model"
	"github.com/gin-gonic/gin"
)

func (d *Dao) GetAdminById(c *gin.Context, id int) (*model.Admin, error) {
	obj := &model.Admin{}
	var wvs []*gcurd.WhereValue
	wvs = append(wvs, gcurd.EQ("id", id))
	return GetWhere(c, d.db, obj, wvs)
}

func (d *Dao) GetAdminByName(c *gin.Context, name string) (*model.Admin, error) {
	obj := &model.Admin{}
	var wvs []*gcurd.WhereValue
	wvs = append(wvs, gcurd.EQ("name", name))
	return GetWhere(c, d.db, obj, wvs)
}

func (d *Dao) GetAdminByAdminId(c *gin.Context, adminId int) (*model.Admin, error) {
	obj := &model.Admin{}
	var wvs []*gcurd.WhereValue
	wvs = append(wvs, gcurd.EQ("admin_id", adminId))
	return GetWhere(c, d.db, obj, wvs)
}

func (d *Dao) UpdateAdminByAdminId(c *gin.Context, adminId int, key string, value interface{}) error {
	obj := &model.Admin{}
	var wvs []*gcurd.WhereValue
	wvs = append(wvs, gcurd.EQ("admin_id", adminId))
	return UpdateWhere(c, d.db, obj, key, value, wvs)
}

func (d *Dao) DeleteAdminByAdminId(c *gin.Context, adminId int) error {
	obj := &model.Admin{}
	var wvs []*gcurd.WhereValue
	wvs = append(wvs, gcurd.EQ("admin_id", adminId))
	return gcurd.DeleteWhere(c, d.db, obj, wvs)
}

func (d *Dao) DeleteAdminRoleByAdminId(c *gin.Context, adminId int) error {
	obj := &model.AdminRole{}
	var wvs []*gcurd.WhereValue
	wvs = append(wvs, gcurd.EQ("admin_id", adminId))
	return gcurd.DeleteWhere(c, d.db, obj, wvs)
}

func (d *Dao) GetRoleByName(c *gin.Context, name string) (*model.Role, error) {
	obj := &model.Role{}
	var wvs []*gcurd.WhereValue
	wvs = append(wvs, gcurd.EQ("name", name))
	return GetWhere(c, d.db, obj, wvs)
}

func (d *Dao) DeleteRoleMenuByRoleId(c *gin.Context, roleId int) error {
	obj := &model.RoleMenu{}
	var wvs []*gcurd.WhereValue
	wvs = append(wvs, gcurd.EQ("role_id", roleId))
	return gcurd.DeleteWhere(c, d.db, obj, wvs)
}

func (d *Dao) DeleteRoleApiByRoleId(c *gin.Context, roleId int) error {
	obj := &model.RoleApi{}
	var wvs []*gcurd.WhereValue
	wvs = append(wvs, gcurd.EQ("role_id", roleId))
	return gcurd.DeleteWhere(c, d.db, obj, wvs)
}

func (d *Dao) GetMenuByName(c *gin.Context, name string) (*model.Menu, error) {
	obj := &model.Menu{}
	var wvs []*gcurd.WhereValue
	wvs = append(wvs, gcurd.EQ("name", name))
	return GetWhere(c, d.db, obj, wvs)
}

func (d *Dao) GetApiByName(c *gin.Context, name string) (*model.Api, error) {
	obj := &model.Api{}
	var wvs []*gcurd.WhereValue
	wvs = append(wvs, gcurd.EQ("name", name))
	return GetWhere(c, d.db, obj, wvs)
}

func (d *Dao) GetApiByUrl(c *gin.Context, url string) (*model.Api, error) {
	obj := &model.Api{}
	var wvs []*gcurd.WhereValue
	wvs = append(wvs, gcurd.EQ("url", url))
	return GetWhere(c, d.db, obj, wvs)
}

func (d *Dao) GetMenuByUrl(c *gin.Context, url string) (*model.Menu, error) {
	obj := &model.Menu{}
	var wvs []*gcurd.WhereValue
	wvs = append(wvs, gcurd.EQ("url", url))
	return GetWhere(c, d.db, obj, wvs)
}

func (d *Dao) GetRoleById(c *gin.Context, id int) (*model.Role, error) {
	obj := &model.Role{}
	var wvs []*gcurd.WhereValue
	wvs = append(wvs, gcurd.EQ("id", id))
	return GetWhere(c, d.db, obj, wvs)
}
