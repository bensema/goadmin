package model

import xtime "github.com/bensema/library/time"

type AdminSession struct {
	UserId int
	Name   string
}

// admin 数据v1
type AdminV1 struct {
	Id        int        `json:"id"`         // 自增编号
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
	Id     int   `json:"id"`
	Status int   `json:"status"`
	Roles  []int `json:"roles"`
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
