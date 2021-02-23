package contorller

import (
	"github.com/bensema/goadmin/server/http/internal"
	"github.com/gin-gonic/gin"
	"net/http"
)

type HtmlWeb struct{}

func (_this *HtmlWeb) RegisterRoute(g *gin.RouterGroup) {
	g.GET("/login", _this.login)
	g.GET("/logout", _this.login)
	g.GET("/admin/form", _this.adminForm)
	g.GET("/admin/add", _this.adminAdd)
	g.GET("/role/add", _this.roleAdd)
	g.GET("/role/form", _this.roleForm)

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
	_, err := c.Cookie(internal.AdminSession)
	if err != nil {

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
