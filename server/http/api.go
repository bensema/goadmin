package http

import (
	"errors"
	"github.com/bensema/goadmin/config"
	"github.com/bensema/goadmin/model"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"image/gif"
	"image/png"
	"library/ecode"
	"net/http"
)

type Api struct{}

func (_this *Api) RegisterRoute(g *gin.RouterGroup) {
	g.GET("/api/v1/rsa", _this.rsa)
	g.POST("/api/v1/login", _this.login)
	g.GET("/api/v1/captcha/img", _this.captchaImg)
	g.GET("/api/v1/captcha/gif", _this.captchaGif)
}

func (_this *Api) rsa(c *gin.Context) {
	JSON(c, string(config.PublicKey), ecode.OK)
	return
}

func (_this *Api) login(c *gin.Context) {
	verCode := c.PostForm("vercode")
	username := c.PostForm("username")
	password := c.PostForm("password")

	if username == "" {
		JSON(c, nil, ecode.UsernameIsEmpty)
		return
	}

	if password == "" {
		JSON(c, nil, ecode.PasswordIsEmpty)
		return
	}

	capKey, _ := c.Cookie(CaptchaKey)
	b, err := srv.CaptchaVerify(c, capKey, verCode)
	if err == model.ErrNil {
		JSON(c, nil, errors.New("请刷新验证码试试"))
		return
	}
	if err != nil {
		log.Errorf(err.Error())
		JSON(c, nil, err)
		return
	}
	if !b {
		JSON(c, nil, errors.New("验证码错误"))
		return
	}

	session, err := srv.AdminLogin(c, username, password)
	if err != nil {
		JSON(c, nil, err)
		return
	}
	c.SetCookie(AdminSession, session, 60*60*8, "/", "", false, true)
	JSON(c, nil, ecode.OK)
	return

}

func (_this *Api) captchaImg(c *gin.Context) {
	data, code := srv.CaptchaImg(c)
	v, err := srv.SetCaptchaCache(c, code)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.SetCookie(CaptchaKey, v, 60*3, "/", "", false, true)
	png.Encode(c.Writer, data)
}

func (_this *Api) captchaGif(c *gin.Context) {
	data, code := srv.CaptchaGif(c)
	v, err := srv.SetCaptchaCache(c, code)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.SetCookie(CaptchaKey, v, 60*3, "/", "", false, true)
	gif.EncodeAll(c.Writer, data)
}
