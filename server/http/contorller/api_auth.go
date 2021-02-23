package contorller

import (
	"errors"
	"fmt"
	"github.com/bensema/goadmin/global"
	"github.com/bensema/goadmin/model"
	"github.com/bensema/goadmin/server/http/internal"
	"github.com/bensema/goadmin/utils"
	"github.com/gin-gonic/gin"
	"strings"
)

type ApiAuth struct{}

func (_this *ApiAuth) RegisterRoute(g *gin.RouterGroup) {
	g.GET("/api/v1/menu", _this.menu)
	g.GET("/api/v1/admin/pages", _this.adminPagesV1)
	g.GET("/api/v1/admin/info", _this.adminInfoV1)
	g.POST("/api/v1/admin/update", _this.updateAdmin)
	g.POST("/api/v1/admin/delete", _this.deleteAdmin)
	g.POST("/api/v1/admin/add", _this.addAdmin)
	g.GET("/api/v1/role/all", _this.roleAll)

	g.GET("/api/v1/role/pages", _this.rolePagesV1)
	g.GET("/api/v1/role/info", _this.roleInfoV1)
	g.POST("/api/v1/role/update", _this.updateRole)
	g.POST("/api/v1/role/add", _this.addRole)
	g.POST("/api/v1/role/delete", _this.deleteRole)

	g.GET("/api/v1/permission/all", _this.permissionAll)

}

func (_this *ApiAuth) menu(c *gin.Context) {
	session, _ := c.Cookie(internal.AdminSession)
	adminSession, _ := global.Srv.GetAdminSessionCache(c, session)
	menus, err := global.Srv.FindAdminMenu(c, adminSession.UserId)
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
	id, _ := c.GetQuery("id")
	reply, err := global.Srv.GetAdminV1(c, utils.GetInt(id))
	internal.JSON(c, reply, err)
}

func (_this *ApiAuth) roleAll(c *gin.Context) {
	reply, err := global.Srv.FindAllRole(c)
	internal.JSON(c, reply, err)
}

func (_this *ApiAuth) updateAdmin(c *gin.Context) {
	var filed []string
	arg := &model.UpdateAdmin{}
	id, _ := c.GetPostForm("id")
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
	arg.Id = utils.GetInt(id)
	arg.Status = utils.GetInt(status)
	arg.Roles = r

	internal.JSON(c, nil, global.Srv.UpdateAdmin(c, arg, filed))
}

func (_this *ApiAuth) deleteAdmin(c *gin.Context) {
	id, b := c.GetPostForm("id")
	if !b {
		internal.JSON(c, nil, errors.New("id不能空"))
		return
	}

	internal.JSON(c, nil, global.Srv.DeleteAdmin(c, utils.GetInt(id)))
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
		internal.JSON(c, nil, errors.New("选择适当的角色"))
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
