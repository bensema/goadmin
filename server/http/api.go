package http

import (
	"github.com/bensema/goadmin/config"
	"github.com/bensema/goadmin/model"
	"github.com/gin-gonic/gin"
	"image/gif"
	"image/png"
	"library/ecode"
	"net/http"
)

type Api struct{}

func (_this *Api) RegisterRoute(g *gin.RouterGroup) {
	g.GET("/api/rsa", _this.rsa)
	g.POST("/api/login", _this.login)
	g.GET("/api/captcha/img", _this.captchaImg)
	g.GET("/api/captcha/gif", _this.captchaGif)
}

func (_this *Api) rsa(c *gin.Context) {
	JSON(c, string(config.PublicKey), ecode.OK)
	return
}

func (_this *Api) login(c *gin.Context) {
	//verCode := c.PostForm("vercode")
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

	//capKey, _ := c.Cookie(CaptchaKey)
	//b, err := srv.CaptchaVerify(c, capKey, verCode)
	//if err == model.ErrNil {
	//	JSON(c, nil, errors.New("请刷新验证码试试"))
	//	return
	//}
	//if err != nil {
	//	log.Errorf(err.Error())
	//	JSON(c, nil, err)
	//	return
	//}
	//if !b {
	//	JSON(c, nil, errors.New("验证码错误"))
	//	return
	//}

	err := srv.AdminLogin(c, username, password)
	if err != nil {
		JSON(c, nil, err)
		return
	}
	u, _ := srv.GetAdminByName(c, username)
	_ = sessionSrv.GinSetSession(c, &model.GinSession{AdminId: u.Id, Name: u.Name})
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
