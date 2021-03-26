package contorller

import (
	"errors"
	"fmt"
	"github.com/bensema/goadmin/global"
	"github.com/bensema/goadmin/model"
	"github.com/bensema/goadmin/server/http/internal"
	"github.com/bensema/goadmin/utils"
	"github.com/bensema/library/ecode"
	xtime "github.com/bensema/library/time"
	"github.com/gin-gonic/gin"
	"strings"
)

type ApiAuth struct{}

func (_this *ApiAuth) RegisterRoute(g *gin.RouterGroup) {
	g.GET("/api/v1/menu", _this.menu)                 // 获取菜单
	g.GET("/api/v1/admin/pages", _this.adminPagesV1)  // 分页查询管理员
	g.GET("/api/v1/admin/info", _this.adminInfoV1)    // 管理员信息
	g.POST("/api/v1/admin/update", _this.updateAdmin) // 更新管理员 （管理员，角色）
	g.POST("/api/v1/admin/delete", _this.deleteAdmin) // 删除管理员
	g.POST("/api/v1/admin/add", _this.addAdmin)       // 添加管理员

	g.GET("/api/v1/role/pages", _this.rolePagesV1)  // 分页查询角色
	g.GET("/api/v1/role/info", _this.roleInfoV1)    // 角色信息
	g.POST("/api/v1/role/update", _this.updateRole) // 更新角色（角色，权限）
	g.POST("/api/v1/role/add", _this.addRole)       // 添加角色
	g.POST("/api/v1/role/delete", _this.deleteRole) // 删除角色

	g.GET("/api/v1/permission/pages", _this.permissionPages)                  // 分页查询权限
	g.POST("/api/v1/permission/add", _this.addPermission)                     // 添加权限
	g.POST("/api/v1/permission/update", _this.updatePermission)               // 更新权限（权限，菜单，操作）
	g.POST("/api/v1/permission/delete", _this.deletePermission)               // 删除权限
	g.GET("/api/v1/permission_menu/find", _this.findPermissionMenu)           // 查询指定权限菜单
	g.GET("/api/v1/permission_operation/find", _this.findPermissionOperation) // 查询指定权限操作

	g.POST("/api/v1/menu/add", _this.addMenu)       // 添加菜单
	g.POST("/api/v1/menu/delete", _this.deleteMenu) // 删除菜单
	g.POST("/api/v1/menu/update", _this.updateMenu) // 更新菜单

	g.POST("/api/v1/operation/add", _this.addOperation)       // 添加操作功能
	g.POST("/api/v1/operation/delete", _this.deleteOperation) // 删除操作功能
	g.POST("/api/v1/operation/update", _this.updateOperation) // 更新操作功能

	g.GET("/api/v1/log_login/pages", _this.logLogin)         // 分页查询登录信息
	g.GET("/api/v1/log_operation/pages", _this.logOperation) // 分页查询操作记录

	g.GET("/api/v1/role/all", _this.roleAll)             // 获取所有角色
	g.GET("/api/v1/permission/all", _this.permissionAll) // 获取所有权限
	g.GET("/api/v1/menu/all", _this.menuAll)             // 获取所有菜单
	g.GET("/api/v1/operation/all", _this.operationAll)   // 获取所有操作

	_this.RegisterBBAdminRoute(g)
}

func (_this *ApiAuth) menu(c *gin.Context) {
	session, _ := c.Cookie(internal.AdminSession)
	adminSession, _ := global.Srv.GetAdminSessionCache(c, session)
	menus, err := global.Srv.FindAdminMenu(c, adminSession.AdminId)
	internal.JSON(c, menus, err)
}

func (_this *ApiAuth) adminPages(c *gin.Context) {
	arg := &model.FindAdminReq{}
	name, _ := c.GetQuery("name")
	orderBy, _ := c.GetQuery("order_by")
	sort, _ := c.GetQuery("sort")
	num, _ := c.GetQuery("num")
	size, _ := c.GetQuery("size")

	arg.Name = name
	arg.OrderBy = orderBy
	arg.Sort = sort
	arg.Num = utils.GetInt(num)
	arg.Size = utils.GetInt(size)

	arg.Verify()
	reply, err := global.Srv.FindAdminPage(c, arg)
	internal.JSON(c, reply, err)
}

// 返回
func (_this *ApiAuth) adminPagesV1(c *gin.Context) {
	arg := &model.FindAdminReq{}
	name, _ := c.GetQuery("name")
	orderBy, _ := c.GetQuery("order_by")
	sort, _ := c.GetQuery("sort")
	num, _ := c.GetQuery("num")
	size, _ := c.GetQuery("size")

	arg.Name = name
	arg.OrderBy = orderBy
	arg.Sort = sort
	arg.Num = utils.GetInt(num)
	arg.Size = utils.GetInt(size)

	arg.Verify()
	reply, err := global.Srv.FindAdminPageV1(c, arg)
	internal.JSON(c, reply, err)
}

func (_this *ApiAuth) adminInfoV1(c *gin.Context) {
	adminId, _ := c.GetQuery("admin_id")
	reply, err := global.Srv.GetAdminV1(c, adminId)
	internal.JSON(c, reply, err)
}

func (_this *ApiAuth) roleAll(c *gin.Context) {
	reply, err := global.Srv.FindAllRole(c)
	internal.JSON(c, reply, err)
}

func (_this *ApiAuth) updateAdmin(c *gin.Context) {
	var filed []string
	arg := &model.UpdateAdmin{}
	adminId, _ := c.GetPostForm("admin_id")
	status, b := c.GetPostForm("status")
	if b {
		filed = append(filed, "status")
	}
	roles, b := c.GetPostForm("roles")
	if b {
		filed = append(filed, "roles")
	}
	r, err := utils.S2IList(strings.Split(roles, ","))
	if err != nil {
		internal.JSON(c, nil, err)
		return
	}
	arg.AdminId = adminId
	arg.Status = utils.GetInt(status)
	arg.Roles = r

	internal.JSON(c, nil, global.Srv.UpdateAdmin(c, arg, filed))
}

func (_this *ApiAuth) deleteAdmin(c *gin.Context) {
	adminId, b := c.GetPostForm("admin_id")
	if !b {
		internal.JSON(c, nil, errors.New("admin_id 不能空"))
		return
	}

	if adminId == internal.RootId {
		internal.JSON(c, nil, ecode.RejectOperation)
		return
	}

	internal.JSON(c, nil, global.Srv.DeleteAdmin(c, adminId))
}

func (_this *ApiAuth) addAdmin(c *gin.Context) {
	arg := &model.AddAdmin{}
	name, _ := c.GetPostForm("name")
	password, _ := c.GetPostForm("password")
	status, _ := c.GetPostForm("status")
	roles, _ := c.GetPostForm("roles")

	r, err := utils.S2IList(strings.Split(roles, ","))
	if err != nil {
		internal.JSON(c, nil, errors.New("选择适当的角色"))
		return
	}
	arg.Name = name
	arg.Password = password
	arg.Status = utils.GetInt(status)
	arg.Roles = r
	internal.JSON(c, nil, global.Srv.AddAdmin(c, arg))
}

// 角色

func (_this *ApiAuth) rolePagesV1(c *gin.Context) {
	arg := &model.FindRoleReq{}
	//name, _ := c.GetQuery("name")
	orderBy, _ := c.GetQuery("order_by")
	sort, _ := c.GetQuery("sort")
	num, _ := c.GetQuery("num")
	size, _ := c.GetQuery("size")

	//arg.Name = name
	arg.OrderBy = orderBy
	arg.Sort = sort
	arg.Num = utils.GetInt(num)
	arg.Size = utils.GetInt(size)

	arg.Verify()
	reply, err := global.Srv.FindRolePageV1(c, arg)
	internal.JSON(c, reply, err)
}

func (_this *ApiAuth) roleInfoV1(c *gin.Context) {
	id, _ := c.GetQuery("id")
	reply, err := global.Srv.GetRoleV1(c, utils.GetInt(id))
	internal.JSON(c, reply, err)
}

func (_this *ApiAuth) permissionAll(c *gin.Context) {
	reply, err := global.Srv.FindAllPermission(c)
	internal.JSON(c, reply, err)
}

func (_this *ApiAuth) menuAll(c *gin.Context) {
	reply, err := global.Srv.FindAllMenu(c)
	internal.JSON(c, reply, err)
}

func (_this *ApiAuth) operationAll(c *gin.Context) {
	reply, err := global.Srv.FindAllOperation(c)
	internal.JSON(c, reply, err)
}

func (_this *ApiAuth) updateRole(c *gin.Context) {
	var filed []string
	arg := &model.UpdateRole{}
	id, _ := c.GetPostForm("id")
	name, b := c.GetPostForm("name")
	if b {
		filed = append(filed, "name")
	}
	permissions, b := c.GetPostForm("permissions")
	if b {
		filed = append(filed, "permissions")
	}
	r, err := utils.S2IList(strings.Split(permissions, ","))
	if err != nil {
		internal.JSON(c, nil, err)
		return
	}
	arg.Id = utils.GetInt(id)
	arg.Name = name
	arg.Permissions = r
	fmt.Println(id, name, r)
	internal.JSON(c, nil, global.Srv.UpdateRole(c, arg, filed))
}

func (_this *ApiAuth) addRole(c *gin.Context) {
	arg := &model.AddRole{}
	name, _ := c.GetPostForm("name")
	permissions, _ := c.GetPostForm("permissions")

	r, err := utils.S2IList(strings.Split(permissions, ","))
	if err != nil {
		internal.JSON(c, nil, errors.New("选择适当的权限"))
		return
	}
	arg.Name = name
	arg.Permissions = r
	internal.JSON(c, nil, global.Srv.AddRole(c, arg))
}

func (_this *ApiAuth) deleteRole(c *gin.Context) {
	id, b := c.GetPostForm("id")
	if !b {
		internal.JSON(c, nil, errors.New("id不能空"))
		return
	}

	internal.JSON(c, nil, global.Srv.DeleteRole(c, utils.GetInt(id)))
}

// 权限

func (_this *ApiAuth) permissionPages(c *gin.Context) {
	arg := &model.FindPermissionReq{}
	//name, _ := c.GetQuery("name")
	orderBy, _ := c.GetQuery("order_by")
	sort, _ := c.GetQuery("sort")
	num, _ := c.GetQuery("num")
	size, _ := c.GetQuery("size")

	//arg.Name = name
	arg.OrderBy = orderBy
	arg.Sort = sort
	arg.Num = utils.GetInt(num)
	arg.Size = utils.GetInt(size)

	arg.Verify()
	reply, err := global.Srv.FindPermissionPage(c, arg)
	internal.JSON(c, reply, err)
}

func (_this *ApiAuth) addPermission(c *gin.Context) {
	arg := &model.Permission{}
	name, _ := c.GetPostForm("name")
	if name == "" {
		internal.JSON(c, nil, errors.New("name不能为空"))
	}

	arg.Name = name
	internal.JSON(c, nil, global.Srv.AddPermission(c, arg))
}

func (_this *ApiAuth) updatePermission(c *gin.Context) {
	var filed []string
	arg := &model.UpdatePermission{}
	id, _ := c.GetPostForm("id")
	name, _ := c.GetPostForm("name")
	var ms []int
	var ops []int
	var err error

	menus, b := c.GetPostForm("menus")
	if b && menus != "" {
		filed = append(filed, "menus")
		ms, err = utils.S2IList(strings.Split(menus, ","))
		if err != nil {
			internal.JSON(c, nil, errors.New("菜单数据错误"))
			return
		}
	}
	operations, b := c.GetPostForm("operations")
	if b && operations != "" {
		filed = append(filed, "operations")
		ops, err = utils.S2IList(strings.Split(operations, ","))
		if err != nil {
			internal.JSON(c, nil, errors.New("操作数据错误"))
			return
		}
	}

	arg.Id = utils.GetInt(id)
	arg.Name = name
	arg.Menus = ms
	arg.Operation = ops
	internal.JSON(c, nil, global.Srv.UpdatePermission(c, arg, filed))
}

func (_this *ApiAuth) deletePermission(c *gin.Context) {
	id, b := c.GetPostForm("id")
	if !b {
		internal.JSON(c, nil, errors.New("id不能空"))
		return
	}

	internal.JSON(c, nil, global.Srv.DeletePermission(c, utils.GetInt(id)))
}

func (_this *ApiAuth) findPermissionMenu(c *gin.Context) {
	id, _ := c.GetQuery("id")
	req := &model.FindPermissionMenuReq{}
	req.PermissionId = utils.GetInt(id)

	reply, err := global.Srv.FindPermissionMenu(c, req)
	internal.JSON(c, reply, err)
}

func (_this *ApiAuth) findPermissionOperation(c *gin.Context) {
	id, _ := c.GetQuery("id")
	req := &model.FindPermissionOperationReq{}
	req.PermissionId = utils.GetInt(id)

	reply, err := global.Srv.FindPermissionOperation(c, req)
	internal.JSON(c, reply, err)
}

func (_this *ApiAuth) addMenu(c *gin.Context) {
	arg := &model.Menu{}
	name, _ := c.GetPostForm("name")
	pid, _ := c.GetPostForm("pid")
	icon, _ := c.GetPostForm("icon")
	url, _ := c.GetPostForm("url")
	indexSort, _ := c.GetPostForm("index_sort")

	arg.Name = name
	arg.Pid = utils.GetInt(pid)
	arg.Icon = icon
	arg.Url = url
	arg.IndexSort = utils.GetInt(indexSort)

	internal.JSON(c, nil, global.Srv.AddMenu(c, arg))
}

func (_this *ApiAuth) updateMenu(c *gin.Context) {
	arg := &model.Menu{}
	var filed []string
	id, b := c.GetPostForm("id")
	if !b {
		internal.JSON(c, nil, errors.New("id不能空"))
		return
	}
	name, b := c.GetPostForm("name")
	if b {
		filed = append(filed, "name")
	}
	pid, b := c.GetPostForm("pid")
	if b {
		filed = append(filed, "pid")
	}
	icon, b := c.GetPostForm("icon")
	if b {
		filed = append(filed, "icon")
	}
	url, b := c.GetPostForm("url")
	if b {
		filed = append(filed, "url")
	}
	indexSort, b := c.GetPostForm("index_sort")
	if b {
		filed = append(filed, "index_sort")
	}

	arg.Id = utils.GetInt(id)
	arg.Name = name
	arg.Pid = utils.GetInt(pid)
	arg.Icon = icon
	arg.Url = url
	arg.IndexSort = utils.GetInt(indexSort)

	internal.JSON(c, nil, global.Srv.UpdateMenu(c, arg, filed))
}

func (_this *ApiAuth) deleteMenu(c *gin.Context) {
	id, b := c.GetPostForm("id")
	if !b {
		internal.JSON(c, nil, errors.New("id不能空"))
		return
	}

	internal.JSON(c, nil, global.Srv.DeleteMenu(c, utils.GetInt(id)))
}

func (_this *ApiAuth) addOperation(c *gin.Context) {
	arg := &model.Operation{}
	name, _ := c.GetPostForm("name")
	pid, _ := c.GetPostForm("pid")
	code, _ := c.GetPostForm("code")
	method, _ := c.GetPostForm("method")
	url, _ := c.GetPostForm("url")

	arg.Name = name
	arg.Pid = utils.GetInt(pid)
	arg.Code = code
	arg.Url = url
	arg.Method = method

	internal.JSON(c, nil, global.Srv.AddOperation(c, arg))
}

func (_this *ApiAuth) deleteOperation(c *gin.Context) {
	id, b := c.GetPostForm("id")
	if !b {
		internal.JSON(c, nil, errors.New("id不能空"))
		return
	}

	internal.JSON(c, nil, global.Srv.DeleteOperation(c, utils.GetInt(id)))
}

func (_this *ApiAuth) updateOperation(c *gin.Context) {
	arg := &model.Operation{}
	var filed []string
	id, b := c.GetPostForm("id")
	if !b {
		internal.JSON(c, nil, errors.New("id不能空"))
		return
	}
	name, b := c.GetPostForm("name")
	if b {
		filed = append(filed, "name")
	}
	pid, b := c.GetPostForm("pid")
	if b {
		filed = append(filed, "pid")
	}
	code, b := c.GetPostForm("code")
	if b {
		filed = append(filed, "code")
	}
	url, b := c.GetPostForm("url")
	if b {
		filed = append(filed, "url")
	}
	method, b := c.GetPostForm("method")
	if b {
		filed = append(filed, "method")
	}

	arg.Id = utils.GetInt(id)
	arg.Name = name
	arg.Pid = utils.GetInt(pid)
	arg.Code = code
	arg.Url = url
	arg.Method = method

	internal.JSON(c, nil, global.Srv.UpdateOperation(c, arg, filed))
}

func (_this *ApiAuth) logLogin(c *gin.Context) {
	arg := &model.FindLogAdminLoginReq{}
	name, _ := c.GetQuery("name")
	ip, _ := c.GetQuery("ip")
	result, _ := c.GetQuery("result")
	recordAtFrom, _ := c.GetQuery("record_at_from")
	recordAtTo, _ := c.GetQuery("record_at_to")
	orderBy, _ := c.GetQuery("order_by")
	sort, _ := c.GetQuery("sort")
	num, _ := c.GetQuery("num")
	size, _ := c.GetQuery("size")

	arg.Name = name
	arg.Ip = ip
	arg.Result = utils.GetInt(result)
	arg.RecordAtFrom = xtime.Time(utils.GetInt64(recordAtFrom))
	arg.RecordAtTo = xtime.Time(utils.GetInt64(recordAtTo))
	arg.OrderBy = orderBy
	arg.Sort = sort
	arg.Num = utils.GetInt(num)
	arg.Size = utils.GetInt(size)

	arg.Verify()
	reply, err := global.Srv.FindAdminLoginPage(c, arg)
	internal.JSON(c, reply, err)
}

func (_this *ApiAuth) logOperation(c *gin.Context) {
	arg := &model.FindLogAdminOperationReq{}
	name, _ := c.GetQuery("name")
	ip, _ := c.GetQuery("ip")
	result, _ := c.GetQuery("result")
	recordAtFrom, _ := c.GetQuery("record_at_from")
	recordAtTo, _ := c.GetQuery("record_at_to")
	orderBy, _ := c.GetQuery("order_by")
	sort, _ := c.GetQuery("sort")
	num, _ := c.GetQuery("num")
	size, _ := c.GetQuery("size")

	arg.Name = name
	arg.Ip = ip
	arg.Result = utils.GetInt(result)
	arg.RecordAtFrom = xtime.Time(utils.GetInt64(recordAtFrom))
	arg.RecordAtTo = xtime.Time(utils.GetInt64(recordAtTo))
	arg.OrderBy = orderBy
	arg.Sort = sort
	arg.Num = utils.GetInt(num)
	arg.Size = utils.GetInt(size)

	arg.Verify()
	reply, err := global.Srv.PageFindLogAdminOperation(c, arg)
	internal.JSON(c, reply, err)
}
