package http

import (
	"github.com/bensema/gcurd"
	"github.com/bensema/goadmin/utils"
	"github.com/gin-gonic/gin"
)

func prepareReq(c *gin.Context, req *gcurd.Request) *gcurd.Request {
	if orderByFiled, b := c.GetQuery("order_by_filed"); b {
		req.OrderBy.Filed = orderByFiled
	}
	if orderByDirection, b := c.GetQuery("order_by_direction"); b {
		req.OrderBy.Direction = orderByDirection
	}
	if num, b := c.GetQuery("num"); b {
		req.Pagination.Num = utils.GetInt(num)
	}
	if size, b := c.GetQuery("size"); b {
		req.Pagination.Size = utils.GetInt(size)
	}
	req.Pagination.Verify()
	return req
}
