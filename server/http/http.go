package http

import (
	"fmt"
	"github.com/bensema/goadmin/config"
	"github.com/bensema/goadmin/server/http/controller"
	"github.com/bensema/goadmin/server/http/middleware"
	"github.com/bensema/goadmin/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Init(c *config.Config, s *service.Service) {

	service.Srv = s

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
		new(controller.HtmlWeb).RegisterRoute(web)
	}
	webAuth := e.Group("/", middleware.PermitWeb())
	{
		new(controller.HtmlWebAuth).RegisterRoute(webAuth)
	}
	api := e.Group("/")
	{
		new(controller.Api).RegisterRoute(api)
	}
	apiAuth := e.Group("/", middleware.PermitApi())
	{
		new(controller.ApiAuth).RegisterRoute(apiAuth)
	}
}

func webResources(e *gin.Engine, c *config.Config) {
	e.LoadHTMLGlob(c.Web.Template)
	e.Static("/static", c.Web.Static)
}

func Index(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{})
}
