package http

import (
	"github.com/bensema/goadmin/model"
	"github.com/gin-gonic/gin"
	"library/ecode"
)

func auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		adminSession := &model.GinSession{}
		err := sessionSrv.GinLoadSession(c, adminSession)
		if err != nil {
			err = ecode.AccessTokenExpires
			JSON(c, nil, err)
			c.Abort()
			return
		}
		c.Set("id", adminSession.AdminId)
		c.Set("name", adminSession.Name)

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
