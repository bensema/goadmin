package service

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/bensema/gcurd"
	"github.com/bensema/goadmin/model"
	"github.com/bensema/goadmin/utils"
	"github.com/gin-gonic/gin"
	"github.com/mssola/user_agent"
	"library/ecode"
	"library/xtime"
	"time"
)

func (s *Service) AdminLogin(c *gin.Context, username string, password string) error {
	au, err := s.dao.GetAdminByName(c, username)
	if err != nil {
		return ecode.UsernameOrPasswordErr
	}
	b := utils.ComparePasswords(au.Password, password)
	if !b {
		return ecode.UsernameOrPasswordErr
	}
	ipInfo, _ := s.Ip2Region.BtreeSearch(c.ClientIP())
	ua := user_agent.New(c.Request.UserAgent())
	b1, bv := ua.Browser()
	loginLog := &model.LogAdminLogin{
		AdminId:   au.Id,
		Name:      au.Name,
		Location:  fmt.Sprintf("%s %s", ipInfo.Province, ipInfo.City),
		Os:        ua.OS(),
		Browser:   b1 + bv,
		UserAgent: ua.UA(),
		Url:       c.FullPath(),
		Result:    1,
		Ip:        c.ClientIP(),
		RecordAt:  xtime.Time(time.Now().Unix()),
		Remark:    "",
	}
	_, _ = s.dao.CreateLogAdminLogin(c, loginLog)
	return err
}

func (s *Service) FindAdminMenu(c *gin.Context, adminId int) ([]*model.Menu, error) {
	var res []*model.Menu
	adminRoles, err := s.dao.FindAdminRole(c, []*gcurd.WhereValue{gcurd.EQ("admin_id", adminId)})
	if err != nil {
		return nil, err
	}
	for _, adminRole := range adminRoles {
		roleMenus, err := s.dao.FindRoleMenu(c, []*gcurd.WhereValue{gcurd.EQ("role_id", adminRole.RoleId)})
		if err != nil {
			continue
		}
		for _, roleMenu := range roleMenus {
			menu, err := s.dao.GetMenu(c, roleMenu.MenuId)
			if err != nil {
				continue
			}
			res = append(res, menu)
		}
	}
	return res, err
}

func (s *Service) FindAdminPageV2(c *gin.Context, req *gcurd.Request) (reply *model.PageReply[*model.AdminV1], err error) {
	var count int
	var dataTmp []*model.Admin
	var data []*model.AdminV1

	if count, err = s.dao.PageTotalAdmin(c, req); err != nil {
		return
	}
	if count <= 0 {
		return
	}
	if dataTmp, err = s.dao.PageFindAdmin(c, req); err != nil {
		return
	}
	for _, d := range dataTmp {
		rls := make([]*model.Role, 0)
		adminRoles, _ := s.dao.FindAdminRole(c, []*gcurd.WhereValue{gcurd.EQ("admin_id", d.Id)})

		for _, adminRole := range adminRoles {
			roles, _ := s.dao.FindRole(c, []*gcurd.WhereValue{gcurd.EQ("id", adminRole.RoleId)})
			for _, role := range roles {
				rls = append(rls, role)
			}
		}

		data = append(data, &model.AdminV1{
			Id:        d.Id,
			Name:      d.Name,
			Status:    d.Status,
			CreatedAt: d.CreatedAt,
			UpdatedAt: d.UpdatedAt,
			Roles:     rls,
		})
	}
	reply = &model.PageReply[*model.AdminV1]{}
	reply.Rows = data
	reply.RowsTotal = count
	reply.Page = req.Pagination.Page
	reply.PageSize = req.Pagination.PageSize
	return
}

// GetAdminV1 获取管理员信息过滤密码，添加角色
func (s *Service) GetAdminV1(c *gin.Context, adminId int) (*model.AdminV1, error) {
	var data *model.AdminV1
	dataTmp, err := s.dao.GetAdmin(c, adminId)
	if err != nil {
		fmt.Println(err)
	}
	if err == sql.ErrNoRows {
		return nil, errors.New("账户不存在")
	}

	if err != nil {
		return nil, err
	}

	rls := make([]*model.Role, 0)
	adminRoles, _ := s.dao.FindAdminRole(c, []*gcurd.WhereValue{gcurd.EQ("admin_id", adminId)})

	for _, adminRole := range adminRoles {
		roles, _ := s.dao.FindRole(c, []*gcurd.WhereValue{gcurd.EQ("id", adminRole.RoleId)})
		for _, role := range roles {
			rls = append(rls, role)
		}
	}

	data = &model.AdminV1{
		Id:        dataTmp.Id,
		Name:      dataTmp.Name,
		Status:    dataTmp.Status,
		CreatedAt: dataTmp.CreatedAt,
		UpdatedAt: dataTmp.UpdatedAt,
		Remark:    dataTmp.Remark,
		Roles:     rls,
	}

	return data, err
}

func (s *Service) FindAllRole(c *gin.Context) (reply []*model.Role, err error) {
	return s.dao.FindRole(c, nil)
}

func (s *Service) UpdateAdminV2(c *gin.Context, info *model.UpdateAdmin, filed []string) error {
	var content string
	aInfo, err := s.dao.GetAdmin(c, info.AdminId)
	if err != nil {
		fmt.Println(err)
		return errors.New("Not found admin")
	}
	for _, v := range filed {
		switch v {
		case "status":
			content += fmt.Sprintf("状态:%s;", info.Status)
			_ = s.dao.UpdateAdmin(c, info.AdminId, "status", info.Status)
		case "remark":
			content += fmt.Sprintf("备注:%s;", info.Remark)
			_ = s.dao.UpdateAdmin(c, info.AdminId, "remark", info.Remark)
		case "roles":
			for _, role := range info.Roles {
				_, err := s.dao.GetRoleById(c, role)
				if err != nil {
					return errors.New("角色不存在")
				}
			}
			err = s.dao.DeleteAdminRoleByAdminId(c, info.AdminId)
			if err != nil {
				return err
			}
			for _, role := range info.Roles {
				rInfo, err := s.dao.GetRoleById(c, role)
				if err != nil {
					return errors.New("角色不存在")
				}

				content += fmt.Sprintf("角色编号:%d;角色:%s;", rInfo.Id, rInfo.Name)
				_, _ = s.dao.CreateAdminRole(c, &model.AdminRole{
					AdminId: info.AdminId,
					RoleId:  rInfo.Id,
				})
			}
		}
	}

	cc := fmt.Sprintf("修改管理员:账户:%s;账户编号:%d;%s", aInfo.Name, aInfo.Id, content)
	return s.logAction(c, "update_admin", cc, 1)

}

func (s *Service) DeleteAdminV2(c *gin.Context, adminId int) error {

	var content string
	aInfo, err := s.dao.GetAdmin(c, adminId)
	if err != nil {
		return errors.New("账户不存在")
	}

	err = s.dao.DeleteAdmin(c, adminId)
	err = s.dao.DeleteAdminRoleByAdminId(c, adminId)

	cc := fmt.Sprintf("删除管理员:账户:%s;账户编号:%d;%s", aInfo.Name, aInfo.Id, content)
	return s.logAction(c, "delete_admin", cc, 1)
}

func (s *Service) AddAdmin(c *gin.Context, info *model.AddAdmin) error {
	if err := utils.CheckNameLegal(info.Name); err != nil {
		return err
	}
	if err := utils.CheckPasswordLegal(info.Password); err != nil {
		return err
	}

	user := &model.Admin{}
	_, err := s.dao.GetAdminByName(c, info.Name)
	if err != sql.ErrNoRows {
		return errors.New("账户已存在")
	}

	hashPwd, err := utils.HashAndSalt(info.Password)
	if err != nil {
		return ecode.PasswordEncodeErr
	}

	//user.AdminId = utils.RandInt(7)
	user.Name = info.Name
	user.Password = hashPwd
	user.Status = info.Status
	user.CreatedAt = xtime.Time(time.Now().Unix())
	user.UpdatedAt = xtime.Time(time.Now().Unix())
	res, err := s.dao.CreateAdmin(c, user)
	if err != nil {
		return err
	}
	id, _ := res.LastInsertId()
	uInfo, err := s.dao.GetAdminByName(c, info.Name)
	if err != nil {
		return err
	}
	for _, role := range info.Roles {
		_, _ = s.dao.CreateAdminRole(c, &model.AdminRole{
			AdminId: uInfo.Id,
			RoleId:  role,
		})
	}

	cc := fmt.Sprintf("添加管理员:账户:%s;账户编号:%d;", info.Name, id)
	return s.logAction(c, "add_admin", cc, 1)
}

// 角色

func (s *Service) FindRolePageV1(c *gin.Context, req *gcurd.Request) (reply *model.PageReply[*model.RoleInfo], err error) {
	var count int
	var dataTmp []*model.Role
	var data []*model.RoleInfo
	if count, err = s.dao.PageTotalRole(c, req); err != nil {
		return
	}
	if count <= 0 {
		return
	}
	if dataTmp, err = s.dao.FindRole(c, req.Where); err != nil {
		return
	}
	for _, d := range dataTmp {
		var ms []*model.Menu
		var as []*model.Api
		roleMenus, _ := s.dao.FindRoleMenu(c, []*gcurd.WhereValue{gcurd.EQ("role_id", d.Id)})
		for _, roleMenu := range roleMenus {
			menus, _ := s.dao.FindMenu(c, []*gcurd.WhereValue{gcurd.EQ("id", roleMenu.MenuId)})
			ms = append(ms, menus...)
		}
		roleApis, _ := s.dao.FindRoleApi(c, []*gcurd.WhereValue{gcurd.EQ("role_id", d.Id)})
		for _, roleApi := range roleApis {
			apis, _ := s.dao.FindApi(c, []*gcurd.WhereValue{gcurd.EQ("id", roleApi.ApiId)})
			as = append(as, apis...)
		}

		data = append(data, &model.RoleInfo{
			Id:    d.Id,
			Name:  d.Name,
			Menus: ms,
			Apis:  as,
		})
	}
	reply.Rows = data
	reply.RowsTotal = count
	reply.Page = req.Pagination.Page
	reply.PageSize = req.Pagination.PageSize
	return
}

func (s *Service) GetRoleInfo(c *gin.Context, roleId int) (*model.RoleInfo, error) {
	var data *model.RoleInfo
	dataTmp, err := s.dao.GetRoleById(c, roleId)
	if err == sql.ErrNoRows {
		return nil, errors.New("角色不存在")
	}

	if err != nil {
		return nil, err
	}

	var ms []*model.Menu
	var as []*model.Api
	roleMenus, _ := s.dao.FindRoleMenu(c, []*gcurd.WhereValue{gcurd.EQ("role_id", roleId)})
	for _, roleMenu := range roleMenus {
		menus, _ := s.dao.FindMenu(c, []*gcurd.WhereValue{gcurd.EQ("id", roleMenu.MenuId)})
		ms = append(ms, menus...)
	}
	roleApis, _ := s.dao.FindRoleApi(c, []*gcurd.WhereValue{gcurd.EQ("role_id", roleId)})
	for _, roleApi := range roleApis {
		apis, _ := s.dao.FindApi(c, []*gcurd.WhereValue{gcurd.EQ("id", roleApi.ApiId)})
		as = append(as, apis...)
	}

	data = &model.RoleInfo{
		Id:    dataTmp.Id,
		Name:  dataTmp.Name,
		Menus: ms,
		Apis:  as,
	}

	return data, err
}

func (s *Service) FindAllMenu(c *gin.Context) (reply []*model.Menu, err error) {
	return s.dao.FindMenu(c, nil)
}

func (s *Service) FindAllApi(c *gin.Context) (reply []*model.Api, err error) {
	return s.dao.FindApi(c, nil)
}

func (s *Service) AddRole(c *gin.Context, role *model.Role) error {

	aInfo, err := s.dao.GetRoleByName(c, role.Name)
	if err != sql.ErrNoRows {
		return errors.New("角色名已存在")
	}
	res, err := s.dao.CreateRole(c, role)
	if err != nil {
		return err
	}
	id, _ := res.LastInsertId()

	content := fmt.Sprintf("添加角色:角色:%s;角色编号:%d;备注：%s", aInfo.Name, id, aInfo.Remark)
	return s.logAction(c, "add_role", content, 1)
}

func (s *Service) DeleteRoleV1(c *gin.Context, id int) error {

	var content string
	roleInfo, err := s.dao.GetRoleById(c, id)
	if err != nil {
		return errors.New("角色不存在")
	}

	ss := []*gcurd.WhereValue{gcurd.EQ("role_id", id)}
	fmt.Println(ss)

	var wvs []*gcurd.WhereValue
	wvs = append(wvs, gcurd.EQ("role_id", id))
	arl, err := s.dao.FindAdminRole(c, wvs)
	if err != nil {
		return errors.New("获取用户-角色失败")
	}
	if len(arl) > 0 {
		return errors.New("还有账户在使用此角色")
	}

	err = s.dao.DeleteRole(c, id)
	err = s.dao.DeleteRoleMenuByRoleId(c, id)
	err = s.dao.DeleteRoleApiByRoleId(c, id)

	return s.logAction(c, "delete_role", fmt.Sprintf("删除角色:角色:%s;角色编号:%d;%s", roleInfo.Name, roleInfo.Id, content), 1)
}

func (s *Service) logAction(c *gin.Context, opCode string, content string, result int) error {
	return logAction(c, s.dao.DB(), opCode, content, result)
}

func (s *Service) GetAdminByName(c *gin.Context, name string) (*model.Admin, error) {
	return s.dao.GetAdminByName(c, name)
}
