package controller

import (
	"github.com/bensema/goadmin/server/http/internal"
	"github.com/bensema/goadmin/service"
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

	g.GET("/advertise/add", _this.advertiseAdd)
	g.GET("/advertise/form", _this.advertiseForm)

	g.GET("/announcement/add", _this.announcementAdd)
	g.GET("/announcement/form", _this.announcementForm)

	g.GET("/game/add", _this.gameAdd)
	g.GET("/game/form", _this.gameForm)

	g.GET("/game_result/add", _this.gameResultAdd)
	g.GET("/game_result/form", _this.gameResultForm)
	g.GET("/game_result/detail", _this.gameResultDetail)

	g.GET("/admin/info", _this.adminInfo)
	g.GET("/admin/password", _this.adminPassword)

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
		service.Srv.DeleteAdminSessionCache(c, adminSession)
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

func (_this *HtmlWeb) advertiseAdd(c *gin.Context) {
	c.HTML(http.StatusOK, "advertise/advertiseadd.html", gin.H{})
}

func (_this *HtmlWeb) advertiseForm(c *gin.Context) {
	c.HTML(http.StatusOK, "advertise/advertiseform.html", gin.H{})
}

func (_this *HtmlWeb) announcementAdd(c *gin.Context) {
	c.HTML(http.StatusOK, "announcement/announcementadd.html", gin.H{})
}

func (_this *HtmlWeb) announcementForm(c *gin.Context) {
	c.HTML(http.StatusOK, "announcement/announcementform.html", gin.H{})
}

func (_this *HtmlWeb) gameAdd(c *gin.Context) {
	c.HTML(http.StatusOK, "game/game_add.html", gin.H{})
}

func (_this *HtmlWeb) gameForm(c *gin.Context) {
	c.HTML(http.StatusOK, "game/game_form.html", gin.H{})
}

func (_this *HtmlWeb) gameResultAdd(c *gin.Context) {
	c.HTML(http.StatusOK, "game_result/game_result_add.html", gin.H{})
}

func (_this *HtmlWeb) gameResultForm(c *gin.Context) {
	c.HTML(http.StatusOK, "game_result/game_result_form.html", gin.H{})
}

func (_this *HtmlWeb) gameResultDetail(c *gin.Context) {
	c.HTML(http.StatusOK, "game_result/game_result_detail.html", gin.H{})
}

func (_this *HtmlWeb) adminInfo(c *gin.Context) {
	c.HTML(http.StatusOK, "admin/info.html", gin.H{})
}

func (_this *HtmlWeb) adminPassword(c *gin.Context) {
	c.HTML(http.StatusOK, "admin/password.html", gin.H{})
}
