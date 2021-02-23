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
	g.GET("/home/dashboard", _this.common)
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

func (_this *HtmlWebAuth) common(c *gin.Context) {
	name := fmt.Sprintf("%s.html", c.Request.URL.String())
	fmt.Println(name)
	c.HTML(http.StatusOK, name, gin.H{})
}
