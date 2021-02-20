package service

import (
	"fmt"
	"github.com/bensema/goadmin/ecode"
	"github.com/bensema/goadmin/model"
	"github.com/bensema/goadmin/utils"
	xtime "github.com/bensema/library/time"
	"github.com/gin-gonic/gin"
	"github.com/mssola/user_agent"
	"time"
)

func (s *Service) GetAdminById(c *gin.Context, id int) (*model.Admin, error) {
	return s.dao.GetAdminById(c, id)
}

func (s *Service) AdminLogin(c *gin.Context, arg *model.AdminLoginReq) (string, error) {
	au, err := s.dao.GetAdminByName(c, arg.Username)
	if err != nil {
		fmt.Println(err)
		return "", ecode.UsernameOrPasswordErr
	}
	b := utils.ComparePasswords(au.Password, arg.Password)
	if !b {
		return "", ecode.UsernameOrPasswordErr
	}
	adminSessionKey := utils.RandomString(40)
	adminSession := model.AdminSession{
		UserId: au.Id,
		Name:   au.Name,
	}
	err = s.dao.SetAdminSessionCache(c, adminSessionKey, &adminSession)
	ipInfo, _ := s.Ip2Region.BtreeSearch(c.ClientIP())
	ua := user_agent.New(c.Request.UserAgent())
	b1, bv := ua.Browser()
	loginLog := &model.LogAdminLogin{
		AdminId:    au.Id,
		Name:       au.Name,
		Location:   fmt.Sprintf("%s %s", ipInfo.Province, ipInfo.City),
		Os:         ua.OS(),
		Browser:    b1 + bv,
		UserAgent:  ua.UA(),
		Url:        c.FullPath(),
		Result:     1,
		Ip:         c.ClientIP(),
		RecordTime: xtime.Time(time.Now().Unix()),
		Remark:     "",
	}
	_ = s.dao.LogAdminLogin(c, loginLog)

	return adminSessionKey, err
}

func (s *Service) FindAdminMenu(c *gin.Context, uid int) ([]*model.Menu, error) {
	var res []*model.Menu
	temp := map[int]struct{}{}
	farr := &model.FindAdminRoleReq{}
	farr.AdminId = uid
	adminRoles, err := s.dao.FindAdminRole(c, farr)
	if err != nil {
		return nil, err
	}
	for _, adminRole := range adminRoles {
		frpr := &model.FindRolePermissionReq{}
		frpr.RoleId = adminRole.RoleId
		rolePermissions, err := s.dao.FindRolePermission(c, frpr)
		if err != nil {
			continue
		}
		for _, rolePermission := range rolePermissions {
			fpmr := &model.FindPermissionMenuReq{}
			fpmr.PermissionId = rolePermission.PermissionId
			permissionMenus, err := s.dao.FindPermissionMenu(c, fpmr)
			if err != nil {
				continue
			}
			for _, permissionMenu := range permissionMenus {
				menu, err := s.dao.GetMenuById(c, permissionMenu.MenuId)
				if err != nil {
					continue
				}
				//去重
				if _, ok := temp[menu.Id]; !ok {
					temp[menu.Id] = struct{}{}
					res = append(res, menu)
				}
			}
		}
	}
	return res, err
}
