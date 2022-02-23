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
		Result:    model.SUCCESS,
		Ip:        c.ClientIP(),
		RecordAt:  xtime.Time(time.Now().Unix()),
		Remark:    "",
	}
	_, _ = s.dao.CreateLogAdminLogin(c, loginLog)
	return err
}

func (s *Service) FindAdminMenu(c *gin.Context, adminId int) ([]*model.Menu, error) {
	var res []*model.Menu
	temp := map[int]struct{}{}
	adminRoles, err := s.dao.FindAdminRole(c, []*gcurd.WhereValue{gcurd.EQ("admin_id", adminId)})
	if err != nil {
		return nil, err
	}
	for _, adminRole := range adminRoles {
		rolePermissions, err := s.dao.FindRolePermission(c, []*gcurd.WhereValue{gcurd.EQ("role_id", adminRole.RoleId)})
		if err != nil {
			continue
		}
		for _, rolePermission := range rolePermissions {
			permissionMenus, err := s.dao.FindPermissionMenu(c, []*gcurd.WhereValue{gcurd.EQ("permission_id", rolePermission.PermissionId)})
			if err != nil {
				continue
			}
			for _, permissionMenu := range permissionMenus {
				menu, err := s.dao.GetMenu(c, permissionMenu.MenuId)
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

func (s *Service) UpdateRoleV2(c *gin.Context, info *model.UpdateRole, filed []string) error {
	var content string
	aInfo, err := s.dao.GetRole(c, info.RoleId)
	if err != nil {
		fmt.Println(err)
		return errors.New("not found role")
	}
	for _, v := range filed {
		switch v {
		case "name":
			content += fmt.Sprintf("name:%s;", info.Name)
			_ = s.dao.UpdateRole(c, info.RoleId, "name", info.Name)
		case "remark":
			content += fmt.Sprintf("remark:%s;", info.Remark)
			_ = s.dao.UpdateRole(c, info.RoleId, "remark", info.Remark)
		case "permissions":
			for _, permission := range info.Permissions {
				_, err := s.dao.GetPermission(c, permission)
				if err != nil {
					return errors.New("permission not exist")
				}
			}
			err = s.dao.DeleteRolePermissionByRoleId(c, info.RoleId)
			if err != nil {
				return err
			}
			for _, permission := range info.Permissions {
				rInfo, err := s.dao.GetPermission(c, permission)
				if err != nil {
					return errors.New("permission not exist")
				}

				content += fmt.Sprintf("permission id:%d;permission:%s;", rInfo.Id, rInfo.Name)
				_, _ = s.dao.CreateRolePermission(c, &model.RolePermission{
					RoleId:       info.RoleId,
					PermissionId: rInfo.Id,
				})
			}
		}
	}

	cc := fmt.Sprintf("修改管理员:账户:%s;账户编号:%d;%s", aInfo.Name, aInfo.Id, content)
	result := model.SUCCESS
	return s.logAction(c, "update_admin", cc, result)

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
				_, err := s.dao.GetRole(c, role)
				if err != nil {
					return errors.New("角色不存在")
				}
			}
			err = s.dao.DeleteAdminRoleByAdminId(c, info.AdminId)
			if err != nil {
				return err
			}
			for _, role := range info.Roles {
				rInfo, err := s.dao.GetRole(c, role)
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
	result := model.SUCCESS
	return s.logAction(c, "update_admin", cc, result)

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
	result := model.SUCCESS
	return s.logAction(c, "delete_admin", cc, result)
}

func (s *Service) AddAdminV2(c *gin.Context, info *model.AddAdmin) error {
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
	result := model.SUCCESS
	return s.logAction(c, "add_admin", cc, result)
}

// 角色

func (s *Service) FindRolePageV2(c *gin.Context, req *gcurd.Request) (reply *model.PageReply[*model.RoleInfo], err error) {
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
		ps := make([]*model.Permission, 0)

		rolePermissions, _ := s.dao.FindRolePermission(c, []*gcurd.WhereValue{gcurd.EQ("role_id", d.Id)})
		for _, rolePermission := range rolePermissions {
			menus, _ := s.dao.FindPermission(c, []*gcurd.WhereValue{gcurd.EQ("id", rolePermission.PermissionId)})
			ps = append(ps, menus...)
		}

		data = append(data, &model.RoleInfo{
			Id:          d.Id,
			Name:        d.Name,
			Permissions: ps,
		})
	}
	reply = &model.PageReply[*model.RoleInfo]{}

	reply.Rows = data
	reply.RowsTotal = count
	reply.Page = req.Pagination.Page
	reply.PageSize = req.Pagination.PageSize
	return
}

func (s *Service) GetRoleInfoV1(c *gin.Context, roleId int) (*model.RoleInfo, error) {
	var data *model.RoleInfo
	dataTmp, err := s.dao.GetRole(c, roleId)
	if err == sql.ErrNoRows {
		return nil, errors.New("角色不存在")
	}

	if err != nil {
		return nil, err
	}

	ps := make([]*model.Permission, 0)

	rolePermissions, _ := s.dao.FindRolePermission(c, []*gcurd.WhereValue{gcurd.EQ("role_id", dataTmp.Id)})
	for _, rolePermission := range rolePermissions {
		menus, _ := s.dao.FindPermission(c, []*gcurd.WhereValue{gcurd.EQ("id", rolePermission.PermissionId)})
		ps = append(ps, menus...)
	}

	data = &model.RoleInfo{
		Id:          dataTmp.Id,
		Name:        dataTmp.Name,
		Remark:      dataTmp.Remark,
		Permissions: ps,
	}

	return data, err
}

func (s *Service) FindAllPermission(c *gin.Context) (reply []*model.Permission, err error) {
	return s.dao.FindPermission(c, nil)
}

func (s *Service) FindAllMenu(c *gin.Context) (reply []*model.Menu, err error) {
	return s.dao.FindMenu(c, nil)
}

func (s *Service) FindAllApi(c *gin.Context) (reply []*model.Api, err error) {
	return s.dao.FindApi(c, nil)
}

func (s *Service) AddRoleV2(c *gin.Context, role *model.AddRole) error {

	aInfo, err := s.dao.GetRoleByName(c, role.Name)
	if err != sql.ErrNoRows {
		return errors.New("name exist")
	}
	res, err := s.dao.CreateRole(c, &model.Role{
		Name:   role.Name,
		Remark: role.Remark,
	})
	if err != nil {
		return err
	}
	id, _ := res.LastInsertId()

	for _, permission := range role.Permissions {
		_, _ = s.dao.CreateRolePermission(c, &model.RolePermission{
			RoleId:       int(id),
			PermissionId: permission,
		})
	}

	content := fmt.Sprintf("添加角色:角色:%s;角色编号:%d;备注：%s", aInfo.Name, id, aInfo.Remark)
	result := model.SUCCESS
	return s.logAction(c, "add_role", content, result)
}

func (s *Service) DeleteRoleV1(c *gin.Context, id int) error {

	var content string
	roleInfo, err := s.dao.GetRole(c, id)
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
	err = s.dao.DeleteRolePermission(c, id)
	result := model.SUCCESS
	return s.logAction(c, "delete_role", fmt.Sprintf("删除角色:角色:%s;角色编号:%d;%s", roleInfo.Name, roleInfo.Id, content), result)
}

func (s *Service) logAction(c *gin.Context, opCode string, content string, result string) error {
	return logAction(c, s.dao.DB(), opCode, content, result)
}

func (s *Service) GetAdminByName(c *gin.Context, name string) (*model.Admin, error) {
	return s.dao.GetAdminByName(c, name)
}

func (s *Service) GetPermissionInfoV1(c *gin.Context, permissionId int) (*model.PermissionInfo, error) {
	var data *model.PermissionInfo
	dataTmp, err := s.dao.GetPermission(c, permissionId)
	if err == sql.ErrNoRows {
		return nil, errors.New("permission not exist")
	}

	if err != nil {
		return nil, err
	}

	ms := make([]*model.Menu, 0)
	permissionMenus, _ := s.dao.FindPermissionMenu(c, []*gcurd.WhereValue{gcurd.EQ("permission_id", dataTmp.Id)})
	for _, permissionMenu := range permissionMenus {
		menus, _ := s.dao.FindMenu(c, []*gcurd.WhereValue{gcurd.EQ("id", permissionMenu.MenuId)})
		ms = append(ms, menus...)
	}
	as := make([]*model.Api, 0)
	permissionApis, _ := s.dao.FindPermissionApi(c, []*gcurd.WhereValue{gcurd.EQ("permission_id", dataTmp.Id)})
	for _, permissionApi := range permissionApis {
		apis, _ := s.dao.FindApi(c, []*gcurd.WhereValue{gcurd.EQ("id", permissionApi.ApiId)})
		as = append(as, apis...)
	}

	data = &model.PermissionInfo{
		Id:     dataTmp.Id,
		Name:   dataTmp.Name,
		Remark: dataTmp.Remark,
		Menus:  ms,
		Apis:   as,
	}

	return data, err
}

func (s *Service) UpdatePermissionV1(c *gin.Context, info *model.UpdatePermission, filed []string) error {
	var content string
	pInfo, err := s.dao.GetPermission(c, info.Id)
	if err != nil {
		return errors.New("permission not exit")
	}
	for _, v := range filed {
		switch v {
		case "name":
			content += fmt.Sprintf("name:%s;", info.Name)
			_ = s.dao.UpdatePermission(c, info.Id, "name", info.Name)
		case "permission_group":
			content += fmt.Sprintf("permission_group:%s;", info.PermissionGroup)
			_ = s.dao.UpdatePermission(c, info.Id, "permission_group", info.PermissionGroup)
		case "remark":
			content += fmt.Sprintf("remark:%s;", info.Remark)
			_ = s.dao.UpdatePermission(c, info.Id, "remark", info.Remark)
		case "menus":
			for _, menu := range info.Menus {
				_, err := s.dao.GetMenu(c, menu)
				if err != nil {
					return errors.New("menu not exist")
				}
			}
			err = s.dao.DeletePermissionMenuByPermissionId(c, info.Id)
			if err != nil {
				return err
			}
			for _, menu := range info.Menus {
				mInfo, err := s.dao.GetMenu(c, menu)
				if err != nil {
					return errors.New("menu not exist")
				}

				content += fmt.Sprintf("menu id:%d;menu:%s;", mInfo.Id, mInfo.Name)
				_, _ = s.dao.CreatePermissionMenu(c, &model.PermissionMenu{
					PermissionId: pInfo.Id,
					MenuId:       mInfo.Id,
				})
			}
		case "apis":
			for _, api := range info.Apis {
				_, err := s.dao.GetApi(c, api)
				if err != nil {
					return errors.New("api not exist")
				}
			}
			err = s.dao.DeletePermissionApiByPermissionId(c, info.Id)
			if err != nil {
				return err
			}
			for _, api := range info.Apis {
				aInfo, err := s.dao.GetApi(c, api)
				if err != nil {
					return errors.New("api not exist")
				}

				content += fmt.Sprintf("api id:%d;api:%s;", aInfo.Id, aInfo.Name)
				_, _ = s.dao.CreatePermissionApi(c, &model.PermissionApi{
					PermissionId: pInfo.Id,
					ApiId:        aInfo.Id,
				})
			}
		}
	}

	cc := fmt.Sprintf("set permission:permission:%s;permission id:%d;%s", pInfo.Name, pInfo.Id, content)
	result := model.SUCCESS
	return s.logAction(c, "update_permission", cc, result)

}

func (s *Service) AddPermissionV1(c *gin.Context, permission *model.AddPermission) error {

	aInfo, err := s.dao.GetPermissionByName(c, permission.Name)
	if err != sql.ErrNoRows {
		return errors.New("name exist")
	}
	res, err := s.dao.CreatePermission(c, &model.Permission{
		Name:            permission.Name,
		PermissionGroup: permission.PermissionGroup,
		Remark:          permission.Remark,
	})
	if err != nil {
		return err
	}
	id, _ := res.LastInsertId()

	for _, menuId := range permission.Menus {
		_, _ = s.dao.CreatePermissionMenu(c, &model.PermissionMenu{
			PermissionId: int(id),
			MenuId:       menuId,
		})
	}

	for _, apiId := range permission.Apis {
		_, _ = s.dao.CreatePermissionApi(c, &model.PermissionApi{
			PermissionId: int(id),
			ApiId:        apiId,
		})
	}

	content := fmt.Sprintf("add permission:permission:%s;permission id:%d;permission group:%s;remark：%s", aInfo.Name, id, aInfo.PermissionGroup, aInfo.Remark)
	result := model.SUCCESS
	return s.logAction(c, "add_permission", content, result)
}

func (s *Service) DeletePermissionV1(c *gin.Context, permissionId int) error {

	var content string
	permission, err := s.dao.GetPermission(c, permissionId)
	if err != nil {
		return errors.New("permission not exist")
	}

	err = s.dao.DeletePermission(c, permissionId)
	err = s.dao.DeletePermissionMenuByPermissionId(c, permissionId)
	err = s.dao.DeletePermissionApiByPermissionId(c, permissionId)
	result := model.SUCCESS
	return s.logAction(c, "delete_permission", fmt.Sprintf("delete permission:permission:%s;permission id:%d;%s", permission.Name, permission.Id, content), result)
}
