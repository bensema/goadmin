package contorller

import (
	"errors"
	"github.com/bensema/goadmin/global"
	"github.com/bensema/goadmin/model"
	"github.com/bensema/goadmin/server/http/internal"
	"github.com/bensema/library/cache/redis"
	"github.com/bensema/library/ecode"
	"github.com/bensema/library/log"
	"github.com/gin-gonic/gin"
	"image/gif"
	"image/png"
	"net/http"
)

type Api struct{}

func (_this *Api) RegisterRoute(g *gin.RouterGroup) {
	g.POST("/api/v1/login", _this.login)
	g.GET("/api/v1/captcha/img", _this.captchaImg)
	g.GET("/api/v1/captcha/gif", _this.captchaGif)
}

func (_this *Api) login(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	verCode := c.PostForm("vercode")
	if username == "" {
		internal.JSON(c, nil, ecode.UsernameIsEmpty)
		return
	}

	if password == "" {
		internal.JSON(c, nil, ecode.PasswordIsEmpty)
		return
	}
	capKey, _ := c.Cookie(internal.CaptchaKey)
	b, err := global.Srv.CaptchaVerify(c, capKey, verCode)
	if err == redis.ErrNil {
		internal.JSON(c, nil, errors.New("请刷新验证码试试"))
		return
	}
	if err != nil {
		log.Errorf(err.Error())
		internal.JSON(c, nil, err)
		return
	}
	if !b {
		internal.JSON(c, nil, errors.New("验证码错误"))
		return
	}

	session, err := global.Srv.AdminLogin(c, &model.AdminLoginReq{Username: username, Password: password})
	if err != nil {
		internal.JSON(c, nil, err)
		return
	}
	c.SetCookie(internal.AdminSession, session, 60*60*8, "/", "", false, true)
	internal.JSON(c, nil, ecode.OK)
	return

}

func (_this *Api) captchaImg(c *gin.Context) {
	data, code := global.Srv.CaptchaImg(c)
	v, err := global.Srv.SetCaptchaCache(c, code)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.SetCookie(internal.CaptchaKey, v, 60*3, "/", "", false, true)
	png.Encode(c.Writer, data)
}

func (_this *Api) captchaGif(c *gin.Context) {
	data, code := global.Srv.CaptchaGif(c)
	v, err := global.Srv.SetCaptchaCache(c, code)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.SetCookie(internal.CaptchaKey, v, 60*3, "/", "", false, true)
	gif.EncodeAll(c.Writer, data)
}
