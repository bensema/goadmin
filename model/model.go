package model

import (
	"library/xtime"
)

type Admin struct {
	Id        int        `json:"id"`         // 自增编号
	Name      string     `json:"name"`       // 管理员
	Password  string     `json:"password"`   // 密码
	Status    string     `json:"status"`     // Enable/Disable
	CreatedAt xtime.Time `json:"created_at"` // 创建时间
	UpdatedAt xtime.Time `json:"updated_at"` // 修改时间
	Remark    string     `json:"remark"`     // remark
}

func (m *Admin) Table() string {
	return "admin"
}

func (m *Admin) SetID(id int) {
	m.Id = id
}

func (m *Admin) GetID() int {
	return m.Id
}

func (m *Admin) Columns() []string {
	return []string{"id", "name", "password", "status", "created_at", "updated_at", "remark"}
}

func (m *Admin) Fields() []interface{} {
	return []interface{}{&m.Id, &m.Name, &m.Password, &m.Status, &m.CreatedAt, &m.UpdatedAt, &m.Remark}
}

func (Admin) New() *Admin {
	return &Admin{}
}

type AdminRole struct {
	Id      int `json:"id"`       // 编号
	AdminId int `json:"admin_id"` // 账户编号
	RoleId  int `json:"role_id"`  // 角色编号
}

func (m *AdminRole) Table() string {
	return "admin_role"
}

func (m *AdminRole) SetID(id int) {
	m.Id = id
}

func (m *AdminRole) GetID() int {
	return m.Id
}

func (m *AdminRole) Columns() []string {
	return []string{"id", "admin_id", "role_id"}
}

func (m *AdminRole) Fields() []interface{} {
	return []interface{}{&m.Id, &m.AdminId, &m.RoleId}
}

func (AdminRole) New() *AdminRole {
	return &AdminRole{}
}

type Api struct {
	Id       int    `json:"id"`        // 编号
	Name     string `json:"name"`      // 名称
	Code     string `json:"code"`      // 编码
	ApiGroup string `json:"api_group"` // 分组
	Method   string `json:"method"`    // 方法
	Url      string `json:"url"`       // url
}

func (m *Api) Table() string {
	return "api"
}

func (m *Api) SetID(id int) {
	m.Id = id
}

func (m *Api) GetID() int {
	return m.Id
}

func (m *Api) Columns() []string {
	return []string{"id", "name", "code", "api_group", "method", "url"}
}

func (m *Api) Fields() []interface{} {
	return []interface{}{&m.Id, &m.Name, &m.Code, &m.ApiGroup, &m.Method, &m.Url}
}

func (Api) New() *Api {
	return &Api{}
}

type LogAdminLogin struct {
	Id        int        `json:"id"`         // 自增编号
	AdminId   int        `json:"admin_id"`   // 管理员编号
	Name      string     `json:"name"`       // 管理员名
	Location  string     `json:"location"`   // 位置
	Os        string     `json:"os"`         // 操作系统
	Browser   string     `json:"browser"`    // 浏览器
	UserAgent string     `json:"user_agent"` // 浏览器详情
	Url       string     `json:"url"`        // url
	Result    int        `json:"result"`     // 2:失败;1:成功
	Ip        string     `json:"ip"`         // IP
	RecordAt  xtime.Time `json:"record_at"`  // 记录时间
	Remark    string     `json:"remark"`     // 备注
}

func (m *LogAdminLogin) Table() string {
	return "log_admin_login"
}

func (m *LogAdminLogin) SetID(id int) {
	m.Id = id
}

func (m *LogAdminLogin) GetID() int {
	return m.Id
}

func (m *LogAdminLogin) Columns() []string {
	return []string{"id", "admin_id", "name", "location", "os", "browser", "user_agent", "url", "result", "ip", "record_at", "remark"}
}

func (m *LogAdminLogin) Fields() []interface{} {
	return []interface{}{&m.Id, &m.AdminId, &m.Name, &m.Location, &m.Os, &m.Browser, &m.UserAgent, &m.Url, &m.Result, &m.Ip, &m.RecordAt, &m.Remark}
}

func (LogAdminLogin) New() *LogAdminLogin {
	return &LogAdminLogin{}
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

func (m *LogAdminOperation) SetID(id int) {
	m.Id = id
}

func (m *LogAdminOperation) GetID() int {
	return m.Id
}

func (m *LogAdminOperation) Columns() []string {
	return []string{"id", "admin_id", "name", "operation_code", "operation_name", "content", "result", "message", "ip", "record_at"}
}

func (m *LogAdminOperation) Fields() []interface{} {
	return []interface{}{&m.Id, &m.AdminId, &m.Name, &m.OperationCode, &m.OperationName, &m.Content, &m.Result, &m.Message, &m.Ip, &m.RecordAt}
}

func (LogAdminOperation) New() *LogAdminOperation {
	return &LogAdminOperation{}
}

type Menu struct {
	Id        int    `json:"id"`         // 编号
	Name      string `json:"name"`       // 菜单名
	Code      string `json:"code"`       // 编码
	Pid       int    `json:"pid"`        // 上级菜单
	Icon      string `json:"icon"`       // icon
	Url       string `json:"url"`        // url
	IndexSort int    `json:"index_sort"` // 排序
}

func (m *Menu) Table() string {
	return "menu"
}

func (m *Menu) SetID(id int) {
	m.Id = id
}

func (m *Menu) GetID() int {
	return m.Id
}

func (m *Menu) Columns() []string {
	return []string{"id", "name", "code", "pid", "icon", "url", "index_sort"}
}

func (m *Menu) Fields() []interface{} {
	return []interface{}{&m.Id, &m.Name, &m.Code, &m.Pid, &m.Icon, &m.Url, &m.IndexSort}
}

func (Menu) New() *Menu {
	return &Menu{}
}

type Permission struct {
	Id              int    `json:"id"`
	Name            string `json:"name"`
	PermissionGroup string `json:"permission_group"`
	Remark          string `json:"remark"`
}

func (m *Permission) Table() string {
	return "permission"
}

func (m *Permission) SetID(id int) {
	m.Id = id
}

func (m *Permission) GetID() int {
	return m.Id
}

func (m *Permission) Columns() []string {
	return []string{"id", "name", "permission_group", "remark"}
}

func (m *Permission) Fields() []interface{} {
	return []interface{}{&m.Id, &m.Name, &m.PermissionGroup, &m.Remark}
}

func (Permission) New() *Permission {
	return &Permission{}
}

type PermissionApi struct {
	Id           int `json:"id"`
	PermissionId int `json:"permission_id"`
	ApiId        int `json:"api_id"`
}

func (m *PermissionApi) Table() string {
	return "permission_api"
}

func (m *PermissionApi) SetID(id int) {
	m.Id = id
}

func (m *PermissionApi) GetID() int {
	return m.Id
}

func (m *PermissionApi) Columns() []string {
	return []string{"id", "permission_id", "api_id"}
}

func (m *PermissionApi) Fields() []interface{} {
	return []interface{}{&m.Id, &m.PermissionId, &m.ApiId}
}

func (PermissionApi) New() *PermissionApi {
	return &PermissionApi{}
}

type PermissionMenu struct {
	Id           int `json:"id"`
	PermissionId int `json:"permission_id"`
	MenuId       int `json:"menu_id"`
}

func (m *PermissionMenu) Table() string {
	return "permission_menu"
}

func (m *PermissionMenu) SetID(id int) {
	m.Id = id
}

func (m *PermissionMenu) GetID() int {
	return m.Id
}

func (m *PermissionMenu) Columns() []string {
	return []string{"id", "permission_id", "menu_id"}
}

func (m *PermissionMenu) Fields() []interface{} {
	return []interface{}{&m.Id, &m.PermissionId, &m.MenuId}
}

func (PermissionMenu) New() *PermissionMenu {
	return &PermissionMenu{}
}

type Role struct {
	Id     int    `json:"id"`   // 编号
	Name   string `json:"name"` // 角色名称
	Remark string `json:"remark"`
}

func (m *Role) Table() string {
	return "role"
}

func (m *Role) SetID(id int) {
	m.Id = id
}

func (m *Role) GetID() int {
	return m.Id
}

func (m *Role) Columns() []string {
	return []string{"id", "name", "remark"}
}

func (m *Role) Fields() []interface{} {
	return []interface{}{&m.Id, &m.Name, &m.Remark}
}

func (Role) New() *Role {
	return &Role{}
}

type RoleMenu struct {
	Id     int `json:"id"`
	RoleId int `json:"role_id"`
	MenuId int `json:"menu_id"`
}

func (m *RoleMenu) Table() string {
	return "role_menu"
}

func (m *RoleMenu) SetID(id int) {
	m.Id = id
}

func (m *RoleMenu) GetID() int {
	return m.Id
}

func (m *RoleMenu) Columns() []string {
	return []string{"id", "role_id", "menu_id"}
}

func (m *RoleMenu) Fields() []interface{} {
	return []interface{}{&m.Id, &m.RoleId, &m.MenuId}
}

func (RoleMenu) New() *RoleMenu {
	return &RoleMenu{}
}

type RolePermission struct {
	Id           int `json:"id"`
	RoleId       int `json:"role_id"`
	PermissionId int `json:"permission_id"`
}

func (m *RolePermission) Table() string {
	return "role_permission"
}

func (m *RolePermission) SetID(id int) {
	m.Id = id
}

func (m *RolePermission) GetID() int {
	return m.Id
}

func (m *RolePermission) Columns() []string {
	return []string{"id", "role_id", "permission_id"}
}

func (m *RolePermission) Fields() []interface{} {
	return []interface{}{&m.Id, &m.RoleId, &m.PermissionId}
}

func (RolePermission) New() *RolePermission {
	return &RolePermission{}
}
