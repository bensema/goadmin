package service

import (
	"github.com/bensema/goadmin/model"
	"github.com/gin-gonic/gin"
)

func (s *Service) GetAdminSessionCache(c *gin.Context, key string) (*model.AdminSession, error) {
	return s.dao.GetAdminSessionCache(c, key)
}
