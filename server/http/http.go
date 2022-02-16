package http

import (
	"fmt"
	"github.com/bensema/goadmin/config"
	"github.com/bensema/goadmin/service"
	"github.com/gin-gonic/gin"
	"library/session"
)

var (
	srv        *service.Service
	sessionSrv *session.Session
)

func Init(c *config.Config, s *service.Service) {
	srv = s
	sessionSrv = session.New(c.Redis)
	engine := gin.Default()
	router(engine)
	go func() {
		_ = engine.Run(fmt.Sprintf(":%d", c.Port))
	}()
}

func router(e *gin.Engine) {

	api := e.Group("/")
	{
		new(Api).RegisterRoute(api)
	}
	apiAuth := e.Group("/", auth())
	{
		new(ApiAuth).RegisterRoute(apiAuth)
	}
}
