package internal

import (
	"github.com/bensema/goadmin/model"
	"github.com/bensema/library/ecode"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type res struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	T       int64       `json:"t"`
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
		T:       time.Now().Unix(),
	})
}

func AdminJSON(c *gin.Context, resp *model.AdminApiReply, err error) {
	code := http.StatusOK
	bCode := ecode.Cause(err)

	if err != nil {
		resp.Code = bCode.Code()
		resp.Message = bCode.Message() + " from admin"
	}

	c.JSON(code, resp)
}
