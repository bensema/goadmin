package http

import (
	"errors"
	"fmt"
	"github.com/bensema/gcurd"
	"github.com/bensema/goadmin/model"
	"github.com/bensema/goadmin/utils"
	"github.com/gin-gonic/gin"
	"library/ecode"
	"library/xtime"
	"strings"
)

type ApiAuth struct{}

func (_this *ApiAuth) RegisterRoute(g *gin.RouterGroup) {
	g.GET("/api/menus", _this.menu)                // 获取菜单
	g.GET("/api/admin/pages", _this.pageAdminV2)   // 分页查询管理员
	g.GET("/api/admin/info", _this.adminInfoV1)    // 管理员信息
	g.POST("/api/admin/update", _this.updateAdmin) // 更新管理员 （管理员，角色）
	g.POST("/api/admin/add", _this.addAdminV2)     // 添加管理员
	g.POST("/api/admin/delete", _this.deleteAdmin) // 删除管理员

	g.GET("/api/role/pages", _this.rolePagesV1)    // 分页查询角色
	g.GET("/api/role/info", _this.roleInfoV1)      // 角色信息
	g.POST("/api/role/update", _this.updateRoleV2) // 更新角色（角色，权限）
	g.POST("/api/role/add", _this.addRole)         // 添加角色
	g.POST("/api/role/delete", _this.deleteRole)   // 删除角色

	g.GET("/api/permission/pages", _this.permissionPages)
	g.GET("/api/permission/info", _this.permissionInfoV1)
	g.POST("/api/permission/update", _this.updatePermissionV1) // 更新权限（menu，api）
	g.POST("/api/permission/add", _this.addPermissionV1)
	g.POST("/api/permission/delete", _this.deletePermissionV1)

	g.GET("/api/role_api/find", _this.findRolePermission) // 查询指定权限操作

	g.GET("/api/menu/pages", _this.menuPages)
	g.GET("/api/menu/info", _this.menuInfo)
	g.POST("/api/menu/add", _this.addMenu)       // 添加菜单
	g.POST("/api/menu/delete", _this.deleteMenu) // 删除菜单
	g.POST("/api/menu/update", _this.updateMenu) // 更新菜单

	g.GET("/api/api/pages", _this.apiPages)
	g.GET("/api/api/info", _this.apiInfo)
	g.POST("/api/api/add", _this.addApi)       // 添加操作功能
	g.POST("/api/api/delete", _this.deleteApi) // 删除操作功能
	g.POST("/api/api/update", _this.updateApi) // 更新操作功能

	g.GET("/api/log_login/pages", _this.logLogin)              // 分页查询登录信息
	g.GET("/api/log_operation/pages", _this.logAdminOperation) // 分页查询操作记录

	g.GET("/api/role/all", _this.roleAll)             // 获取所有角色
	g.GET("/api/permission/all", _this.permissionAll) // 获取所有权限
	g.GET("/api/menu/all", _this.menuAll)             // 获取所有菜单
	g.GET("/api/api/all", _this.apiAll)               // 获取所有操作

}

func (_this *ApiAuth) menu(c *gin.Context) {
	obj := &model.GinSession{}
	err := sessionSrv.GinLoadSession(c, obj)
	menus, err := srv.FindAdminMenu(c, obj.AdminId)
	JSON(c, menus, err)
}

func (_this *ApiAuth) pageAdmin(c *gin.Context) {
	req := &gcurd.Request{}
	req = prepareReq(c, req)

	if name, b := c.GetQuery("name"); b {
		req.Where = append(req.Where, gcurd.EQ("name", name))
	}

	reply, err := srv.PageAdmin(c, req)
	for _, d := range reply.Rows {
		d.Password = "***"
	}
	JSON(c, reply, err)
}

// 返回
func (_this *ApiAuth) pageAdminV2(c *gin.Context) {
	req := &gcurd.Request{}
	req = prepareReq(c, req)

	if name, b := c.GetQuery("name"); b {
		req.Where = append(req.Where, gcurd.EQ("name", name))
	}
	reply, err := srv.FindAdminPageV2(c, req)
	JSON(c, reply, err)
}

func (_this *ApiAuth) adminInfoV1(c *gin.Context) {
	adminId, _ := c.GetQuery("id")
	reply, err := srv.GetAdminV1(c, utils.GetInt(adminId))
	JSON(c, reply, err)
}

func (_this *ApiAuth) roleAll(c *gin.Context) {
	reply, err := srv.FindAllRole(c)
	JSON(c, reply, err)
}

func (_this *ApiAuth) updateAdmin(c *gin.Context) {
	var filed []string
	arg := &model.UpdateAdmin{}
	adminId, _ := c.GetPostForm("id")
	status, b := c.GetPostForm("status")
	if b {
		filed = append(filed, "status")
	}
	remark, b := c.GetPostForm("remark")
	if b {
		filed = append(filed, "remark")
	}
	roles, b := c.GetPostForm("roles")
	if b {
		filed = append(filed, "roles")
	}
	var r []int
	var err error
	if roles == "" {
		r = []int{}
	} else {
		r, err = utils.S2IList(strings.Split(roles, ","))
		if err != nil {
			fmt.Println(err)
			JSON(c, nil, err)
			return
		}
	}
	arg.AdminId = utils.GetInt(adminId)
	arg.Status = status
	arg.Remark = remark
	arg.Roles = r

	JSON(c, nil, srv.UpdateAdminV2(c, arg, filed))
}

func (_this *ApiAuth) deleteAdmin(c *gin.Context) {
	adminId, b := c.GetPostForm("id")
	if !b {
		JSON(c, nil, errors.New("id 不能空"))
		return
	}

	if adminId == RootId {
		JSON(c, nil, ecode.RejectOperation)
		return
	}

	JSON(c, nil, srv.DeleteAdminV2(c, utils.GetInt(adminId)))
}

func (_this *ApiAuth) addAdminV2(c *gin.Context) {
	arg := &model.AddAdmin{}
	name, _ := c.GetPostForm("username")
	password, _ := c.GetPostForm("password")
	status, _ := c.GetPostForm("status")
	roles, _ := c.GetPostForm("roles")

	r, err := utils.S2IList(strings.Split(roles, ","))
	if err != nil {
		JSON(c, nil, errors.New("选择适当的角色"))
		return
	}
	arg.Name = name
	arg.Password = password
	arg.Status = status
	arg.Roles = r
	JSON(c, nil, srv.AddAdminV2(c, arg))
}

// 角色

func (_this *ApiAuth) rolePagesV1(c *gin.Context) {
	req := &gcurd.Request{}
	req = prepareReq(c, req)
	if name, b := c.GetQuery("name"); b {
		req.Where = append(req.Where, gcurd.EQ("name", name))
	}
	reply, err := srv.FindRolePageV2(c, req)
	JSON(c, reply, err)
}
func (_this *ApiAuth) permissionAll(c *gin.Context) {
	reply, err := srv.FindAllPermission(c)
	JSON(c, reply, err)
}

func (_this *ApiAuth) menuAll(c *gin.Context) {
	reply, err := srv.FindAllMenu(c)
	JSON(c, reply, err)
}

func (_this *ApiAuth) apiAll(c *gin.Context) {
	reply, err := srv.FindAllApi(c)
	JSON(c, reply, err)
}

func (_this *ApiAuth) roleInfoV1(c *gin.Context) {
	id, _ := c.GetQuery("id")
	reply, err := srv.GetRoleInfoV1(c, utils.GetInt(id))
	JSON(c, reply, err)
}

func (_this *ApiAuth) updateRoleV2(c *gin.Context) {
	var filed []string
	arg := &model.UpdateRole{}
	roleId, _ := c.GetPostForm("id")
	name, b := c.GetPostForm("name")
	if b {
		filed = append(filed, "name")
	}
	remark, b := c.GetPostForm("remark")
	if b {
		filed = append(filed, "remark")
	}
	permissions, b := c.GetPostForm("permissions")
	if b {
		filed = append(filed, "permissions")
	}
	var r []int
	var err error
	if permissions == "" {
		r = []int{}
	} else {
		r, err = utils.S2IList(strings.Split(permissions, ","))
		if err != nil {
			fmt.Println(err)
			JSON(c, nil, err)
			return
		}
	}
	arg.RoleId = utils.GetInt(roleId)
	arg.Name = name
	arg.Remark = remark
	arg.Permissions = r

	JSON(c, nil, srv.UpdateRoleV2(c, arg, filed))
}

func (_this *ApiAuth) addRole(c *gin.Context) {
	obj := &model.AddRole{}
	name, _ := c.GetPostForm("name")
	remark, _ := c.GetPostForm("remark")
	permissions, _ := c.GetPostForm("permissions")

	r, err := utils.S2IList(strings.Split(permissions, ","))
	if err != nil {
		JSON(c, nil, errors.New("选择适当的权限"))
		return
	}

	obj.Name = name
	obj.Remark = remark
	obj.Permissions = r
	JSON(c, nil, srv.AddRoleV2(c, obj))
}

func (_this *ApiAuth) deleteRole(c *gin.Context) {
	id, b := c.GetPostForm("id")
	if !b {
		JSON(c, nil, errors.New("id不能空"))
		return
	}

	JSON(c, nil, srv.DeleteRoleV1(c, utils.GetInt(id)))
}

// permission

func (_this *ApiAuth) permissionPages(c *gin.Context) {
	req := &gcurd.Request{}
	req = prepareReq(c, req)
	if name, b := c.GetQuery("name"); b {
		req.Where = append(req.Where, gcurd.EQ("name", name))
	}
	reply, err := srv.PagePermission(c, req)
	JSON(c, reply, err)
}

func (_this *ApiAuth) permissionInfoV1(c *gin.Context) {
	id, _ := c.GetQuery("id")
	reply, err := srv.GetPermissionInfoV1(c, utils.GetInt(id))
	JSON(c, reply, err)
}

func (_this *ApiAuth) updatePermissionV1(c *gin.Context) {
	var filed []string
	arg := &model.UpdatePermission{}
	permissionId, _ := c.GetPostForm("id")
	name, b := c.GetPostForm("name")
	if b {
		filed = append(filed, "name")
	}
	permissionGroup, b := c.GetPostForm("permission_group")
	if b {
		filed = append(filed, "permission_group")
	}
	remark, b := c.GetPostForm("remark")
	if b {
		filed = append(filed, "remark")
	}

	menus, b := c.GetPostForm("menus")
	if b {
		filed = append(filed, "menus")
	}
	apis, b := c.GetPostForm("apis")
	if b {
		filed = append(filed, "apis")
	}

	var ms []int
	var err error
	if menus == "" {
		ms = []int{}
	} else {
		ms, err = utils.S2IList(strings.Split(menus, ","))
		if err != nil {
			JSON(c, nil, err)
			return
		}
	}
	var as []int
	if apis == "" {
		as = []int{}
	} else {
		as, err = utils.S2IList(strings.Split(apis, ","))
		if err != nil {
			JSON(c, nil, err)
			return
		}
	}

	arg.Id = utils.GetInt(permissionId)
	arg.Name = name
	arg.PermissionGroup = permissionGroup
	arg.Remark = remark
	arg.Menus = ms
	arg.Apis = as

	JSON(c, nil, srv.UpdatePermissionV1(c, arg, filed))
}

func (_this *ApiAuth) addPermissionV1(c *gin.Context) {
	obj := &model.AddPermission{}
	name, _ := c.GetPostForm("name")
	permissionGroup, _ := c.GetPostForm("permission_group")
	remark, _ := c.GetPostForm("remark")
	menus, _ := c.GetPostForm("menus")
	apis, _ := c.GetPostForm("apis")

	ms, err := utils.S2IList(strings.Split(menus, ","))
	if err != nil {
		JSON(c, nil, errors.New("select menu"))
		return
	}

	as, err := utils.S2IList(strings.Split(apis, ","))
	if err != nil {
		JSON(c, nil, errors.New("select api"))
		return
	}

	obj.Name = name
	obj.PermissionGroup = permissionGroup
	obj.Remark = remark
	obj.Menus = ms
	obj.Apis = as
	JSON(c, nil, srv.AddPermissionV1(c, obj))
}

func (_this *ApiAuth) deletePermissionV1(c *gin.Context) {
	permissionId, b := c.GetPostForm("id")
	if !b {
		JSON(c, nil, errors.New("id is null"))
		return
	}

	JSON(c, nil, srv.DeletePermissionV1(c, utils.GetInt(permissionId)))
}

func (_this *ApiAuth) menuInfo(c *gin.Context) {
	id, _ := c.GetQuery("id")
	reply, err := srv.GetMenu(c, utils.GetInt(id))
	JSON(c, reply, err)
}

func (_this *ApiAuth) findRolePermission(c *gin.Context) {
	var wvs []*gcurd.WhereValue
	id, _ := c.GetQuery("id")
	wvs = append(wvs, gcurd.EQ("role_id", utils.GetInt(id)))
	// todo
	//reply, err := srv.FindRolePermission(c, wvs)
	//JSON(c, reply, err)
}

// menus
func (_this *ApiAuth) menuPages(c *gin.Context) {
	req := &gcurd.Request{}
	req = prepareReq(c, req)
	if name, b := c.GetQuery("name"); b {
		req.Where = append(req.Where, gcurd.EQ("name", name))
	}
	reply, err := srv.PageMenu(c, req)
	JSON(c, reply, err)
}

func (_this *ApiAuth) addMenu(c *gin.Context) {
	obj := &model.Menu{}
	name, _ := c.GetPostForm("name")
	pid, _ := c.GetPostForm("pid")
	icon, _ := c.GetPostForm("icon")
	url, _ := c.GetPostForm("url")
	indexSort, _ := c.GetPostForm("index_sort")

	obj.Name = name
	obj.Pid = utils.GetInt(pid)
	obj.Icon = icon
	obj.Url = url
	obj.IndexSort = utils.GetInt(indexSort)

	JSON(c, nil, srv.CreateMenu(c, obj))
}

func (_this *ApiAuth) updateMenu(c *gin.Context) {
	obj := &model.Menu{}
	var (
		Id  int
		err error
	)
	if id, b := c.GetPostForm("id"); !b {
		JSON(c, nil, errors.New("id不能空"))
		return
	} else {
		Id = utils.GetInt(id)
	}
	obj, err = srv.GetMenu(c, Id)
	if err != nil {
		JSON(c, nil, errors.New("menu 不存在"))
		return
	}

	if name, b := c.GetPostForm("name"); b {
		obj.Name = name
	}
	if pid, b := c.GetPostForm("pid"); b {
		obj.Pid = utils.GetInt(pid)
	}
	if icon, b := c.GetPostForm("icon"); b {
		obj.Icon = icon
	}
	if url, b := c.GetPostForm("url"); b {
		obj.Url = url
	}
	if indexSort, b := c.GetPostForm("index_sort"); b {
		obj.IndexSort = utils.GetInt(indexSort)
	}
	fmt.Println(obj, obj.Url)
	JSON(c, nil, srv.UpdateMenu(c, obj, Id, []string{}, []string{}))

}

func (_this *ApiAuth) deleteMenu(c *gin.Context) {
	id, b := c.GetPostForm("id")
	if !b {
		JSON(c, nil, errors.New("id不能空"))
		return
	}

	JSON(c, nil, srv.DeleteMenu(c, utils.GetInt(id)))
}

// api

func (_this *ApiAuth) apiPages(c *gin.Context) {
	req := &gcurd.Request{}
	req = prepareReq(c, req)
	if name, b := c.GetQuery("name"); b {
		req.Where = append(req.Where, gcurd.EQ("name", name))
	}
	reply, err := srv.PageApi(c, req)
	JSON(c, reply, err)
}

func (_this *ApiAuth) apiInfo(c *gin.Context) {
	id, _ := c.GetQuery("id")
	reply, err := srv.GetApi(c, utils.GetInt(id))
	JSON(c, reply, err)
}

func (_this *ApiAuth) addApi(c *gin.Context) {
	obj := &model.Api{}
	name, _ := c.GetPostForm("name")
	method, _ := c.GetPostForm("method")
	url, _ := c.GetPostForm("url")

	obj.Name = name
	obj.Url = url
	obj.Method = method

	JSON(c, nil, srv.CreateApi(c, obj))
}

func (_this *ApiAuth) deleteApi(c *gin.Context) {
	id, b := c.GetPostForm("id")
	if !b {
		JSON(c, nil, errors.New("id不能空"))
		return
	}

	JSON(c, nil, srv.DeleteApi(c, utils.GetInt(id)))
}

func (_this *ApiAuth) updateApi(c *gin.Context) {
	obj := &model.Api{}
	var (
		Id  int
		err error
	)
	if id, b := c.GetPostForm("id"); !b {
		JSON(c, nil, errors.New("id不能空"))
		return
	} else {
		Id = utils.GetInt(id)
	}
	obj, err = srv.GetApi(c, Id)
	if err != nil {
		JSON(c, nil, errors.New("api 不存在"))
		return
	}

	if name, b := c.GetPostForm("name"); b {
		obj.Name = name
	}
	if url, b := c.GetPostForm("url"); b {
		obj.Url = url
	}
	if method, b := c.GetPostForm("method"); b {
		obj.Method = method
	}

	JSON(c, nil, srv.UpdateApi(c, obj, Id, []string{}, []string{}))
}

func (_this *ApiAuth) logLogin(c *gin.Context) {
	req := &gcurd.Request{}
	req = prepareReq(c, req)

	if name, b := c.GetQuery("name"); b && name != "" {
		req.Where = append(req.Where, gcurd.EQ("name", name))
	}
	if ip, b := c.GetQuery("ip"); b && ip != "" {
		req.Where = append(req.Where, gcurd.EQ("ip", ip))
	}
	if result, b := c.GetQuery("result"); b && result != "" {
		req.Where = append(req.Where, gcurd.EQ("result", result))
	}
	if recordAtFrom, b := c.GetQuery("record_at_from"); b && recordAtFrom != "" {
		req.Where = append(req.Where, gcurd.GTE("record_at", xtime.Time(utils.GetInt64(recordAtFrom))))
	}
	if recordAtTo, b := c.GetQuery("record_at_to"); b && recordAtTo != "" {
		req.Where = append(req.Where, gcurd.LT("record_at", xtime.Time(utils.GetInt64(recordAtTo))))
	}

	reply, err := srv.PageLogAdminLogin(c, req)
	JSON(c, reply, err)
}

func (_this *ApiAuth) logAdminOperation(c *gin.Context) {
	req := &gcurd.Request{}
	req = prepareReq(c, req)

	if name, b := c.GetQuery("name"); b && name != "" {
		req.Where = append(req.Where, gcurd.EQ("name", name))
	}
	if ip, b := c.GetQuery("ip"); b && ip != "" {
		req.Where = append(req.Where, gcurd.EQ("ip", ip))
	}
	if result, b := c.GetQuery("result"); b && result != "" {
		req.Where = append(req.Where, gcurd.EQ("result", result))
	}
	if action, b := c.GetQuery("action"); b && action != "" {
		req.Where = append(req.Where, gcurd.EQ("action", action))
	}
	if module, b := c.GetQuery("module"); b && module != "" {
		req.Where = append(req.Where, gcurd.EQ("module", module))
	}
	if recordAtFrom, b := c.GetQuery("record_at_from"); b && recordAtFrom != "" {
		req.Where = append(req.Where, gcurd.GTE("record_at", xtime.Time(utils.GetInt64(recordAtFrom))))
	}
	if recordAtTo, b := c.GetQuery("record_at_to"); b && recordAtTo != "" {
		req.Where = append(req.Where, gcurd.LT("record_at", xtime.Time(utils.GetInt64(recordAtTo))))
	}

	reply, err := srv.PageLogAdminOperation(c, req)
	JSON(c, reply, err)
}
