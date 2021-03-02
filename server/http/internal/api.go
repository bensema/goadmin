package internal

import (
	"github.com/bensema/library/ecode"
	"github.com/gin-gonic/gin"
	"net/http"
)

type res struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func JSON(c *gin.Context, data interface{}, err error) {
	code := http.StatusOK
	bCode := ecode.Cause(err)
	if data == nil {
		data = gin.H{}
	}
	c.JSON(code, res{
		Code:    bCode.Code(),
		Message: bCode.Message(),
		Data:    data,
	})
}
