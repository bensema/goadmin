package dao

import (
	"database/sql"
	"github.com/bensema/gcurd"
	"github.com/bensema/goadmin/model"
	"github.com/gin-gonic/gin"
)

func (d *Dao) CreateAdmin(c *gin.Context, obj *model.Admin) (sql.Result, error) {
	return Create(c, d.db, obj)
}

func (d *Dao) DeleteAdmin(c *gin.Context, id int) error {
	obj := &model.Admin{}
	return Delete(c, d.db, obj, id)
}

func (d *Dao) UpdateAdmin(c *gin.Context, id int, key string, value interface{}) error {
	obj := &model.Admin{}
	return Update(c, d.db, obj, id, key, value)
}

func (d *Dao) UpdateWhereAdmin(c *gin.Context, key string, value interface{}, wvs []*gcurd.WhereValue) error {
	obj := &model.Admin{}
	return UpdateWhere(c, d.db, obj, key, value, wvs)
}

func (d *Dao) UpdateWhereKVAdmin(c *gin.Context, kvs []gcurd.KeyValue, wvs []*gcurd.WhereValue) error {
	obj := &model.Admin{}
	return UpdateWhereKV(c, d.db, obj, kvs, wvs)
}

func (d *Dao) GetAdmin(c *gin.Context, id int) (*model.Admin, error) {
	obj := &model.Admin{}
	return Get(c, d.db, obj, id)
}

func (d *Dao) GetWhereAdmin(c *gin.Context, wvs []*gcurd.WhereValue) (*model.Admin, error) {
	obj := &model.Admin{}
	return GetWhere(c, d.db, obj, wvs)
}

func (d *Dao) FindAdmin(c *gin.Context, wvs []*gcurd.WhereValue) ([]*model.Admin, error) {
	obj := &model.Admin{}
	return Find(c, d.db, obj, wvs, obj.New)
}

func (d *Dao) PageTotalAdmin(c *gin.Context, req *gcurd.Request) (int, error) {
	obj := &model.Admin{}
	return PageTotal(c, d.db, obj, req)
}

func (d *Dao) PageFindAdmin(c *gin.Context, req *gcurd.Request) ([]*model.Admin, error) {
	obj := &model.Admin{}
	return PageFind(c, d.db, obj, req, obj.New)
}

func (d *Dao) CreateAdminRole(c *gin.Context, obj *model.AdminRole) (sql.Result, error) {
	return Create(c, d.db, obj)
}

func (d *Dao) DeleteAdminRole(c *gin.Context, id int) error {
	obj := &model.AdminRole{}
	return Delete(c, d.db, obj, id)
}

func (d *Dao) UpdateAdminRole(c *gin.Context, id int, key string, value interface{}) error {
	obj := &model.AdminRole{}
	return Update(c, d.db, obj, id, key, value)
}

func (d *Dao) UpdateWhereAdminRole(c *gin.Context, key string, value interface{}, wvs []*gcurd.WhereValue) error {
	obj := &model.AdminRole{}
	return UpdateWhere(c, d.db, obj, key, value, wvs)
}

func (d *Dao) UpdateWhereKVAdminRole(c *gin.Context, kvs []gcurd.KeyValue, wvs []*gcurd.WhereValue) error {
	obj := &model.AdminRole{}
	return UpdateWhereKV(c, d.db, obj, kvs, wvs)
}

func (d *Dao) GetAdminRole(c *gin.Context, id int) (*model.AdminRole, error) {
	obj := &model.AdminRole{}
	return Get(c, d.db, obj, id)
}

func (d *Dao) GetWhereAdminRole(c *gin.Context, wvs []*gcurd.WhereValue) (*model.AdminRole, error) {
	obj := &model.AdminRole{}
	return GetWhere(c, d.db, obj, wvs)
}

func (d *Dao) FindAdminRole(c *gin.Context, wvs []*gcurd.WhereValue) ([]*model.AdminRole, error) {
	obj := &model.AdminRole{}
	return Find(c, d.db, obj, wvs, obj.New)
}

func (d *Dao) PageTotalAdminRole(c *gin.Context, req *gcurd.Request) (int, error) {
	obj := &model.AdminRole{}
	return PageTotal(c, d.db, obj, req)
}

func (d *Dao) PageFindAdminRole(c *gin.Context, req *gcurd.Request) ([]*model.AdminRole, error) {
	obj := &model.AdminRole{}
	return PageFind(c, d.db, obj, req, obj.New)
}

func (d *Dao) CreateApi(c *gin.Context, obj *model.Api) (sql.Result, error) {
	return Create(c, d.db, obj)
}

func (d *Dao) DeleteApi(c *gin.Context, id int) error {
	obj := &model.Api{}
	return Delete(c, d.db, obj, id)
}

func (d *Dao) UpdateApi(c *gin.Context, id int, key string, value interface{}) error {
	obj := &model.Api{}
	return Update(c, d.db, obj, id, key, value)
}

func (d *Dao) UpdateWhereApi(c *gin.Context, key string, value interface{}, wvs []*gcurd.WhereValue) error {
	obj := &model.Api{}
	return UpdateWhere(c, d.db, obj, key, value, wvs)
}

func (d *Dao) UpdateWhereKVApi(c *gin.Context, kvs []gcurd.KeyValue, wvs []*gcurd.WhereValue) error {
	obj := &model.Api{}
	return UpdateWhereKV(c, d.db, obj, kvs, wvs)
}

func (d *Dao) GetApi(c *gin.Context, id int) (*model.Api, error) {
	obj := &model.Api{}
	return Get(c, d.db, obj, id)
}

func (d *Dao) GetWhereApi(c *gin.Context, wvs []*gcurd.WhereValue) (*model.Api, error) {
	obj := &model.Api{}
	return GetWhere(c, d.db, obj, wvs)
}

func (d *Dao) FindApi(c *gin.Context, wvs []*gcurd.WhereValue) ([]*model.Api, error) {
	obj := &model.Api{}
	return Find(c, d.db, obj, wvs, obj.New)
}

func (d *Dao) PageTotalApi(c *gin.Context, req *gcurd.Request) (int, error) {
	obj := &model.Api{}
	return PageTotal(c, d.db, obj, req)
}

func (d *Dao) PageFindApi(c *gin.Context, req *gcurd.Request) ([]*model.Api, error) {
	obj := &model.Api{}
	return PageFind(c, d.db, obj, req, obj.New)
}

func (d *Dao) CreateLogAdminLogin(c *gin.Context, obj *model.LogAdminLogin) (sql.Result, error) {
	return Create(c, d.db, obj)
}

func (d *Dao) DeleteLogAdminLogin(c *gin.Context, id int) error {
	obj := &model.LogAdminLogin{}
	return Delete(c, d.db, obj, id)
}

func (d *Dao) UpdateLogAdminLogin(c *gin.Context, id int, key string, value interface{}) error {
	obj := &model.LogAdminLogin{}
	return Update(c, d.db, obj, id, key, value)
}

func (d *Dao) UpdateWhereLogAdminLogin(c *gin.Context, key string, value interface{}, wvs []*gcurd.WhereValue) error {
	obj := &model.LogAdminLogin{}
	return UpdateWhere(c, d.db, obj, key, value, wvs)
}

func (d *Dao) UpdateWhereKVLogAdminLogin(c *gin.Context, kvs []gcurd.KeyValue, wvs []*gcurd.WhereValue) error {
	obj := &model.LogAdminLogin{}
	return UpdateWhereKV(c, d.db, obj, kvs, wvs)
}

func (d *Dao) GetLogAdminLogin(c *gin.Context, id int) (*model.LogAdminLogin, error) {
	obj := &model.LogAdminLogin{}
	return Get(c, d.db, obj, id)
}

func (d *Dao) GetWhereLogAdminLogin(c *gin.Context, wvs []*gcurd.WhereValue) (*model.LogAdminLogin, error) {
	obj := &model.LogAdminLogin{}
	return GetWhere(c, d.db, obj, wvs)
}

func (d *Dao) FindLogAdminLogin(c *gin.Context, wvs []*gcurd.WhereValue) ([]*model.LogAdminLogin, error) {
	obj := &model.LogAdminLogin{}
	return Find(c, d.db, obj, wvs, obj.New)
}

func (d *Dao) PageTotalLogAdminLogin(c *gin.Context, req *gcurd.Request) (int, error) {
	obj := &model.LogAdminLogin{}
	return PageTotal(c, d.db, obj, req)
}

func (d *Dao) PageFindLogAdminLogin(c *gin.Context, req *gcurd.Request) ([]*model.LogAdminLogin, error) {
	obj := &model.LogAdminLogin{}
	return PageFind(c, d.db, obj, req, obj.New)
}

func (d *Dao) CreateLogAdminOperation(c *gin.Context, obj *model.LogAdminOperation) (sql.Result, error) {
	return Create(c, d.db, obj)
}

func (d *Dao) DeleteLogAdminOperation(c *gin.Context, id int) error {
	obj := &model.LogAdminOperation{}
	return Delete(c, d.db, obj, id)
}

func (d *Dao) UpdateLogAdminOperation(c *gin.Context, id int, key string, value interface{}) error {
	obj := &model.LogAdminOperation{}
	return Update(c, d.db, obj, id, key, value)
}

func (d *Dao) UpdateWhereLogAdminOperation(c *gin.Context, key string, value interface{}, wvs []*gcurd.WhereValue) error {
	obj := &model.LogAdminOperation{}
	return UpdateWhere(c, d.db, obj, key, value, wvs)
}

func (d *Dao) UpdateWhereKVLogAdminOperation(c *gin.Context, kvs []gcurd.KeyValue, wvs []*gcurd.WhereValue) error {
	obj := &model.LogAdminOperation{}
	return UpdateWhereKV(c, d.db, obj, kvs, wvs)
}

func (d *Dao) GetLogAdminOperation(c *gin.Context, id int) (*model.LogAdminOperation, error) {
	obj := &model.LogAdminOperation{}
	return Get(c, d.db, obj, id)
}

func (d *Dao) GetWhereLogAdminOperation(c *gin.Context, wvs []*gcurd.WhereValue) (*model.LogAdminOperation, error) {
	obj := &model.LogAdminOperation{}
	return GetWhere(c, d.db, obj, wvs)
}

func (d *Dao) FindLogAdminOperation(c *gin.Context, wvs []*gcurd.WhereValue) ([]*model.LogAdminOperation, error) {
	obj := &model.LogAdminOperation{}
	return Find(c, d.db, obj, wvs, obj.New)
}

func (d *Dao) PageTotalLogAdminOperation(c *gin.Context, req *gcurd.Request) (int, error) {
	obj := &model.LogAdminOperation{}
	return PageTotal(c, d.db, obj, req)
}

func (d *Dao) PageFindLogAdminOperation(c *gin.Context, req *gcurd.Request) ([]*model.LogAdminOperation, error) {
	obj := &model.LogAdminOperation{}
	return PageFind(c, d.db, obj, req, obj.New)
}

func (d *Dao) CreateMenu(c *gin.Context, obj *model.Menu) (sql.Result, error) {
	return Create(c, d.db, obj)
}

func (d *Dao) DeleteMenu(c *gin.Context, id int) error {
	obj := &model.Menu{}
	return Delete(c, d.db, obj, id)
}

func (d *Dao) UpdateMenu(c *gin.Context, id int, key string, value interface{}) error {
	obj := &model.Menu{}
	return Update(c, d.db, obj, id, key, value)
}

func (d *Dao) UpdateWhereMenu(c *gin.Context, key string, value interface{}, wvs []*gcurd.WhereValue) error {
	obj := &model.Menu{}
	return UpdateWhere(c, d.db, obj, key, value, wvs)
}

func (d *Dao) UpdateWhereKVMenu(c *gin.Context, kvs []gcurd.KeyValue, wvs []*gcurd.WhereValue) error {
	obj := &model.Menu{}
	return UpdateWhereKV(c, d.db, obj, kvs, wvs)
}

func (d *Dao) GetMenu(c *gin.Context, id int) (*model.Menu, error) {
	obj := &model.Menu{}
	return Get(c, d.db, obj, id)
}

func (d *Dao) GetWhereMenu(c *gin.Context, wvs []*gcurd.WhereValue) (*model.Menu, error) {
	obj := &model.Menu{}
	return GetWhere(c, d.db, obj, wvs)
}

func (d *Dao) FindMenu(c *gin.Context, wvs []*gcurd.WhereValue) ([]*model.Menu, error) {
	obj := &model.Menu{}
	return Find(c, d.db, obj, wvs, obj.New)
}

func (d *Dao) PageTotalMenu(c *gin.Context, req *gcurd.Request) (int, error) {
	obj := &model.Menu{}
	return PageTotal(c, d.db, obj, req)
}

func (d *Dao) PageFindMenu(c *gin.Context, req *gcurd.Request) ([]*model.Menu, error) {
	obj := &model.Menu{}
	return PageFind(c, d.db, obj, req, obj.New)
}

func (d *Dao) CreateRole(c *gin.Context, obj *model.Role) (sql.Result, error) {
	return Create(c, d.db, obj)
}

func (d *Dao) DeleteRole(c *gin.Context, id int) error {
	obj := &model.Role{}
	return Delete(c, d.db, obj, id)
}

func (d *Dao) UpdateRole(c *gin.Context, id int, key string, value interface{}) error {
	obj := &model.Role{}
	return Update(c, d.db, obj, id, key, value)
}

func (d *Dao) UpdateWhereRole(c *gin.Context, key string, value interface{}, wvs []*gcurd.WhereValue) error {
	obj := &model.Role{}
	return UpdateWhere(c, d.db, obj, key, value, wvs)
}

func (d *Dao) UpdateWhereKVRole(c *gin.Context, kvs []gcurd.KeyValue, wvs []*gcurd.WhereValue) error {
	obj := &model.Role{}
	return UpdateWhereKV(c, d.db, obj, kvs, wvs)
}

func (d *Dao) GetRole(c *gin.Context, id int) (*model.Role, error) {
	obj := &model.Role{}
	return Get(c, d.db, obj, id)
}

func (d *Dao) GetWhereRole(c *gin.Context, wvs []*gcurd.WhereValue) (*model.Role, error) {
	obj := &model.Role{}
	return GetWhere(c, d.db, obj, wvs)
}

func (d *Dao) FindRole(c *gin.Context, wvs []*gcurd.WhereValue) ([]*model.Role, error) {
	obj := &model.Role{}
	return Find(c, d.db, obj, wvs, obj.New)
}

func (d *Dao) PageTotalRole(c *gin.Context, req *gcurd.Request) (int, error) {
	obj := &model.Role{}
	return PageTotal(c, d.db, obj, req)
}

func (d *Dao) PageFindRole(c *gin.Context, req *gcurd.Request) ([]*model.Role, error) {
	obj := &model.Role{}
	return PageFind(c, d.db, obj, req, obj.New)
}

func (d *Dao) CreateRoleApi(c *gin.Context, obj *model.RoleApi) (sql.Result, error) {
	return Create(c, d.db, obj)
}

func (d *Dao) DeleteRoleApi(c *gin.Context, id int) error {
	obj := &model.RoleApi{}
	return Delete(c, d.db, obj, id)
}

func (d *Dao) UpdateRoleApi(c *gin.Context, id int, key string, value interface{}) error {
	obj := &model.RoleApi{}
	return Update(c, d.db, obj, id, key, value)
}

func (d *Dao) UpdateWhereRoleApi(c *gin.Context, key string, value interface{}, wvs []*gcurd.WhereValue) error {
	obj := &model.RoleApi{}
	return UpdateWhere(c, d.db, obj, key, value, wvs)
}

func (d *Dao) UpdateWhereKVRoleApi(c *gin.Context, kvs []gcurd.KeyValue, wvs []*gcurd.WhereValue) error {
	obj := &model.RoleApi{}
	return UpdateWhereKV(c, d.db, obj, kvs, wvs)
}

func (d *Dao) GetRoleApi(c *gin.Context, id int) (*model.RoleApi, error) {
	obj := &model.RoleApi{}
	return Get(c, d.db, obj, id)
}

func (d *Dao) GetWhereRoleApi(c *gin.Context, wvs []*gcurd.WhereValue) (*model.RoleApi, error) {
	obj := &model.RoleApi{}
	return GetWhere(c, d.db, obj, wvs)
}

func (d *Dao) FindRoleApi(c *gin.Context, wvs []*gcurd.WhereValue) ([]*model.RoleApi, error) {
	obj := &model.RoleApi{}
	return Find(c, d.db, obj, wvs, obj.New)
}

func (d *Dao) PageTotalRoleApi(c *gin.Context, req *gcurd.Request) (int, error) {
	obj := &model.RoleApi{}
	return PageTotal(c, d.db, obj, req)
}

func (d *Dao) PageFindRoleApi(c *gin.Context, req *gcurd.Request) ([]*model.RoleApi, error) {
	obj := &model.RoleApi{}
	return PageFind(c, d.db, obj, req, obj.New)
}

func (d *Dao) CreateRoleMenu(c *gin.Context, obj *model.RoleMenu) (sql.Result, error) {
	return Create(c, d.db, obj)
}

func (d *Dao) DeleteRoleMenu(c *gin.Context, id int) error {
	obj := &model.RoleMenu{}
	return Delete(c, d.db, obj, id)
}

func (d *Dao) UpdateRoleMenu(c *gin.Context, id int, key string, value interface{}) error {
	obj := &model.RoleMenu{}
	return Update(c, d.db, obj, id, key, value)
}

func (d *Dao) UpdateWhereRoleMenu(c *gin.Context, key string, value interface{}, wvs []*gcurd.WhereValue) error {
	obj := &model.RoleMenu{}
	return UpdateWhere(c, d.db, obj, key, value, wvs)
}

func (d *Dao) UpdateWhereKVRoleMenu(c *gin.Context, kvs []gcurd.KeyValue, wvs []*gcurd.WhereValue) error {
	obj := &model.RoleMenu{}
	return UpdateWhereKV(c, d.db, obj, kvs, wvs)
}

func (d *Dao) GetRoleMenu(c *gin.Context, id int) (*model.RoleMenu, error) {
	obj := &model.RoleMenu{}
	return Get(c, d.db, obj, id)
}

func (d *Dao) GetWhereRoleMenu(c *gin.Context, wvs []*gcurd.WhereValue) (*model.RoleMenu, error) {
	obj := &model.RoleMenu{}
	return GetWhere(c, d.db, obj, wvs)
}

func (d *Dao) FindRoleMenu(c *gin.Context, wvs []*gcurd.WhereValue) ([]*model.RoleMenu, error) {
	obj := &model.RoleMenu{}
	return Find(c, d.db, obj, wvs, obj.New)
}

func (d *Dao) PageTotalRoleMenu(c *gin.Context, req *gcurd.Request) (int, error) {
	obj := &model.RoleMenu{}
	return PageTotal(c, d.db, obj, req)
}

func (d *Dao) PageFindRoleMenu(c *gin.Context, req *gcurd.Request) ([]*model.RoleMenu, error) {
	obj := &model.RoleMenu{}
	return PageFind(c, d.db, obj, req, obj.New)
}
