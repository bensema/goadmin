package service

import (
	"github.com/bensema/gcurd"
	"github.com/bensema/goadmin/model"
	"github.com/gin-gonic/gin"
)

func (s *Service) PageAdmin(c *gin.Context, req *gcurd.Request) (reply *model.PageReply[*model.Admin], err error) {
	obj := &model.Admin{}
	return gPage(c, s.dao.DB(), obj, req, obj.New)
}

func (s *Service) CreateAdmin(c *gin.Context, obj *model.Admin) error {
	return gCreate(c, s.dao.DB(), obj, model.OpAdminAdd, []string{})
}

func (s *Service) DeleteAdmin(c *gin.Context, id int) error {
	obj := &model.Admin{}
	return gDelete(c, s.dao.DB(), obj, id, model.OpAdminDel, []string{})
}

func (s *Service) GetAdmin(c *gin.Context, id int) (*model.Admin, error) {
	obj := &model.Admin{}
	return gGet(c, s.dao.DB(), obj, id)
}

func (s *Service) FindAdmin(c *gin.Context, wvs []*gcurd.WhereValue) ([]*model.Admin, error) {
	obj := &model.Admin{}
	return gFind(c, s.dao.DB(), obj, wvs, obj.New)
}

func (s *Service) UpdateAdmin(c *gin.Context, obj *model.Admin, id int, ignoreColumns []string, mosaicsColumns []string) error {
	return gUpdate(c, s.dao.DB(), obj, id, model.OpAdminUpdate, ignoreColumns, mosaicsColumns)
}

func (s *Service) PageAdminRole(c *gin.Context, req *gcurd.Request) (reply *model.PageReply[*model.AdminRole], err error) {
	obj := &model.AdminRole{}
	return gPage(c, s.dao.DB(), obj, req, obj.New)
}

func (s *Service) CreateAdminRole(c *gin.Context, obj *model.AdminRole) error {
	return gCreate(c, s.dao.DB(), obj, model.OpAdminRoleAdd, []string{})
}

func (s *Service) DeleteAdminRole(c *gin.Context, id int) error {
	obj := &model.AdminRole{}
	return gDelete(c, s.dao.DB(), obj, id, model.OpAdminRoleDel, []string{})
}

func (s *Service) GetAdminRole(c *gin.Context, id int) (*model.AdminRole, error) {
	obj := &model.AdminRole{}
	return gGet(c, s.dao.DB(), obj, id)
}

func (s *Service) FindAdminRole(c *gin.Context, wvs []*gcurd.WhereValue) ([]*model.AdminRole, error) {
	obj := &model.AdminRole{}
	return gFind(c, s.dao.DB(), obj, wvs, obj.New)
}

func (s *Service) UpdateAdminRole(c *gin.Context, obj *model.AdminRole, id int, ignoreColumns []string, mosaicsColumns []string) error {
	return gUpdate(c, s.dao.DB(), obj, id, model.OpAdminRoleUpdate, ignoreColumns, mosaicsColumns)
}

func (s *Service) PageApi(c *gin.Context, req *gcurd.Request) (reply *model.PageReply[*model.Api], err error) {
	obj := &model.Api{}
	return gPage(c, s.dao.DB(), obj, req, obj.New)
}

func (s *Service) CreateApi(c *gin.Context, obj *model.Api) error {
	return gCreate(c, s.dao.DB(), obj, model.OpApiAdd, []string{})
}

func (s *Service) DeleteApi(c *gin.Context, id int) error {
	obj := &model.Api{}
	return gDelete(c, s.dao.DB(), obj, id, model.OpApiDel, []string{})
}

func (s *Service) GetApi(c *gin.Context, id int) (*model.Api, error) {
	obj := &model.Api{}
	return gGet(c, s.dao.DB(), obj, id)
}

func (s *Service) FindApi(c *gin.Context, wvs []*gcurd.WhereValue) ([]*model.Api, error) {
	obj := &model.Api{}
	return gFind(c, s.dao.DB(), obj, wvs, obj.New)
}

func (s *Service) UpdateApi(c *gin.Context, obj *model.Api, id int, ignoreColumns []string, mosaicsColumns []string) error {
	return gUpdate(c, s.dao.DB(), obj, id, model.OpApiUpdate, ignoreColumns, mosaicsColumns)
}

func (s *Service) PageLogAdminLogin(c *gin.Context, req *gcurd.Request) (reply *model.PageReply[*model.LogAdminLogin], err error) {
	obj := &model.LogAdminLogin{}
	return gPage(c, s.dao.DB(), obj, req, obj.New)
}

func (s *Service) CreateLogAdminLogin(c *gin.Context, obj *model.LogAdminLogin) error {
	return gCreate(c, s.dao.DB(), obj, model.OpLogAdminLoginAdd, []string{})
}

func (s *Service) DeleteLogAdminLogin(c *gin.Context, id int) error {
	obj := &model.LogAdminLogin{}
	return gDelete(c, s.dao.DB(), obj, id, model.OpLogAdminLoginDel, []string{})
}

func (s *Service) GetLogAdminLogin(c *gin.Context, id int) (*model.LogAdminLogin, error) {
	obj := &model.LogAdminLogin{}
	return gGet(c, s.dao.DB(), obj, id)
}

func (s *Service) FindLogAdminLogin(c *gin.Context, wvs []*gcurd.WhereValue) ([]*model.LogAdminLogin, error) {
	obj := &model.LogAdminLogin{}
	return gFind(c, s.dao.DB(), obj, wvs, obj.New)
}

func (s *Service) UpdateLogAdminLogin(c *gin.Context, obj *model.LogAdminLogin, id int, ignoreColumns []string, mosaicsColumns []string) error {
	return gUpdate(c, s.dao.DB(), obj, id, model.OpLogAdminLoginUpdate, ignoreColumns, mosaicsColumns)
}

func (s *Service) PageLogAdminOperation(c *gin.Context, req *gcurd.Request) (reply *model.PageReply[*model.LogAdminOperation], err error) {
	obj := &model.LogAdminOperation{}
	return gPage(c, s.dao.DB(), obj, req, obj.New)
}

func (s *Service) CreateLogAdminOperation(c *gin.Context, obj *model.LogAdminOperation) error {
	return gCreate(c, s.dao.DB(), obj, model.OpLogAdminOperationAdd, []string{})
}

func (s *Service) DeleteLogAdminOperation(c *gin.Context, id int) error {
	obj := &model.LogAdminOperation{}
	return gDelete(c, s.dao.DB(), obj, id, model.OpLogAdminOperationDel, []string{})
}

func (s *Service) GetLogAdminOperation(c *gin.Context, id int) (*model.LogAdminOperation, error) {
	obj := &model.LogAdminOperation{}
	return gGet(c, s.dao.DB(), obj, id)
}

func (s *Service) FindLogAdminOperation(c *gin.Context, wvs []*gcurd.WhereValue) ([]*model.LogAdminOperation, error) {
	obj := &model.LogAdminOperation{}
	return gFind(c, s.dao.DB(), obj, wvs, obj.New)
}

func (s *Service) UpdateLogAdminOperation(c *gin.Context, obj *model.LogAdminOperation, id int, ignoreColumns []string, mosaicsColumns []string) error {
	return gUpdate(c, s.dao.DB(), obj, id, model.OpLogAdminOperationUpdate, ignoreColumns, mosaicsColumns)
}

func (s *Service) PageMenu(c *gin.Context, req *gcurd.Request) (reply *model.PageReply[*model.Menu], err error) {
	obj := &model.Menu{}
	return gPage(c, s.dao.DB(), obj, req, obj.New)
}

func (s *Service) CreateMenu(c *gin.Context, obj *model.Menu) error {
	return gCreate(c, s.dao.DB(), obj, model.OpMenuAdd, []string{})
}

func (s *Service) DeleteMenu(c *gin.Context, id int) error {
	obj := &model.Menu{}
	return gDelete(c, s.dao.DB(), obj, id, model.OpMenuDel, []string{})
}

func (s *Service) GetMenu(c *gin.Context, id int) (*model.Menu, error) {
	obj := &model.Menu{}
	return gGet(c, s.dao.DB(), obj, id)
}

func (s *Service) FindMenu(c *gin.Context, wvs []*gcurd.WhereValue) ([]*model.Menu, error) {
	obj := &model.Menu{}
	return gFind(c, s.dao.DB(), obj, wvs, obj.New)
}

func (s *Service) UpdateMenu(c *gin.Context, obj *model.Menu, id int, ignoreColumns []string, mosaicsColumns []string) error {
	return gUpdate(c, s.dao.DB(), obj, id, model.OpMenuUpdate, ignoreColumns, mosaicsColumns)
}

func (s *Service) PagePermission(c *gin.Context, req *gcurd.Request) (reply *model.PageReply[*model.Permission], err error) {
	obj := &model.Permission{}
	return gPage(c, s.dao.DB(), obj, req, obj.New)
}

func (s *Service) CreatePermission(c *gin.Context, obj *model.Permission) error {
	return gCreate(c, s.dao.DB(), obj, model.OpPermissionAdd, []string{})
}

func (s *Service) DeletePermission(c *gin.Context, id int) error {
	obj := &model.Permission{}
	return gDelete(c, s.dao.DB(), obj, id, model.OpPermissionDel, []string{})
}

func (s *Service) GetPermission(c *gin.Context, id int) (*model.Permission, error) {
	obj := &model.Permission{}
	return gGet(c, s.dao.DB(), obj, id)
}

func (s *Service) FindPermission(c *gin.Context, wvs []*gcurd.WhereValue) ([]*model.Permission, error) {
	obj := &model.Permission{}
	return gFind(c, s.dao.DB(), obj, wvs, obj.New)
}

func (s *Service) UpdatePermission(c *gin.Context, obj *model.Permission, id int, ignoreColumns []string, mosaicsColumns []string) error {
	return gUpdate(c, s.dao.DB(), obj, id, model.OpPermissionUpdate, ignoreColumns, mosaicsColumns)
}

func (s *Service) PagePermissionApi(c *gin.Context, req *gcurd.Request) (reply *model.PageReply[*model.PermissionApi], err error) {
	obj := &model.PermissionApi{}
	return gPage(c, s.dao.DB(), obj, req, obj.New)
}

func (s *Service) CreatePermissionApi(c *gin.Context, obj *model.PermissionApi) error {
	return gCreate(c, s.dao.DB(), obj, model.OpPermissionApiAdd, []string{})
}

func (s *Service) DeletePermissionApi(c *gin.Context, id int) error {
	obj := &model.PermissionApi{}
	return gDelete(c, s.dao.DB(), obj, id, model.OpPermissionApiDel, []string{})
}

func (s *Service) GetPermissionApi(c *gin.Context, id int) (*model.PermissionApi, error) {
	obj := &model.PermissionApi{}
	return gGet(c, s.dao.DB(), obj, id)
}

func (s *Service) FindPermissionApi(c *gin.Context, wvs []*gcurd.WhereValue) ([]*model.PermissionApi, error) {
	obj := &model.PermissionApi{}
	return gFind(c, s.dao.DB(), obj, wvs, obj.New)
}

func (s *Service) UpdatePermissionApi(c *gin.Context, obj *model.PermissionApi, id int, ignoreColumns []string, mosaicsColumns []string) error {
	return gUpdate(c, s.dao.DB(), obj, id, model.OpPermissionApiUpdate, ignoreColumns, mosaicsColumns)
}

func (s *Service) PagePermissionMenu(c *gin.Context, req *gcurd.Request) (reply *model.PageReply[*model.PermissionMenu], err error) {
	obj := &model.PermissionMenu{}
	return gPage(c, s.dao.DB(), obj, req, obj.New)
}

func (s *Service) CreatePermissionMenu(c *gin.Context, obj *model.PermissionMenu) error {
	return gCreate(c, s.dao.DB(), obj, model.OpPermissionMenuAdd, []string{})
}

func (s *Service) DeletePermissionMenu(c *gin.Context, id int) error {
	obj := &model.PermissionMenu{}
	return gDelete(c, s.dao.DB(), obj, id, model.OpPermissionMenuDel, []string{})
}

func (s *Service) GetPermissionMenu(c *gin.Context, id int) (*model.PermissionMenu, error) {
	obj := &model.PermissionMenu{}
	return gGet(c, s.dao.DB(), obj, id)
}

func (s *Service) FindPermissionMenu(c *gin.Context, wvs []*gcurd.WhereValue) ([]*model.PermissionMenu, error) {
	obj := &model.PermissionMenu{}
	return gFind(c, s.dao.DB(), obj, wvs, obj.New)
}

func (s *Service) UpdatePermissionMenu(c *gin.Context, obj *model.PermissionMenu, id int, ignoreColumns []string, mosaicsColumns []string) error {
	return gUpdate(c, s.dao.DB(), obj, id, model.OpPermissionMenuUpdate, ignoreColumns, mosaicsColumns)
}

func (s *Service) PageRole(c *gin.Context, req *gcurd.Request) (reply *model.PageReply[*model.Role], err error) {
	obj := &model.Role{}
	return gPage(c, s.dao.DB(), obj, req, obj.New)
}

func (s *Service) CreateRole(c *gin.Context, obj *model.Role) error {
	return gCreate(c, s.dao.DB(), obj, model.OpRoleAdd, []string{})
}

func (s *Service) DeleteRole(c *gin.Context, id int) error {
	obj := &model.Role{}
	return gDelete(c, s.dao.DB(), obj, id, model.OpRoleDel, []string{})
}

func (s *Service) GetRole(c *gin.Context, id int) (*model.Role, error) {
	obj := &model.Role{}
	return gGet(c, s.dao.DB(), obj, id)
}

func (s *Service) FindRole(c *gin.Context, wvs []*gcurd.WhereValue) ([]*model.Role, error) {
	obj := &model.Role{}
	return gFind(c, s.dao.DB(), obj, wvs, obj.New)
}

func (s *Service) UpdateRole(c *gin.Context, obj *model.Role, id int, ignoreColumns []string, mosaicsColumns []string) error {
	return gUpdate(c, s.dao.DB(), obj, id, model.OpRoleUpdate, ignoreColumns, mosaicsColumns)
}

func (s *Service) PageRoleMenu(c *gin.Context, req *gcurd.Request) (reply *model.PageReply[*model.RoleMenu], err error) {
	obj := &model.RoleMenu{}
	return gPage(c, s.dao.DB(), obj, req, obj.New)
}

func (s *Service) CreateRoleMenu(c *gin.Context, obj *model.RoleMenu) error {
	return gCreate(c, s.dao.DB(), obj, model.OpRoleMenuAdd, []string{})
}

func (s *Service) DeleteRoleMenu(c *gin.Context, id int) error {
	obj := &model.RoleMenu{}
	return gDelete(c, s.dao.DB(), obj, id, model.OpRoleMenuDel, []string{})
}

func (s *Service) GetRoleMenu(c *gin.Context, id int) (*model.RoleMenu, error) {
	obj := &model.RoleMenu{}
	return gGet(c, s.dao.DB(), obj, id)
}

func (s *Service) FindRoleMenu(c *gin.Context, wvs []*gcurd.WhereValue) ([]*model.RoleMenu, error) {
	obj := &model.RoleMenu{}
	return gFind(c, s.dao.DB(), obj, wvs, obj.New)
}

func (s *Service) UpdateRoleMenu(c *gin.Context, obj *model.RoleMenu, id int, ignoreColumns []string, mosaicsColumns []string) error {
	return gUpdate(c, s.dao.DB(), obj, id, model.OpRoleMenuUpdate, ignoreColumns, mosaicsColumns)
}

func (s *Service) PageRolePermission(c *gin.Context, req *gcurd.Request) (reply *model.PageReply[*model.RolePermission], err error) {
	obj := &model.RolePermission{}
	return gPage(c, s.dao.DB(), obj, req, obj.New)
}

func (s *Service) CreateRolePermission(c *gin.Context, obj *model.RolePermission) error {
	return gCreate(c, s.dao.DB(), obj, model.OpRolePermissionAdd, []string{})
}

func (s *Service) DeleteRolePermission(c *gin.Context, id int) error {
	obj := &model.RolePermission{}
	return gDelete(c, s.dao.DB(), obj, id, model.OpRolePermissionDel, []string{})
}

func (s *Service) GetRolePermission(c *gin.Context, id int) (*model.RolePermission, error) {
	obj := &model.RolePermission{}
	return gGet(c, s.dao.DB(), obj, id)
}

func (s *Service) FindRolePermission(c *gin.Context, wvs []*gcurd.WhereValue) ([]*model.RolePermission, error) {
	obj := &model.RolePermission{}
	return gFind(c, s.dao.DB(), obj, wvs, obj.New)
}

func (s *Service) UpdateRolePermission(c *gin.Context, obj *model.RolePermission, id int, ignoreColumns []string, mosaicsColumns []string) error {
	return gUpdate(c, s.dao.DB(), obj, id, model.OpRolePermissionUpdate, ignoreColumns, mosaicsColumns)
}
