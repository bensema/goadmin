package model

import (
	"library/xtime"
)

// admin 数据v1
type AdminV1 struct {
	AdminId   int        `json:"admin_id"`   // 自增编号
	Name      string     `json:"name"`       // 账户
	Status    int        `json:"status"`     // 1:正常;2:禁用
	CreatedAt xtime.Time `json:"created_at"` // 创建时间
	UpdatedAt xtime.Time `json:"updated_at"` // 修改时间
	Roles     []*Role    `json:"roles"`      // 角色
}

type UpdateAdmin struct {
	AdminId int   `json:"admin_id"`
	Status  int   `json:"status"`
	Roles   []int `json:"roles"`
}

type AddAdmin struct {
	Name     string `json:"name"`
	Password string `json:"password"`
	Status   int    `json:"status"`
	Roles    []int  `json:"roles"`
}

type RoleInfo struct {
	Id    int     `json:"id"`
	Name  string  `json:"name"`
	Menus []*Menu `json:"menus"`
	Apis  []*Api  `json:"apis"`
}
