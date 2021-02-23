package model

type FindAdminReply struct {
	Data []*Admin `json:"data"`
	PaginationReply
}

type FindAdminRoleReply struct {
	Data []*AdminRole `json:"data"`
	PaginationReply
}

type FindLogAdminLoginReply struct {
	Data []*LogAdminLogin `json:"data"`
	PaginationReply
}

type FindLogAdminOperationReply struct {
	Data []*LogAdminOperation `json:"data"`
	PaginationReply
}

type FindMenuReply struct {
	Data []*Menu `json:"data"`
	PaginationReply
}

type FindOperationReply struct {
	Data []*Operation `json:"data"`
	PaginationReply
}

type FindPermissionReply struct {
	Data []*Permission `json:"data"`
	PaginationReply
}

type FindPermissionMenuReply struct {
	Data []*PermissionMenu `json:"data"`
	PaginationReply
}

type FindPermissionOperationReply struct {
	Data []*PermissionOperation `json:"data"`
	PaginationReply
}

type FindRoleReply struct {
	Data []*Role `json:"data"`
	PaginationReply
}

type FindRolePermissionReply struct {
	Data []*RolePermission `json:"data"`
	PaginationReply
}
