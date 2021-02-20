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
}

func (_this *HtmlWeb) login(c *gin.Context) {
	_, err := c.Cookie(internal.AdminSession)
	if err != nil {
		c.HTML(http.StatusOK, "base/login-v2.html", gin.H{})
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
