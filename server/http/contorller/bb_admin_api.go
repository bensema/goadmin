package contorller

import (
	"github.com/bensema/goadmin/global"
	"github.com/bensema/goadmin/server/http/internal"
	"github.com/gin-gonic/gin"
)

func (_this *ApiAuth) RegisterBBAdminRoute(g *gin.RouterGroup) {
	g.GET("/api/v1/advertise/pages", _this.advertisePages) // 分页查询广告
}

// 返回
func (_this *ApiAuth) advertisePages(c *gin.Context) {
	reply, err := global.Srv.FindAdvertisePage(c)
	internal.AdminJSON(c, reply, err)
}
