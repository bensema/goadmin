package http

import (
	"fmt"
	"github.com/bensema/goadmin/config"
	"github.com/bensema/goadmin/global"
	"github.com/bensema/goadmin/server/http/contorller"
	"github.com/bensema/goadmin/server/http/middleware"
	"github.com/bensema/goadmin/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Init(c *config.Config, s *service.Service) {
	global.Srv = s
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
		new(contorller.HtmlWeb).RegisterRoute(web)
	}
	webAuth := e.Group("/", middleware.AuthAdmin())
	{
		new(contorller.HtmlWebAuth).RegisterRoute(webAuth)
	}
	api := e.Group("/")
	{
		new(contorller.Api).RegisterRoute(api)
	}
	apiAuth := e.Group("/", middleware.AuthApi())
	{
		new(contorller.ApiAuth).RegisterRoute(apiAuth)
	}

}

func webResources(e *gin.Engine, c *config.Config) {
	e.LoadHTMLGlob(fmt.Sprintf("%s/%s", c.Web.Dir, c.Web.Template))
	e.Static("/static", fmt.Sprintf("%s/%s", c.Web.Dir, c.Web.Static))
}

func Index(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{})
}
