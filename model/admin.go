package model

import (
	"library/xtime"
)

type AdminV1 struct {
	Id        int        `json:"id"`         // 自增编号
	Name      string     `json:"name"`       // 账户
	Status    string     `json:"status"`     // 1:正常;2:禁用
	CreatedAt xtime.Time `json:"created_at"` // 创建时间
	UpdatedAt xtime.Time `json:"updated_at"` // 修改时间
	Remark    string     `json:"remark"`
	Roles     []*Role    `json:"roles"` // 角色
}

type UpdateAdmin struct {
	AdminId int    `json:"admin_id"`
	Status  string `json:"status"`
	Remark  string `json:"remark"`
	Roles   []int  `json:"roles"`
}

type AddAdmin struct {
	Name     string `json:"name"`
	Password string `json:"password"`
	Status   string `json:"status"`
	Roles    []int  `json:"roles"`
}

type AddRole struct {
	Name        string `json:"name"`
	Remark      string `json:"remark"`
	Permissions []int  `json:"permissions"`
}

type RoleInfo struct {
	Id          int           `json:"id"`
	Name        string        `json:"name"`
	Remark      string        `json:"remark"`
	Permissions []*Permission `json:"permissions"`
}

type UpdateRole struct {
	RoleId      int    `json:"role_id"`
	Name        string `json:"name"`
	Remark      string `json:"remark"`
	Permissions []int  `json:"permissions"`
}

type PermissionInfo struct {
	Id              int     `json:"id"`
	Name            string  `json:"name"`
	PermissionGroup string  `json:"Permission_group"`
	Remark          string  `json:"remark"`
	Menus           []*Menu `json:"menus"`
	Apis            []*Api  `json:"apis"`
}

type UpdatePermission struct {
	Id              int    `json:"id"`
	Name            string `json:"name"`
	PermissionGroup string `json:"Permission_group"`
	Remark          string `json:"remark"`
	Menus           []int  `json:"menus"`
	Apis            []int  `json:"apis"`
}

type AddPermission struct {
	Name            string `json:"name"`
	PermissionGroup string `json:"Permission_group"`
	Remark          string `json:"remark"`
	Menus           []int  `json:"menus"`
	Apis            []int  `json:"apis"`
}
