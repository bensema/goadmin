package service

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/bensema/goadmin/model"
	"github.com/bensema/goadmin/utils"
	"github.com/bensema/library/ecode"
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
	_ = s.dao.CreateLogAdminLogin(c, loginLog)

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

func (s *Service) FindAdminPage(c *gin.Context, req *model.FindAdminReq) (reply *model.FindAdminReply, err error) {
	reply = &model.FindAdminReply{}
	var count int
	var dataTmp []*model.Admin
	if count, err = s.dao.PageFindAdminTotal(c, req); err != nil {
		return
	}
	if count <= 0 {
		return
	}
	if dataTmp, err = s.dao.FindAdmin(c, req); err != nil {
		return
	}
	for _, d := range dataTmp {
		d.Password = "***"
	}
	reply.Data = dataTmp
	reply.Total = count
	reply.Num = req.Num
	reply.Size = req.Size
	return
}

func (s *Service) FindAdminPageV1(c *gin.Context, req *model.FindAdminReq) (reply *model.FindAdminReplyV1, err error) {
	reply = &model.FindAdminReplyV1{}
	var count int
	var dataTmp []*model.Admin
	var data []*model.AdminV1
	if count, err = s.dao.PageFindAdminTotal(c, req); err != nil {
		return
	}
	if count <= 0 {
		return
	}
	if dataTmp, err = s.dao.FindAdmin(c, req); err != nil {
		return
	}
	for _, d := range dataTmp {
		var rls []*model.Role
		farr := &model.FindAdminRoleReq{}
		farr.AdminId = d.Id
		adminRoles, _ := s.dao.FindAdminRole(c, farr)

		for _, adminRole := range adminRoles {
			frr := &model.FindRoleReq{}
			frr.Id = adminRole.RoleId
			roles, _ := s.dao.FindRole(c, frr)
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
	reply.Data = data
	reply.Total = count
	reply.Num = req.Num
	reply.Size = req.Size
	return
}

// 获取管理员信息过滤密码，添加角色
func (s *Service) GetAdminV1(c *gin.Context, id int) (*model.AdminV1, error) {
	var data *model.AdminV1
	dataTmp, err := s.dao.GetAdminById(c, id)
	if err == sql.ErrNoRows {
		return nil, errors.New("账户不存在")
	}

	if err != nil {
		return nil, err
	}

	var rls []*model.Role
	farr := &model.FindAdminRoleReq{}
	farr.AdminId = id
	adminRoles, _ := s.dao.FindAdminRole(c, farr)

	for _, adminRole := range adminRoles {
		frr := &model.FindRoleReq{}
		frr.Id = adminRole.RoleId
		roles, _ := s.dao.FindRole(c, frr)
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
		Roles:     rls,
	}

	return data, err
}

func (s *Service) FindAllRole(c *gin.Context) (reply []*model.Role, err error) {
	return s.dao.FindRole(c, &model.FindRoleReq{})
}

func (s *Service) UpdateAdmin(c *gin.Context, info *model.UpdateAdmin, filed []string) error {
	var content string
	aInfo, err := s.dao.GetAdminById(c, info.Id)
	if err != nil {
		return err
	}
	for _, v := range filed {
		switch v {
		case "status":
			content += fmt.Sprintf("状态:%d;", info.Status)
			_ = s.dao.UpdateAdminById(c, info.Id, "status", info.Status)
		case "roles":
			for _, role := range info.Roles {
				_, err := s.dao.GetRoleById(c, role)
				if err != nil {
					return errors.New("角色不存在")
				}
			}
			err = s.dao.DeleteAdminRoleByAdminId(c, info.Id)
			if err != nil {
				return err
			}
			for _, role := range info.Roles {
				rInfo, err := s.dao.GetRoleById(c, role)
				if err != nil {
					return errors.New("角色不存在")
				}

				content += fmt.Sprintf("角色编号:%d;角色:%s;", rInfo.Id, rInfo.Name)
				_ = s.dao.CreateAdminRole(c, &model.AdminRole{
					AdminId: info.Id,
					RoleId:  rInfo.Id,
				})
			}
		}
	}

	operatorInfo, err := s.getAdminFromContext(c)
	if err != nil {
		return err
	}
	recordLog := &model.LogAdminOperation{
		AdminId:       operatorInfo.Id,
		Name:          operatorInfo.Name,
		OperationCode: "",
		OperationName: "",
		Content:       fmt.Sprintf("修改管理员:账户:%s;账户编号:%d;%s", aInfo.Name, aInfo.Id, content),
		Result:        1,
		Ip:            c.ClientIP(),
		RecordAt:      xtime.Time(time.Now().Unix()),
	}
	err = s.dao.CreateLogAdminOperation(c, recordLog)
	return err
}

func (s *Service) DeleteAdmin(c *gin.Context, id int) error {
	if id == 0 || id == 1 {
		return errors.New("权限不足")
	}

	var content string
	aInfo, err := s.dao.GetAdminById(c, id)
	if err != nil {
		return errors.New("账户不存在")
	}

	err = s.dao.DeleteAdminById(c, id)
	err = s.dao.DeleteAdminRoleByAdminId(c, id)

	operatorInfo, err := s.getAdminFromContext(c)
	if err != nil {
		return err
	}
	recordLog := &model.LogAdminOperation{
		AdminId:       operatorInfo.Id,
		Name:          operatorInfo.Name,
		OperationCode: "delete_admin",
		OperationName: "删除账户",
		Content:       fmt.Sprintf("删除管理员:账户:%s;账户编号:%d;%s", aInfo.Name, aInfo.Id, content),
		Result:        1,
		Ip:            c.ClientIP(),
		RecordAt:      xtime.Time(time.Now().Unix()),
	}
	err = s.dao.CreateLogAdminOperation(c, recordLog)
	return err
}

func (s *Service) AddAdmin(c *gin.Context, info *model.AddAdmin) error {
	if err := utils.CheckNameLegal(info.Name); err != nil {
		return err
	}
	if err := utils.CheckPasswordLegal(info.Password); err != nil {
		return err
	}

	user := &model.Admin{}
	aInfo, err := s.dao.GetAdminByName(c, info.Name)
	if err != sql.ErrNoRows {
		return errors.New("账户已存在")
	}

	hashPwd, err := utils.HashAndSalt(info.Password)
	if err != nil {
		return ecode.PasswordEncodeErr
	}

	user.Name = info.Name
	user.Password = hashPwd
	user.Status = info.Status
	user.CreatedAt = xtime.Time(time.Now().Unix())
	user.UpdatedAt = xtime.Time(time.Now().Unix())
	err = s.dao.CreateAdmin(c, user)
	if err != nil {
		return err
	}
	uInfo, err := s.dao.GetAdminByName(c, info.Name)
	if err != nil {
		return err
	}
	for _, role := range info.Roles {
		_ = s.dao.CreateAdminRole(c, &model.AdminRole{
			AdminId: uInfo.Id,
			RoleId:  role,
		})
	}

	operatorInfo, err := s.getAdminFromContext(c)
	if err != nil {
		return err
	}
	recordLog := &model.LogAdminOperation{
		AdminId:       operatorInfo.Id,
		Name:          operatorInfo.Name,
		OperationCode: "add_admin",
		OperationName: "添加账户",
		Content:       fmt.Sprintf("添加管理员:账户:%s;账户编号:%d;", aInfo.Name, aInfo.Id),
		Result:        1,
		Ip:            c.ClientIP(),
		RecordAt:      xtime.Time(time.Now().Unix()),
	}
	err = s.dao.CreateLogAdminOperation(c, recordLog)
	return err
}

// 角色

func (s *Service) FindRolePageV1(c *gin.Context, req *model.FindRoleReq) (reply *model.FindRoleReplyV1, err error) {
	reply = &model.FindRoleReplyV1{}
	var count int
	var dataTmp []*model.Role
	var data []*model.RoleV1
	if count, err = s.dao.PageFindRoleTotal(c, req); err != nil {
		return
	}
	if count <= 0 {
		return
	}
	if dataTmp, err = s.dao.FindRole(c, req); err != nil {
		return
	}
	for _, d := range dataTmp {
		var ps []*model.Permission
		frpr := &model.FindRolePermissionReq{}
		frpr.RoleId = d.Id
		rolePermissions, _ := s.dao.FindRolePermission(c, frpr)

		for _, rolePermission := range rolePermissions {
			fpr := &model.FindPermissionReq{}
			fpr.Id = rolePermission.PermissionId
			permissions, _ := s.dao.FindPermission(c, fpr)
			for _, permission := range permissions {
				ps = append(ps, permission)
			}
		}

		data = append(data, &model.RoleV1{
			Id:          d.Id,
			Name:        d.Name,
			Permissions: ps,
		})
	}
	reply.Data = data
	reply.Total = count
	reply.Num = req.Num
	reply.Size = req.Size
	return
}

func (s *Service) GetRoleV1(c *gin.Context, id int) (*model.RoleV1, error) {
	var data *model.RoleV1
	dataTmp, err := s.dao.GetRoleById(c, id)
	if err == sql.ErrNoRows {
		return nil, errors.New("角色不存在")
	}

	if err != nil {
		return nil, err
	}

	var rls []*model.Permission
	farr := &model.FindRolePermissionReq{}
	farr.RoleId = id
	adminRoles, _ := s.dao.FindRolePermission(c, farr)

	for _, adminRole := range adminRoles {
		frr := &model.FindPermissionReq{}
		frr.Id = adminRole.PermissionId
		roles, _ := s.dao.FindPermission(c, frr)
		for _, role := range roles {
			rls = append(rls, role)
		}
	}

	data = &model.RoleV1{
		Id:          dataTmp.Id,
		Name:        dataTmp.Name,
		Permissions: rls,
	}

	return data, err
}

func (s *Service) FindAllPermission(c *gin.Context) (reply []*model.Permission, err error) {
	return s.dao.FindPermission(c, &model.FindPermissionReq{})
}

func (s *Service) FindAllMenu(c *gin.Context) (reply []*model.Menu, err error) {
	return s.dao.FindMenu(c, &model.FindMenuReq{})
}

func (s *Service) FindAllOperation(c *gin.Context) (reply []*model.Operation, err error) {
	return s.dao.FindOperation(c, &model.FindOperationReq{})
}

func (s *Service) UpdateRole(c *gin.Context, info *model.UpdateRole, filed []string) error {
	var content string
	aInfo, err := s.dao.GetRoleById(c, info.Id)
	if err != nil {
		return err
	}
	for _, v := range filed {
		switch v {
		case "name":
			content += fmt.Sprintf("名称:%s;", info.Name)
			_ = s.dao.UpdateRoleById(c, info.Id, "name", info.Name)
		case "permissions":
			for _, permission := range info.Permissions {
				_, err := s.dao.GetPermissionById(c, permission)
				if err != nil {
					return errors.New("权限不存在")
				}
			}
			err = s.dao.DeleteRolePermissionByRoleId(c, info.Id)
			if err != nil {
				return err
			}
			for _, permission := range info.Permissions {
				rInfo, err := s.dao.GetPermissionById(c, permission)
				if err != nil {
					return errors.New("权限不存在")
				}

				content += fmt.Sprintf("权限编号:%d;权限:%s;", rInfo.Id, rInfo.Name)
				_ = s.dao.CreateRolePermission(c, &model.RolePermission{
					RoleId:       info.Id,
					PermissionId: permission,
				})
			}
		}
	}

	operatorInfo, err := s.getAdminFromContext(c)
	if err != nil {
		return err
	}
	recordLog := &model.LogAdminOperation{
		AdminId:       operatorInfo.Id,
		Name:          operatorInfo.Name,
		OperationCode: "role_update",
		OperationName: "修改角色",
		Content:       fmt.Sprintf("修改角色:角色:%s;角色编号:%d;%s", aInfo.Name, aInfo.Id, content),
		Result:        1,
		Ip:            c.ClientIP(),
		RecordAt:      xtime.Time(time.Now().Unix()),
	}
	err = s.dao.CreateLogAdminOperation(c, recordLog)
	return err
}

func (s *Service) AddRole(c *gin.Context, info *model.AddRole) error {

	role := &model.Role{}
	aInfo, err := s.dao.GetRoleByName(c, info.Name)
	if err != sql.ErrNoRows {
		return errors.New("角色名已存在")
	}
	role.Name = info.Name

	err = s.dao.CreateRole(c, role)
	if err != nil {
		return err
	}
	uInfo, err := s.dao.GetRoleByName(c, info.Name)
	if err != nil {
		return err
	}
	for _, permission := range info.Permissions {
		_ = s.dao.CreateRolePermission(c, &model.RolePermission{
			RoleId:       uInfo.Id,
			PermissionId: permission,
		})
	}

	operatorInfo, err := s.getAdminFromContext(c)
	if err != nil {
		return err
	}
	recordLog := &model.LogAdminOperation{
		AdminId:       operatorInfo.Id,
		Name:          operatorInfo.Name,
		OperationCode: "add_role",
		OperationName: "添加角色",
		Content:       fmt.Sprintf("添加角色:角色:%s;角色编号:%d;", aInfo.Name, aInfo.Id),
		Result:        1,
		Ip:            c.ClientIP(),
		RecordAt:      xtime.Time(time.Now().Unix()),
	}
	err = s.dao.CreateLogAdminOperation(c, recordLog)
	return err
}

func (s *Service) DeleteRole(c *gin.Context, id int) error {

	var content string
	aInfo, err := s.dao.GetRoleById(c, id)
	if err != nil {
		return errors.New("角色不存在")
	}

	u := &model.FindAdminRoleReq{}
	u.RoleId = id
	arl, err := s.dao.FindAdminRole(c, u)
	if err != nil {
		return errors.New("获取用户-角色失败")
	}
	if len(arl) > 0 {
		return errors.New("还有账户在使用此角色")
	}

	err = s.dao.DeleteRoleById(c, id)
	err = s.dao.DeleteRolePermissionByRoleId(c, id)

	operatorInfo, err := s.getAdminFromContext(c)
	if err != nil {
		return err
	}
	recordLog := &model.LogAdminOperation{
		AdminId:       operatorInfo.Id,
		Name:          operatorInfo.Name,
		OperationCode: "delete_role",
		OperationName: "删除角色",
		Content:       fmt.Sprintf("删除角色:角色:%s;角色编号:%d;%s", aInfo.Name, aInfo.Id, content),
		Result:        1,
		Ip:            c.ClientIP(),
		RecordAt:      xtime.Time(time.Now().Unix()),
	}
	err = s.dao.CreateLogAdminOperation(c, recordLog)
	return err
}

// 权限

func (s *Service) FindPermissionPage(c *gin.Context, req *model.FindPermissionReq) (reply *model.FindPermissionReply, err error) {
	reply = &model.FindPermissionReply{}
	var count int
	var data []*model.Permission
	if count, err = s.dao.PageFindPermissionTotal(c, req); err != nil {
		return
	}
	if count <= 0 {
		return
	}
	if data, err = s.dao.FindPermission(c, req); err != nil {
		return
	}
	reply.Data = data
	reply.Total = count
	reply.Num = req.Num
	reply.Size = req.Size
	return
}

func (s *Service) AddPermission(c *gin.Context, info *model.Permission) error {

	_, err := s.dao.GetPermissionByName(c, info.Name)
	if err != sql.ErrNoRows {
		return errors.New("权限已存在")
	}

	err = s.dao.CreatePermission(c, &model.Permission{Name: info.Name})
	if err != nil {
		return err
	}
	uInfo, err := s.dao.GetPermissionByName(c, info.Name)
	operatorInfo, err := s.getAdminFromContext(c)
	if err != nil {
		return err
	}
	recordLog := &model.LogAdminOperation{
		AdminId:       operatorInfo.Id,
		Name:          operatorInfo.Name,
		OperationCode: "add_permission",
		OperationName: "添加权限",
		Content:       fmt.Sprintf("添加权限:权限:%s;权限编号:%d;", uInfo.Name, uInfo.Id),
		Result:        1,
		Ip:            c.ClientIP(),
		RecordAt:      xtime.Time(time.Now().Unix()),
	}
	err = s.dao.CreateLogAdminOperation(c, recordLog)
	return err
}

func (s *Service) DeletePermission(c *gin.Context, id int) error {

	var content string
	aInfo, err := s.dao.GetPermissionById(c, id)
	if err != nil {
		return errors.New("权限不存在")
	}

	u := &model.FindRolePermissionReq{}
	u.PermissionId = id
	arl, err := s.dao.FindRolePermission(c, u)
	if err != nil {
		return errors.New("获取角色-权限失败")
	}
	if len(arl) > 0 {
		return errors.New("还有角色在使用此权限")
	}

	err = s.dao.DeletePermissionById(c, id)
	err = s.dao.DeletePermissionMenuByPermissionId(c, id)
	err = s.dao.DeletePermissionOperationByPermissionId(c, id)

	operatorInfo, err := s.getAdminFromContext(c)
	if err != nil {
		return err
	}
	recordLog := &model.LogAdminOperation{
		AdminId:       operatorInfo.Id,
		Name:          operatorInfo.Name,
		OperationCode: "delete_permission",
		OperationName: "删除权限",
		Content:       fmt.Sprintf("删除权限:权限:%s;权限编号:%d;%s", aInfo.Name, aInfo.Id, content),
		Result:        1,
		Ip:            c.ClientIP(),
		RecordAt:      xtime.Time(time.Now().Unix()),
	}
	err = s.dao.CreateLogAdminOperation(c, recordLog)
	return err
}

func (s *Service) UpdatePermission(c *gin.Context, info *model.UpdatePermission, filed []string) error {
	var content string
	aInfo, err := s.dao.GetPermissionById(c, info.Id)
	if err == sql.ErrNoRows {
		return errors.New("权限对象不存在")
	}
	if err != nil {
		return err
	}
	for _, v := range filed {
		switch v {
		case "name":
			content += fmt.Sprintf("名称:%s;", info.Name)
			_ = s.dao.UpdatePermissionById(c, info.Id, "name", info.Name)
		case "menus":
			for _, menu := range info.Menus {
				_, err := s.dao.GetMenuById(c, menu)
				if err != nil {
					return errors.New("菜单不存在")
				}
			}
			err = s.dao.DeletePermissionMenuByPermissionId(c, info.Id)
			if err != nil {
				return err
			}
			for _, menu := range info.Menus {
				rInfo, err := s.dao.GetMenuById(c, menu)
				if err != nil {
					return errors.New("菜单不存在")
				}

				content += fmt.Sprintf("菜单编号:%d;菜单:%s;", rInfo.Id, rInfo.Name)
				_ = s.dao.CreatePermissionMenu(c, &model.PermissionMenu{
					PermissionId: info.Id,
					MenuId:       menu,
				})
			}
		case "operations":
			for _, operation := range info.Operation {
				_, err := s.dao.GetOperationById(c, operation)
				if err != nil {
					return errors.New("操作API不存在")
				}
			}
			err = s.dao.DeletePermissionOperationByPermissionId(c, info.Id)
			if err != nil {
				return err
			}
			for _, operation := range info.Operation {
				rInfo, err := s.dao.GetOperationById(c, operation)
				if err != nil {
					return errors.New("操作不存在")
				}

				content += fmt.Sprintf("操作编号:%d;操作:%s;", rInfo.Id, rInfo.Name)
				_ = s.dao.CreatePermissionOperation(c, &model.PermissionOperation{
					PermissionId: info.Id,
					OperationId:  operation,
				})
			}
		}

	}

	operatorInfo, err := s.getAdminFromContext(c)
	if err != nil {
		return err
	}
	recordLog := &model.LogAdminOperation{
		AdminId:       operatorInfo.Id,
		Name:          operatorInfo.Name,
		OperationCode: "permission_update",
		OperationName: "修改权限",
		Content:       fmt.Sprintf("修改权限:权限:%s;权限编号:%d;%s", aInfo.Name, aInfo.Id, content),
		Result:        1,
		Ip:            c.ClientIP(),
		RecordAt:      xtime.Time(time.Now().Unix()),
	}
	err = s.dao.CreateLogAdminOperation(c, recordLog)
	return err
}

func (s *Service) FindPermissionMenu(c *gin.Context, req *model.FindPermissionMenuReq) (reply []*model.PermissionMenu, err error) {
	return s.dao.FindPermissionMenu(c, req)
}

func (s *Service) FindPermissionOperation(c *gin.Context, req *model.FindPermissionOperationReq) (reply []*model.PermissionOperation, err error) {
	return s.dao.FindPermissionOperation(c, req)
}

func (s *Service) AddMenu(c *gin.Context, info *model.Menu) error {

	_, err := s.dao.GetMenuByName(c, info.Name)
	if err != sql.ErrNoRows {
		return errors.New("菜单已存在")
	}

	err = s.dao.CreateMenu(c, info)
	if err != nil {
		return err
	}
	uInfo, err := s.dao.GetMenuByName(c, info.Name)
	operatorInfo, err := s.getAdminFromContext(c)
	if err != nil {
		return err
	}
	recordLog := &model.LogAdminOperation{
		AdminId:       operatorInfo.Id,
		Name:          operatorInfo.Name,
		OperationCode: "add_menu",
		OperationName: "添加菜单",
		Content:       fmt.Sprintf("添加菜单:权菜单限:%s;菜单编号:%d;", uInfo.Name, uInfo.Id),
		Result:        1,
		Ip:            c.ClientIP(),
		RecordAt:      xtime.Time(time.Now().Unix()),
	}
	err = s.dao.CreateLogAdminOperation(c, recordLog)
	return err
}

func (s *Service) AddOperation(c *gin.Context, info *model.Operation) error {

	_, err := s.dao.GetOperationByName(c, info.Name)
	if err != sql.ErrNoRows {
		return errors.New("操作功能已存在")
	}

	if info.Pid != 0 {
		_, err = s.dao.GetMenuById(c, info.Pid)
		if err != nil {
			return errors.New("上级菜单不存在")
		}
	}

	err = s.dao.CreateOperation(c, info)
	if err != nil {
		return err
	}
	uInfo, err := s.dao.GetOperationByName(c, info.Name)
	operatorInfo, err := s.getAdminFromContext(c)
	if err != nil {
		return err
	}
	recordLog := &model.LogAdminOperation{
		AdminId:       operatorInfo.Id,
		Name:          operatorInfo.Name,
		OperationCode: "add_operation",
		OperationName: "添加操作功能",
		Content:       fmt.Sprintf("添加操作功能:操作功能:%s;操作功能编号:%d;", uInfo.Name, uInfo.Id),
		Result:        1,
		Ip:            c.ClientIP(),
		RecordAt:      xtime.Time(time.Now().Unix()),
	}
	err = s.dao.CreateLogAdminOperation(c, recordLog)
	return err
}

func (s *Service) DeleteMenu(c *gin.Context, id int) error {

	var content string
	aInfo, err := s.dao.GetMenuById(c, id)
	if err != nil {
		return errors.New("菜单不存在")
	}

	u := &model.FindOperationReq{}
	u.Pid = id
	arl, err := s.dao.FindOperation(c, u)
	if err != nil {
		return errors.New("获取操作功能失败")
	}
	if len(arl) > 0 {
		return errors.New("菜单下还有API操作功能")
	}

	err = s.dao.DeleteMenuById(c, id)

	operatorInfo, err := s.getAdminFromContext(c)
	if err != nil {
		return err
	}
	recordLog := &model.LogAdminOperation{
		AdminId:       operatorInfo.Id,
		Name:          operatorInfo.Name,
		OperationCode: "delete_menu",
		OperationName: "删除菜单",
		Content:       fmt.Sprintf("删除菜单:菜单:%s;菜单编号:%d;%s", aInfo.Name, aInfo.Id, content),
		Result:        1,
		Ip:            c.ClientIP(),
		RecordAt:      xtime.Time(time.Now().Unix()),
	}
	err = s.dao.CreateLogAdminOperation(c, recordLog)
	return err
}

func (s *Service) UpdateMenu(c *gin.Context, info *model.Menu, filed []string) error {
	var content string
	aInfo, err := s.dao.GetMenuById(c, info.Id)
	if err != nil {
		return errors.New("id不存在")
	}
	for _, v := range filed {
		switch v {
		case "name":
			content += fmt.Sprintf("名称:%s;", info.Name)
			_ = s.dao.UpdateMenuById(c, info.Id, "name", info.Name)
		case "pid":
			if info.Id == info.Pid {
				return errors.New("上级不能是自己")
			}
			content += fmt.Sprintf("pid:%d;", info.Pid)
			_ = s.dao.UpdateMenuById(c, info.Id, "pid", info.Pid)
		case "icon":
			content += fmt.Sprintf("icon:%s;", info.Icon)
			_ = s.dao.UpdateMenuById(c, info.Id, "icon", info.Icon)
		case "url":
			content += fmt.Sprintf("url:%s;", info.Url)
			_ = s.dao.UpdateMenuById(c, info.Id, "url", info.Url)
		case "index_sort":
			content += fmt.Sprintf("index_sort:%d;", info.IndexSort)
			_ = s.dao.UpdateMenuById(c, info.Id, "index_sort", info.IndexSort)

		}
	}

	operatorInfo, err := s.getAdminFromContext(c)
	if err != nil {
		return err
	}
	recordLog := &model.LogAdminOperation{
		AdminId:       operatorInfo.Id,
		Name:          operatorInfo.Name,
		OperationCode: "update_menu",
		OperationName: "更新菜单",
		Content:       fmt.Sprintf("修改菜单:菜单:%s;菜单编号:%d;%s", aInfo.Name, aInfo.Id, content),
		Result:        1,
		Ip:            c.ClientIP(),
		RecordAt:      xtime.Time(time.Now().Unix()),
	}
	err = s.dao.CreateLogAdminOperation(c, recordLog)
	return err
}

func (s *Service) DeleteOperation(c *gin.Context, id int) error {

	var content string
	aInfo, err := s.dao.GetOperationById(c, id)
	if err != nil {
		return errors.New("操作功能不存在")
	}

	err = s.dao.DeleteOperationById(c, id)

	operatorInfo, err := s.getAdminFromContext(c)
	if err != nil {
		return err
	}
	recordLog := &model.LogAdminOperation{
		AdminId:       operatorInfo.Id,
		Name:          operatorInfo.Name,
		OperationCode: "delete_operation",
		OperationName: "删除操作功能",
		Content:       fmt.Sprintf("删除操作功能:操作功能:%s;操作功能编号:%d;%s", aInfo.Name, aInfo.Id, content),
		Result:        1,
		Ip:            c.ClientIP(),
		RecordAt:      xtime.Time(time.Now().Unix()),
	}
	err = s.dao.CreateLogAdminOperation(c, recordLog)
	return err
}

func (s *Service) UpdateOperation(c *gin.Context, info *model.Operation, filed []string) error {
	var content string
	aInfo, err := s.dao.GetOperationById(c, info.Id)
	if err != nil {
		return errors.New("id不存在")
	}
	for _, v := range filed {
		switch v {
		case "name":
			content += fmt.Sprintf("名称:%s;", info.Name)
			_ = s.dao.UpdateOperationById(c, info.Id, "name", info.Name)
		case "pid":
			content += fmt.Sprintf("pid:%d;", info.Pid)
			_ = s.dao.UpdateOperationById(c, info.Id, "pid", info.Pid)
		case "code":
			content += fmt.Sprintf("code:%s;", info.Code)
			_ = s.dao.UpdateOperationById(c, info.Id, "code", info.Code)
		case "url":
			content += fmt.Sprintf("url:%s;", info.Url)
			_ = s.dao.UpdateOperationById(c, info.Id, "url", info.Url)
		case "method":
			content += fmt.Sprintf("method:%s;", info.Method)
			_ = s.dao.UpdateOperationById(c, info.Id, "method", info.Method)

		}
	}

	operatorInfo, err := s.getAdminFromContext(c)
	if err != nil {
		return err
	}
	recordLog := &model.LogAdminOperation{
		AdminId:       operatorInfo.Id,
		Name:          operatorInfo.Name,
		OperationCode: "update_operation",
		OperationName: "更新操作功能",
		Content:       fmt.Sprintf("修改操作功能:操作功能:%s;操作功能编号:%d;%s", aInfo.Name, aInfo.Id, content),
		Result:        1,
		Ip:            c.ClientIP(),
		RecordAt:      xtime.Time(time.Now().Unix()),
	}
	err = s.dao.CreateLogAdminOperation(c, recordLog)
	return err
}

func (s *Service) FindAdminLoginPage(c *gin.Context, req *model.FindLogAdminLoginReq) (reply *model.FindLogAdminLoginReply, err error) {
	reply = &model.FindLogAdminLoginReply{}
	var count int
	var data []*model.LogAdminLogin
	if count, err = s.dao.PageFindLogAdminLoginTotal(c, req); err != nil {
		return
	}
	if count <= 0 {
		return
	}
	if data, err = s.dao.PageFindLogAdminLogin(c, req); err != nil {
		return
	}
	reply.Data = data
	reply.Total = count
	reply.Num = req.Num
	reply.Size = req.Size
	return
}

func (s *Service) PageFindLogAdminOperation(c *gin.Context, req *model.FindLogAdminOperationReq) (reply *model.FindLogAdminOperationReply, err error) {
	reply = &model.FindLogAdminOperationReply{}
	var count int
	var data []*model.LogAdminOperation
	if count, err = s.dao.PageFindLogAdminOperationTotal(c, req); err != nil {
		return
	}
	if count <= 0 {
		return
	}
	if data, err = s.dao.PageFindLogAdminOperation(c, req); err != nil {
		return
	}
	reply.Data = data
	reply.Total = count
	reply.Num = req.Num
	reply.Size = req.Size
	return
}
