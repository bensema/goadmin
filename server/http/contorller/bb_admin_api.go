package contorller

import (
	"github.com/bensema/goadmin/global"
	"github.com/bensema/goadmin/server/http/internal"
	"github.com/gin-gonic/gin"
)

func (_this *ApiAuth) RegisterBBAdminRoute(g *gin.RouterGroup) {
	g.GET("/api/v1/advertise/pages", _this.advertisePages) // 分页查询广告
	g.POST("/api/v1/advertise/add", _this.advertiseAdd)    // 添加广告
	g.POST("/api/v1/advertise/del", _this.advertiseDel)    // 删除广告
}

func (_this *ApiAuth) advertisePages(c *gin.Context) {
	reply, err := global.Srv.FindAdvertisePage(c)
	internal.AdminJSON(c, reply, err)
}

func (_this *ApiAuth) advertiseAdd(c *gin.Context) {
	reply, err := global.Srv.AdvertiseAdd(c)
	internal.AdminJSON(c, reply, err)
}

func (_this *ApiAuth) advertiseDel(c *gin.Context) {
	reply, err := global.Srv.AdvertiseDel(c)
	internal.AdminJSON(c, reply, err)
}
