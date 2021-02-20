package contorller

import (
	"github.com/bensema/goadmin/global"
	"github.com/bensema/goadmin/server/http/internal"
	"github.com/gin-gonic/gin"
)

type ApiAuth struct{}

func (_this *ApiAuth) RegisterRoute(g *gin.RouterGroup) {
	g.GET("/api/v1/menu", _this.menu)
}

func (_this *ApiAuth) menu(c *gin.Context) {
	session, _ := c.Cookie(internal.AdminSession)
	adminSession, _ := global.Srv.GetAdminSessionCache(c, session)
	menus, err := global.Srv.FindAdminMenu(c, adminSession.UserId)
	internal.JSON(c, menus, err)
}
