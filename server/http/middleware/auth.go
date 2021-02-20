package middleware

import (
	"github.com/bensema/goadmin/ecode"
	"github.com/bensema/goadmin/global"
	"github.com/bensema/goadmin/server/http/internal"
	"github.com/gin-gonic/gin"
	"net/http"
)

// 登陆
func AuthAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		sid, err := c.Cookie(internal.AdminSession)
		if err != nil {
			err = ecode.NoLogin
			c.Redirect(http.StatusFound, "/login")
			c.Abort()
			return
		}
		_, err = global.Srv.GetAdminSessionCache(c, sid)
		if err != nil {
			err = ecode.AccessTokenExpires
			c.Redirect(http.StatusFound, "/login")
			c.Abort()
			return
		}
		c.Next()
		return
	}
}

// 登陆
func AuthApi() gin.HandlerFunc {
	return func(c *gin.Context) {
		sid, err := c.Cookie(internal.AdminSession)
		if err != nil {
			err = ecode.NoLogin
			internal.JSON(c, nil, err)
			c.Abort()
			return
		}
		_, err = global.Srv.GetAdminSessionCache(c, sid)
		if err != nil {
			err = ecode.AccessTokenExpires
			internal.JSON(c, nil, err)
			c.Abort()
			return
		}
		c.Next()
		return
	}
}
