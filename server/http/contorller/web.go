package contorller

import (
	"github.com/bensema/goadmin/global"
	"github.com/bensema/goadmin/server/http/internal"
	"github.com/gin-gonic/gin"
	"net/http"
)

type HtmlWeb struct{}

func (_this *HtmlWeb) RegisterRoute(g *gin.RouterGroup) {
	g.GET("/login", _this.login)
	g.GET("/403", _this.e403)
	g.GET("/logout", _this.logout)
	g.GET("/admin/form", _this.adminForm)
	g.GET("/admin/add", _this.adminAdd)
	g.GET("/role/add", _this.roleAdd)
	g.GET("/role/form", _this.roleForm)
	g.GET("/permission/add", _this.permissionAdd)

}

func (_this *HtmlWeb) login(c *gin.Context) {
	_, err := c.Cookie(internal.AdminSession)
	if err != nil {
		c.HTML(http.StatusOK, "base/login.html", gin.H{})
	} else {
		c.Redirect(http.StatusFound, "/")
	}
}

func (_this *HtmlWeb) logout(c *gin.Context) {
	adminSession, err := c.Cookie(internal.AdminSession)
	if err != nil {
	} else {
		c.SetCookie(internal.AdminSession, "", -1, "/", "", false, true)
		global.Srv.DeleteAdminSessionCache(c, adminSession)
	}
	c.Redirect(http.StatusFound, "/login")
}

func (_this *HtmlWeb) adminForm(c *gin.Context) {
	c.HTML(http.StatusOK, "admin/adminform.html", gin.H{})
}

func (_this *HtmlWeb) adminAdd(c *gin.Context) {
	c.HTML(http.StatusOK, "admin/adminadd.html", gin.H{})
}

func (_this *HtmlWeb) roleForm(c *gin.Context) {
	c.HTML(http.StatusOK, "role/roleform.html", gin.H{})
}

func (_this *HtmlWeb) roleAdd(c *gin.Context) {
	c.HTML(http.StatusOK, "role/roleadd.html", gin.H{})
}

func (_this *HtmlWeb) permissionAdd(c *gin.Context) {
	c.HTML(http.StatusOK, "permission/permissionadd.html", gin.H{})
}

func (_this *HtmlWeb) e403(c *gin.Context) {
	c.HTML(http.StatusOK, "base/403.html", gin.H{})
}
