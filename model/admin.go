package model

import xtime "github.com/bensema/library/time"

type AdminSession struct {
	AdminId string
	Name    string
	AesKey  string
}

type AdminLoginV2 struct {
	Username string     `json:"username"`
	Password string     `json:"password"`
	AesKey   string     `json:"aesKey"`
	T        xtime.Time `json:"t"`
}

// admin 数据v1
type AdminV1 struct {
	AdminId   string     `json:"admin_id"`   // 自增编号
	Name      string     `json:"name"`       // 账户
	Status    int        `json:"status"`     // 1:正常;2:禁用
	CreatedAt xtime.Time `json:"created_at"` // 创建时间
	UpdatedAt xtime.Time `json:"updated_at"` // 修改时间
	Roles     []*Role    `json:"roles"`      // 角色
}

type FindAdminReplyV1 struct {
	Data []*AdminV1 `json:"data"`
	PaginationReply
}

type UpdateAdmin struct {
	AdminId string `json:"admin_id"`
	Status  int    `json:"status"`
	Roles   []int  `json:"roles"`
}

type AddAdmin struct {
	Name     string `json:"name"`
	Password string `json:"password"`
	Status   int    `json:"status"`
	Roles    []int  `json:"roles"`
}

// role 数据v1
type RoleV1 struct {
	Id          int           `json:"id"`
	Name        string        `json:"name"`
	Permissions []*Permission `json:"permissions"`
}

type FindRoleReplyV1 struct {
	Data []*RoleV1 `json:"data"`
	PaginationReply
}

type UpdateRole struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Permissions []int  `json:"permissions"`
}

type AddRole struct {
	Name        string `json:"name"`
	Permissions []int  `json:"permissions"`
}

type UpdatePermission struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	Menus     []int  `json:"menus"`
	Operation []int  `json:"operation"`
}

type FindLogAdminLoginReq struct {
	LogAdminLogin
	RecordAtFrom xtime.Time `json:"record_at_from"`
	RecordAtTo   xtime.Time `json:"record_at_to"`
	OrderBy      string     `json:"order_by"`
	Sort         string     `json:"sort"`
	Pagination
}

type FindLogAdminOperationReq struct {
	LogAdminOperation
	OrderBy      string     `json:"order_by"`
	Sort         string     `json:"sort"`
	RecordAtFrom xtime.Time `json:"record_at_from"`
	RecordAtTo   xtime.Time `json:"record_at_to"`
	Pagination
}
