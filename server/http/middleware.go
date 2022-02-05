package http

import (
	"github.com/bensema/goadmin/model"
	"github.com/gin-gonic/gin"
	"library/ecode"
	"net/http"
)

// 登陆
func PermitWeb() gin.HandlerFunc {
	return func(c *gin.Context) {
		sid, err := c.Cookie(AdminSession)
		if err != nil {
			err = ecode.NoLogin
			c.Redirect(http.StatusFound, "/login")
			c.Abort()
			return
		}
		adminSession := &model.AdminSession{}
		err = srv.GetAdminSessionCache(c, sid, adminSession)
		if err != nil {
			err = ecode.AccessTokenExpires
			c.Redirect(http.StatusFound, "/login")
			c.Abort()
			return
		}
		// root无任何限制
		if adminSession.Name == Root {
			c.Next()
			return
		}

		err = srv.PermitWeb(c, adminSession.AdminId)
		if err != nil {
			c.Redirect(http.StatusFound, "/403")
			c.Abort()
			return
		}
		c.Next()
		return
	}
}

func PermitApi() gin.HandlerFunc {
	return func(c *gin.Context) {
		sid, err := c.Cookie(AdminSession)
		if err != nil {
			err = ecode.NoLogin
			JSON(c, nil, err)
			c.Abort()
			return
		}
		adminSession := &model.AdminSession{}
		err = srv.GetAdminSessionCache(c, sid, adminSession)
		if err != nil {
			err = ecode.AccessTokenExpires
			JSON(c, nil, err)
			c.Abort()
			return
		}
		c.Set("admin_id", adminSession.AdminId)
		c.Set("admin_name", adminSession.Name)

		// root无任何限制
		if adminSession.Name == Root {
			c.Next()
			return
		}

		err = srv.PermitAPI(c, adminSession.AdminId)
		if err != nil {
			JSON(c, nil, err)
			c.Abort()
			return
		}

		c.Next()
		return
	}
}
