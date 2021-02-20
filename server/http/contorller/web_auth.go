package contorller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type HtmlWebAuth struct{}

func (_this *HtmlWebAuth) RegisterRoute(g *gin.RouterGroup) {
	g.GET("/", _this.index)
	g.GET("/admin", _this.admin)
}

func (_this *HtmlWebAuth) index(c *gin.Context) {
	c.HTML(http.StatusOK, "base/index.html", gin.H{})
}

func (_this *HtmlWebAuth) admin(c *gin.Context) {
	c.HTML(http.StatusOK, "html/admin.html", gin.H{})
}
