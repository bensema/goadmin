package model

import (
	xtime "github.com/bensema/library/time"
)

type Admin struct {
	Id        int        `json:"id"`         // 自增编号
	Name      string     `json:"name"`       // 账户
	Password  string     `json:"password"`   // 密码
	Status    int        `json:"status"`     // 1:正常;2:禁用
	CreatedAt xtime.Time `json:"created_at"` // 创建时间
	UpdatedAt xtime.Time `json:"updated_at"` // 修改时间
}

func (m *Admin) Table() string {
	return "admin"
}

func (m *Admin) Columns() []string {
	return []string{"id", "name", "password", "status", "created_at", "updated_at"}
}

func (m *Admin) Fields() []interface{} {
	return []interface{}{&m.Id, &m.Name, &m.Password, &m.Status, &m.CreatedAt, &m.UpdatedAt}
}

type AdminRole struct {
	Id      int `json:"id"`       // 编号
	AdminId int `json:"admin_id"` // 账户编号
	RoleId  int `json:"role_id"`  // 角色编号
}

func (m *AdminRole) Table() string {
	return "admin_role"
}

func (m *AdminRole) Columns() []string {
	return []string{"id", "admin_id", "role_id"}
}

func (m *AdminRole) Fields() []interface{} {
	return []interface{}{&m.Id, &m.AdminId, &m.RoleId}
}

type LogAdminLogin struct {
	Id         int        `json:"id"`          // 自增编号
	AdminId    int        `json:"admin_id"`    // 管理员编号
	Name       string     `json:"name"`        // 管理员名
	Location   string     `json:"location"`    // 位置
	Os         string     `json:"os"`          // 操作系统
	Browser    string     `json:"browser"`     // 浏览器
	UserAgent  string     `json:"user_agent"`  // 浏览器详情
	Url        string     `json:"url"`         // url
	Result     int        `json:"result"`      // 2:失败;1:成功
	Ip         string     `json:"ip"`          // IP
	RecordTime xtime.Time `json:"record_time"` // 记录时间
	Remark     string     `json:"remark"`      // 备注
}

func (m *LogAdminLogin) Table() string {
	return "log_admin_login"
}

func (m *LogAdminLogin) Columns() []string {
	return []string{"id", "admin_id", "name", "location", "os", "browser", "user_agent", "url", "result", "ip", "record_time", "remark"}
}

func (m *LogAdminLogin) Fields() []interface{} {
	return []interface{}{&m.Id, &m.AdminId, &m.Name, &m.Location, &m.Os, &m.Browser, &m.UserAgent, &m.Url, &m.Result, &m.Ip, &m.RecordTime, &m.Remark}
}

type LogAdminOperation struct {
	Id            int        `json:"id"`             // 操作编号
	AdminId       int        `json:"admin_id"`       // 管理员编号
	Name          string     `json:"name"`           // 账户
	OperationCode string     `json:"operation_code"` // 行为编号
	OperationName string     `json:"operation_name"` // 行为
	Content       string     `json:"content"`        // 操作内容
	Result        int        `json:"result"`         // 操作结果1:成功；2:失败
	Message       string     `json:"message"`        // 操作消息
	Ip            string     `json:"ip"`             // 操作IP
	RecordAt      xtime.Time `json:"record_at"`      // 操作时间
}

func (m *LogAdminOperation) Table() string {
	return "log_admin_operation"
}

func (m *LogAdminOperation) Columns() []string {
	return []string{"id", "admin_id", "name", "operation_code", "operation_name", "content", "result", "message", "ip", "record_at"}
}

func (m *LogAdminOperation) Fields() []interface{} {
	return []interface{}{&m.Id, &m.AdminId, &m.Name, &m.OperationCode, &m.OperationName, &m.Content, &m.Result, &m.Message, &m.Ip, &m.RecordAt}
}

type Menu struct {
	Id        int    `json:"id"`         // 编号
	Name      string `json:"name"`       // 菜单名
	Pid       int    `json:"pid"`        // 上级菜单
	Icon      string `json:"icon"`       // icon
	Url       string `json:"url"`        // url
	IndexSort int    `json:"index_sort"` // 排序
}

func (m *Menu) Table() string {
	return "menu"
}

func (m *Menu) Columns() []string {
	return []string{"id", "name", "pid", "icon", "url", "index_sort"}
}

func (m *Menu) Fields() []interface{} {
	return []interface{}{&m.Id, &m.Name, &m.Pid, &m.Icon, &m.Url, &m.IndexSort}
}

type Operation struct {
	Id     int    `json:"id"`     // 编号
	Name   string `json:"name"`   // 名称
	Code   string `json:"code"`   // 编码
	Method string `json:"method"` // 方法
	Url    string `json:"url"`    // url
	Pid    int    `json:"pid"`    // 上级编号
}

func (m *Operation) Table() string {
	return "operation"
}

func (m *Operation) Columns() []string {
	return []string{"id", "name", "code", "method", "url", "pid"}
}

func (m *Operation) Fields() []interface{} {
	return []interface{}{&m.Id, &m.Name, &m.Code, &m.Method, &m.Url, &m.Pid}
}

type Permission struct {
	Id   int    `json:"id"`   // 编号
	Name string `json:"name"` // 名称
}

func (m *Permission) Table() string {
	return "permission"
}

func (m *Permission) Columns() []string {
	return []string{"id", "name"}
}

func (m *Permission) Fields() []interface{} {
	return []interface{}{&m.Id, &m.Name}
}

type PermissionMenu struct {
	Id           int `json:"id"`            // 编号
	PermissionId int `json:"permission_id"` // 权限编号
	MenuId       int `json:"menu_id"`       // 菜单编号
}

func (m *PermissionMenu) Table() string {
	return "permission_menu"
}

func (m *PermissionMenu) Columns() []string {
	return []string{"id", "permission_id", "menu_id"}
}

func (m *PermissionMenu) Fields() []interface{} {
	return []interface{}{&m.Id, &m.PermissionId, &m.MenuId}
}

type PermissionOperation struct {
	Id           int `json:"id"`            // 编号
	PermissionId int `json:"permission_id"` // 权限编号
	OperationId  int `json:"operation_id"`  // 操作编号
}

func (m *PermissionOperation) Table() string {
	return "permission_operation"
}

func (m *PermissionOperation) Columns() []string {
	return []string{"id", "permission_id", "operation_id"}
}

func (m *PermissionOperation) Fields() []interface{} {
	return []interface{}{&m.Id, &m.PermissionId, &m.OperationId}
}

type Role struct {
	Id   int    `json:"id"`   // 编号
	Name string `json:"name"` // 角色名称
}

func (m *Role) Table() string {
	return "role"
}

func (m *Role) Columns() []string {
	return []string{"id", "name"}
}

func (m *Role) Fields() []interface{} {
	return []interface{}{&m.Id, &m.Name}
}

type RolePermission struct {
	Id           int `json:"id"`            // 编号
	RoleId       int `json:"role_id"`       // 角色编号
	PermissionId int `json:"permission_id"` // 权限编号
}

func (m *RolePermission) Table() string {
	return "role_permission"
}

func (m *RolePermission) Columns() []string {
	return []string{"id", "role_id", "permission_id"}
}

func (m *RolePermission) Fields() []interface{} {
	return []interface{}{&m.Id, &m.RoleId, &m.PermissionId}
}
