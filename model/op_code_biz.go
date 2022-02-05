package model

type OpCode string

const (
	OpAdminAdd    OpCode = "admin_add"
	OpAdminDel    OpCode = "admin_del"
	OpAdminUpdate OpCode = "admin_update"

	OpAdminRoleAdd    OpCode = "admin_role_add"
	OpAdminRoleDel    OpCode = "admin_role_del"
	OpAdminRoleUpdate OpCode = "admin_role_update"

	OpApiAdd    OpCode = "api_add"
	OpApiDel    OpCode = "api_del"
	OpApiUpdate OpCode = "api_update"

	OpLogAdminLoginAdd    OpCode = "log_admin_login_add"
	OpLogAdminLoginDel    OpCode = "log_admin_login_del"
	OpLogAdminLoginUpdate OpCode = "log_admin_login_update"

	OpLogAdminOperationAdd    OpCode = "log_admin_operation_add"
	OpLogAdminOperationDel    OpCode = "log_admin_operation_del"
	OpLogAdminOperationUpdate OpCode = "log_admin_operation_update"

	OpMenuAdd    OpCode = "menu_add"
	OpMenuDel    OpCode = "menu_del"
	OpMenuUpdate OpCode = "menu_update"

	OpRoleAdd    OpCode = "role_add"
	OpRoleDel    OpCode = "role_del"
	OpRoleUpdate OpCode = "role_update"

	OpRoleApiAdd    OpCode = "role_api_add"
	OpRoleApiDel    OpCode = "role_api_del"
	OpRoleApiUpdate OpCode = "role_api_update"

	OpRoleMenuAdd    OpCode = "role_menu_add"
	OpRoleMenuDel    OpCode = "role_menu_del"
	OpRoleMenuUpdate OpCode = "role_menu_update"
)
