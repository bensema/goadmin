package dao

import (
	"github.com/bensema/goadmin/dao/internal"
	"github.com/bensema/goadmin/model"
	"github.com/gin-gonic/gin"
)

func (d *Dao) CreateAdmin(c *gin.Context, m *model.Admin) error {
	return internal.CreateAdmin(c, d.db, m)
}

func (d *Dao) DeleteAdminById(c *gin.Context, id int) error {
	return internal.DeleteAdminById(c, d.db, id)
}

func (d *Dao) UpdateAdminById(c *gin.Context, id int, key string, value interface{}) error {
	return internal.UpdateAdminById(c, d.db, id, key, value)
}

func (d *Dao) GetAdminById(c *gin.Context, id int) (*model.Admin, error) {
	return internal.GetAdminById(c, d.db, id)
}

func (d *Dao) FindAdmin(c *gin.Context, req *model.FindAdminReq) ([]*model.Admin, error) {
	return internal.FindAdmin(c, d.db, req)
}

func (d *Dao) PageFindAdminTotal(c *gin.Context, req *model.FindAdminReq) (int, error) {
	return internal.PageFindAdminTotal(c, d.db, req)
}

func (d *Dao) PageFindAdmin(c *gin.Context, req *model.FindAdminReq) ([]*model.Admin, error) {
	return internal.PageFindAdmin(c, d.db, req)
}

func (d *Dao) CreateAdminRole(c *gin.Context, m *model.AdminRole) error {
	return internal.CreateAdminRole(c, d.db, m)
}

func (d *Dao) DeleteAdminRoleById(c *gin.Context, id int) error {
	return internal.DeleteAdminRoleById(c, d.db, id)
}

func (d *Dao) UpdateAdminRoleById(c *gin.Context, id int, key string, value interface{}) error {
	return internal.UpdateAdminRoleById(c, d.db, id, key, value)
}

func (d *Dao) GetAdminRoleById(c *gin.Context, id int) (*model.AdminRole, error) {
	return internal.GetAdminRoleById(c, d.db, id)
}

func (d *Dao) FindAdminRole(c *gin.Context, req *model.FindAdminRoleReq) ([]*model.AdminRole, error) {
	return internal.FindAdminRole(c, d.db, req)
}

func (d *Dao) PageFindAdminRoleTotal(c *gin.Context, req *model.FindAdminRoleReq) (int, error) {
	return internal.PageFindAdminRoleTotal(c, d.db, req)
}

func (d *Dao) PageFindAdminRole(c *gin.Context, req *model.FindAdminRoleReq) ([]*model.AdminRole, error) {
	return internal.PageFindAdminRole(c, d.db, req)
}

func (d *Dao) CreateLogAdminLogin(c *gin.Context, m *model.LogAdminLogin) error {
	return internal.CreateLogAdminLogin(c, d.db, m)
}

func (d *Dao) DeleteLogAdminLoginById(c *gin.Context, id int) error {
	return internal.DeleteLogAdminLoginById(c, d.db, id)
}

func (d *Dao) UpdateLogAdminLoginById(c *gin.Context, id int, key string, value interface{}) error {
	return internal.UpdateLogAdminLoginById(c, d.db, id, key, value)
}

func (d *Dao) GetLogAdminLoginById(c *gin.Context, id int) (*model.LogAdminLogin, error) {
	return internal.GetLogAdminLoginById(c, d.db, id)
}

func (d *Dao) FindLogAdminLogin(c *gin.Context, req *model.FindLogAdminLoginReq) ([]*model.LogAdminLogin, error) {
	return internal.FindLogAdminLogin(c, d.db, req)
}

func (d *Dao) PageFindLogAdminLoginTotal(c *gin.Context, req *model.FindLogAdminLoginReq) (int, error) {
	return internal.PageFindLogAdminLoginTotal(c, d.db, req)
}

func (d *Dao) PageFindLogAdminLogin(c *gin.Context, req *model.FindLogAdminLoginReq) ([]*model.LogAdminLogin, error) {
	return internal.PageFindLogAdminLogin(c, d.db, req)
}

func (d *Dao) CreateLogAdminOperation(c *gin.Context, m *model.LogAdminOperation) error {
	return internal.CreateLogAdminOperation(c, d.db, m)
}

func (d *Dao) DeleteLogAdminOperationById(c *gin.Context, id int) error {
	return internal.DeleteLogAdminOperationById(c, d.db, id)
}

func (d *Dao) UpdateLogAdminOperationById(c *gin.Context, id int, key string, value interface{}) error {
	return internal.UpdateLogAdminOperationById(c, d.db, id, key, value)
}

func (d *Dao) GetLogAdminOperationById(c *gin.Context, id int) (*model.LogAdminOperation, error) {
	return internal.GetLogAdminOperationById(c, d.db, id)
}

func (d *Dao) FindLogAdminOperation(c *gin.Context, req *model.FindLogAdminOperationReq) ([]*model.LogAdminOperation, error) {
	return internal.FindLogAdminOperation(c, d.db, req)
}

func (d *Dao) PageFindLogAdminOperationTotal(c *gin.Context, req *model.FindLogAdminOperationReq) (int, error) {
	return internal.PageFindLogAdminOperationTotal(c, d.db, req)
}

func (d *Dao) PageFindLogAdminOperation(c *gin.Context, req *model.FindLogAdminOperationReq) ([]*model.LogAdminOperation, error) {
	return internal.PageFindLogAdminOperation(c, d.db, req)
}

func (d *Dao) CreateMenu(c *gin.Context, m *model.Menu) error {
	return internal.CreateMenu(c, d.db, m)
}

func (d *Dao) DeleteMenuById(c *gin.Context, id int) error {
	return internal.DeleteMenuById(c, d.db, id)
}

func (d *Dao) UpdateMenuById(c *gin.Context, id int, key string, value interface{}) error {
	return internal.UpdateMenuById(c, d.db, id, key, value)
}

func (d *Dao) GetMenuById(c *gin.Context, id int) (*model.Menu, error) {
	return internal.GetMenuById(c, d.db, id)
}

func (d *Dao) FindMenu(c *gin.Context, req *model.FindMenuReq) ([]*model.Menu, error) {
	return internal.FindMenu(c, d.db, req)
}

func (d *Dao) PageFindMenuTotal(c *gin.Context, req *model.FindMenuReq) (int, error) {
	return internal.PageFindMenuTotal(c, d.db, req)
}

func (d *Dao) PageFindMenu(c *gin.Context, req *model.FindMenuReq) ([]*model.Menu, error) {
	return internal.PageFindMenu(c, d.db, req)
}

func (d *Dao) CreateOperation(c *gin.Context, m *model.Operation) error {
	return internal.CreateOperation(c, d.db, m)
}

func (d *Dao) DeleteOperationById(c *gin.Context, id int) error {
	return internal.DeleteOperationById(c, d.db, id)
}

func (d *Dao) UpdateOperationById(c *gin.Context, id int, key string, value interface{}) error {
	return internal.UpdateOperationById(c, d.db, id, key, value)
}

func (d *Dao) GetOperationById(c *gin.Context, id int) (*model.Operation, error) {
	return internal.GetOperationById(c, d.db, id)
}

func (d *Dao) FindOperation(c *gin.Context, req *model.FindOperationReq) ([]*model.Operation, error) {
	return internal.FindOperation(c, d.db, req)
}

func (d *Dao) PageFindOperationTotal(c *gin.Context, req *model.FindOperationReq) (int, error) {
	return internal.PageFindOperationTotal(c, d.db, req)
}

func (d *Dao) PageFindOperation(c *gin.Context, req *model.FindOperationReq) ([]*model.Operation, error) {
	return internal.PageFindOperation(c, d.db, req)
}

func (d *Dao) CreatePermission(c *gin.Context, m *model.Permission) error {
	return internal.CreatePermission(c, d.db, m)
}

func (d *Dao) DeletePermissionById(c *gin.Context, id int) error {
	return internal.DeletePermissionById(c, d.db, id)
}

func (d *Dao) UpdatePermissionById(c *gin.Context, id int, key string, value interface{}) error {
	return internal.UpdatePermissionById(c, d.db, id, key, value)
}

func (d *Dao) GetPermissionById(c *gin.Context, id int) (*model.Permission, error) {
	return internal.GetPermissionById(c, d.db, id)
}

func (d *Dao) FindPermission(c *gin.Context, req *model.FindPermissionReq) ([]*model.Permission, error) {
	return internal.FindPermission(c, d.db, req)
}

func (d *Dao) PageFindPermissionTotal(c *gin.Context, req *model.FindPermissionReq) (int, error) {
	return internal.PageFindPermissionTotal(c, d.db, req)
}

func (d *Dao) PageFindPermission(c *gin.Context, req *model.FindPermissionReq) ([]*model.Permission, error) {
	return internal.PageFindPermission(c, d.db, req)
}

func (d *Dao) CreatePermissionMenu(c *gin.Context, m *model.PermissionMenu) error {
	return internal.CreatePermissionMenu(c, d.db, m)
}

func (d *Dao) DeletePermissionMenuById(c *gin.Context, id int) error {
	return internal.DeletePermissionMenuById(c, d.db, id)
}

func (d *Dao) UpdatePermissionMenuById(c *gin.Context, id int, key string, value interface{}) error {
	return internal.UpdatePermissionMenuById(c, d.db, id, key, value)
}

func (d *Dao) GetPermissionMenuById(c *gin.Context, id int) (*model.PermissionMenu, error) {
	return internal.GetPermissionMenuById(c, d.db, id)
}

func (d *Dao) FindPermissionMenu(c *gin.Context, req *model.FindPermissionMenuReq) ([]*model.PermissionMenu, error) {
	return internal.FindPermissionMenu(c, d.db, req)
}

func (d *Dao) PageFindPermissionMenuTotal(c *gin.Context, req *model.FindPermissionMenuReq) (int, error) {
	return internal.PageFindPermissionMenuTotal(c, d.db, req)
}

func (d *Dao) PageFindPermissionMenu(c *gin.Context, req *model.FindPermissionMenuReq) ([]*model.PermissionMenu, error) {
	return internal.PageFindPermissionMenu(c, d.db, req)
}

func (d *Dao) CreatePermissionOperation(c *gin.Context, m *model.PermissionOperation) error {
	return internal.CreatePermissionOperation(c, d.db, m)
}

func (d *Dao) DeletePermissionOperationById(c *gin.Context, id int) error {
	return internal.DeletePermissionOperationById(c, d.db, id)
}

func (d *Dao) UpdatePermissionOperationById(c *gin.Context, id int, key string, value interface{}) error {
	return internal.UpdatePermissionOperationById(c, d.db, id, key, value)
}

func (d *Dao) GetPermissionOperationById(c *gin.Context, id int) (*model.PermissionOperation, error) {
	return internal.GetPermissionOperationById(c, d.db, id)
}

func (d *Dao) FindPermissionOperation(c *gin.Context, req *model.FindPermissionOperationReq) ([]*model.PermissionOperation, error) {
	return internal.FindPermissionOperation(c, d.db, req)
}

func (d *Dao) PageFindPermissionOperationTotal(c *gin.Context, req *model.FindPermissionOperationReq) (int, error) {
	return internal.PageFindPermissionOperationTotal(c, d.db, req)
}

func (d *Dao) PageFindPermissionOperation(c *gin.Context, req *model.FindPermissionOperationReq) ([]*model.PermissionOperation, error) {
	return internal.PageFindPermissionOperation(c, d.db, req)
}

func (d *Dao) CreateRole(c *gin.Context, m *model.Role) error {
	return internal.CreateRole(c, d.db, m)
}

func (d *Dao) DeleteRoleById(c *gin.Context, id int) error {
	return internal.DeleteRoleById(c, d.db, id)
}

func (d *Dao) UpdateRoleById(c *gin.Context, id int, key string, value interface{}) error {
	return internal.UpdateRoleById(c, d.db, id, key, value)
}

func (d *Dao) GetRoleById(c *gin.Context, id int) (*model.Role, error) {
	return internal.GetRoleById(c, d.db, id)
}

func (d *Dao) FindRole(c *gin.Context, req *model.FindRoleReq) ([]*model.Role, error) {
	return internal.FindRole(c, d.db, req)
}

func (d *Dao) PageFindRoleTotal(c *gin.Context, req *model.FindRoleReq) (int, error) {
	return internal.PageFindRoleTotal(c, d.db, req)
}

func (d *Dao) PageFindRole(c *gin.Context, req *model.FindRoleReq) ([]*model.Role, error) {
	return internal.PageFindRole(c, d.db, req)
}

func (d *Dao) CreateRolePermission(c *gin.Context, m *model.RolePermission) error {
	return internal.CreateRolePermission(c, d.db, m)
}

func (d *Dao) DeleteRolePermissionById(c *gin.Context, id int) error {
	return internal.DeleteRolePermissionById(c, d.db, id)
}

func (d *Dao) UpdateRolePermissionById(c *gin.Context, id int, key string, value interface{}) error {
	return internal.UpdateRolePermissionById(c, d.db, id, key, value)
}

func (d *Dao) GetRolePermissionById(c *gin.Context, id int) (*model.RolePermission, error) {
	return internal.GetRolePermissionById(c, d.db, id)
}

func (d *Dao) FindRolePermission(c *gin.Context, req *model.FindRolePermissionReq) ([]*model.RolePermission, error) {
	return internal.FindRolePermission(c, d.db, req)
}

func (d *Dao) PageFindRolePermissionTotal(c *gin.Context, req *model.FindRolePermissionReq) (int, error) {
	return internal.PageFindRolePermissionTotal(c, d.db, req)
}

func (d *Dao) PageFindRolePermission(c *gin.Context, req *model.FindRolePermissionReq) ([]*model.RolePermission, error) {
	return internal.PageFindRolePermission(c, d.db, req)
}
