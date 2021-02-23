package model

type FindAdminReq struct {
	Admin
	OrderBy string `json:"order_by"`
	Sort    string `json:"sort"`
	Pagination
}

type FindAdminRoleReq struct {
	AdminRole
	OrderBy string `json:"order_by"`
	Sort    string `json:"sort"`
	Pagination
}

type FindLogAdminLoginReq struct {
	LogAdminLogin
	OrderBy string `json:"order_by"`
	Sort    string `json:"sort"`
	Pagination
}

type FindLogAdminOperationReq struct {
	LogAdminOperation
	OrderBy string `json:"order_by"`
	Sort    string `json:"sort"`
	Pagination
}

type FindMenuReq struct {
	Menu
	OrderBy string `json:"order_by"`
	Sort    string `json:"sort"`
	Pagination
}

type FindOperationReq struct {
	Operation
	OrderBy string `json:"order_by"`
	Sort    string `json:"sort"`
	Pagination
}

type FindPermissionReq struct {
	Permission
	OrderBy string `json:"order_by"`
	Sort    string `json:"sort"`
	Pagination
}

type FindPermissionMenuReq struct {
	PermissionMenu
	OrderBy string `json:"order_by"`
	Sort    string `json:"sort"`
	Pagination
}

type FindPermissionOperationReq struct {
	PermissionOperation
	OrderBy string `json:"order_by"`
	Sort    string `json:"sort"`
	Pagination
}

type FindRoleReq struct {
	Role
	OrderBy string `json:"order_by"`
	Sort    string `json:"sort"`
	Pagination
}

type FindRolePermissionReq struct {
	RolePermission
	OrderBy string `json:"order_by"`
	Sort    string `json:"sort"`
	Pagination
}
