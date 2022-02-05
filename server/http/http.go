package http

import (
	"fmt"
	"github.com/bensema/goadmin/config"
	"github.com/bensema/goadmin/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

var (
	srv *service.Service
)

func Init(c *config.Config, s *service.Service) {
	srv = s
	engine := gin.Default()
	webResources(engine, c)
	router(engine)
	go func() {
		_ = engine.Run(fmt.Sprintf(":%d", c.Web.Port))
	}()
}

func router(e *gin.Engine) {
	web := e.Group("/")
	{
		new(HtmlWeb).RegisterRoute(web)
	}
	webAuth := e.Group("/", PermitWeb())
	{
		new(HtmlWebAuth).RegisterRoute(webAuth)
	}
	api := e.Group("/")
	{
		new(Api).RegisterRoute(api)
	}
	apiAuth := e.Group("/", PermitApi())
	{
		new(ApiAuth).RegisterRoute(apiAuth)
	}
}

func webResources(e *gin.Engine, c *config.Config) {
	e.LoadHTMLGlob(c.Web.Template)
	e.Static("/static", c.Web.Static)
}

func Index(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{})
}
