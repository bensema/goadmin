package contorller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type HtmlWebAuth struct{}

func (_this *HtmlWebAuth) RegisterRoute(g *gin.RouterGroup) {
	g.GET("/", _this.index)
	g.GET("/admin", _this.admin)
	g.GET("/role", _this.role)
	g.GET("/permission", _this.permission)
	g.GET("/resources", _this.resources)
	g.GET("/home/dashboard", _this.common)
	g.GET("/log/login", _this.logLogin)
	g.GET("/log/operation", _this.logOperation)
}

func (_this *HtmlWebAuth) index(c *gin.Context) {
	c.HTML(http.StatusOK, "base/index.html", gin.H{})
}

func (_this *HtmlWebAuth) admin(c *gin.Context) {
	c.HTML(http.StatusOK, "admin/admin.html", gin.H{})
}

func (_this *HtmlWebAuth) role(c *gin.Context) {
	c.HTML(http.StatusOK, "role/role.html", gin.H{})
}

func (_this *HtmlWebAuth) permission(c *gin.Context) {
	c.HTML(http.StatusOK, "permission/permission.html", gin.H{})
}

func (_this *HtmlWebAuth) resources(c *gin.Context) {
	c.HTML(http.StatusOK, "resources/resources.html", gin.H{})
}

func (_this *HtmlWebAuth) logLogin(c *gin.Context) {
	c.HTML(http.StatusOK, "log/login.html", gin.H{})
}

func (_this *HtmlWebAuth) logOperation(c *gin.Context) {
	c.HTML(http.StatusOK, "log/operation.html", gin.H{})
}

func (_this *HtmlWebAuth) common(c *gin.Context) {
	name := fmt.Sprintf("%s.html", c.Request.URL.String())
	c.HTML(http.StatusOK, name, gin.H{})
}
