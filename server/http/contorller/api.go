package contorller

import (
	"fmt"
	"github.com/bensema/goadmin/ecode"
	"github.com/bensema/goadmin/global"
	"github.com/bensema/goadmin/model"
	"github.com/bensema/goadmin/server/http/internal"
	"github.com/gin-gonic/gin"
)

const (
	CaptchaKey = "captcha_key"
)

type Api struct{}

func (_this *Api) RegisterRoute(g *gin.RouterGroup) {
	g.POST("/api/v1/login", _this.login)
}

func (_this *Api) login(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	if username == "" {
		internal.JSON(c, nil, ecode.UsernameIsEmpty)
		return
	}

	if password == "" {
		internal.JSON(c, nil, ecode.PasswordIsEmpty)
		return
	}
	fmt.Println(username, password)
	session, err := global.Srv.AdminLogin(c, &model.AdminLoginReq{Username: username, Password: password})
	if err != nil {
		internal.JSON(c, nil, err)
		return
	}
	c.SetCookie(internal.AdminSession, session, 60*60*8, "/", "", false, true)
	internal.JSON(c, nil, ecode.OK)
	return

}
