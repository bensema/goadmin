package http

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type HtmlWebAuth struct{}

func (_this *HtmlWebAuth) RegisterRoute(g *gin.RouterGroup) {
	g.GET("/", _this.index)
	g.GET("/admin", _this.commonV2)
	g.GET("/role", _this.commonV2)
	g.GET("/permission", _this.commonV2)
	g.GET("/resources", _this.commonV2)
	g.GET("/advertise", _this.commonV2)
	g.GET("/announcements", _this.commonV2)
	g.GET("/game", _this.commonV2)
	g.GET("/game_result", _this.commonV2)
	g.GET("/home/dashboard", _this.common)
	g.GET("/log/login", _this.logLogin)
	g.GET("/log/operation", _this.logOperation)
}

func (_this *HtmlWebAuth) index(c *gin.Context) {
	//operatorInfo, _ := service.GetAdminFromContext(c)
	//c.HTML(http.StatusOK, "base/index.html", gin.H{"name": operatorInfo.Name})
	c.HTML(http.StatusOK, "base/index.html", gin.H{"name": "todo"})
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

func (_this *HtmlWebAuth) commonV2(c *gin.Context) {
	s := c.Request.URL.String()[1:]
	p := fmt.Sprintf("%s/%s.html", s, s)
	c.HTML(http.StatusOK, p, gin.H{})
}
