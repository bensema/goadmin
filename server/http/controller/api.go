package controller

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"github.com/bensema/goadmin/global"
	"github.com/bensema/goadmin/model"
	"github.com/bensema/goadmin/server/http/internal"
	"github.com/bensema/library/cache/redis"
	"github.com/bensema/library/crypto"
	"github.com/bensema/library/ecode"
	"github.com/bensema/library/log"
	xtime "github.com/bensema/library/time"
	"github.com/gin-gonic/gin"
	"image/gif"
	"image/png"
	"net/http"
	"strings"
	"time"
)

type Api struct{}

func (_this *Api) RegisterRoute(g *gin.RouterGroup) {
	g.GET("/api/v1/rsa", _this.rsa)
	g.POST("/api/v1/login", _this.login)
	g.GET("/api/v1/captcha/img", _this.captchaImg)
	g.GET("/api/v1/captcha/gif", _this.captchaGif)
}

func (_this *Api) rsa(c *gin.Context) {
	internal.JSON(c, string(global.PublicKey), ecode.OK)
	return
}

func (_this *Api) login(c *gin.Context) {
	verCode := c.PostForm("vercode")
	ciphertext := c.PostForm("data")
	dd, err := base64.StdEncoding.DecodeString(ciphertext)
	if err != nil {
		internal.JSON(c, nil, ecode.MethodNotAllowed)
		return
	}
	var adminLogin model.AdminLoginV2
	jsonStream, err := crypto.RsaDecrypt(dd, global.PrivateKey)
	if err != nil {
		internal.JSON(c, nil, ecode.MethodNotAllowed)
		return
	}
	dec := json.NewDecoder(strings.NewReader(string(jsonStream)))
	err = dec.Decode(&adminLogin)
	//fmt.Println(adminLogin)
	if err != nil {
		internal.JSON(c, nil, ecode.MethodNotAllowed)
		return
	}

	if adminLogin.Username == "" {
		internal.JSON(c, nil, ecode.UsernameIsEmpty)
		return
	}

	if adminLogin.Password == "" {
		internal.JSON(c, nil, ecode.PasswordIsEmpty)
		return
	}

	st, _ := time.ParseDuration("-20s")
	if xtime.Time(time.Now().Add(st).UnixNano())/1e6 > adminLogin.T {
		internal.JSON(c, nil, ecode.Deadline)
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

	session, err := global.Srv.AdminLogin(c, &model.AdminLoginReq{Username: adminLogin.Username, Password: adminLogin.Password, AesKey: adminLogin.AesKey})
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
