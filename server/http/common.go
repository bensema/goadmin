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
	if page, b := c.GetQuery("page"); b {
		req.Pagination.Page = utils.GetInt(page)
	}
	if pageSize, b := c.GetQuery("page_size"); b {
		req.Pagination.PageSize = utils.GetInt(pageSize)
	}
	req.Pagination.Verify()
	return req
}
