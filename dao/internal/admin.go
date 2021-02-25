package internal

import (
	"database/sql"
	entsql "entgo.io/ent/dialect/sql"
	"github.com/bensema/goadmin/model"
	"github.com/gin-gonic/gin"
)

func GetAdminByName(c *gin.Context, db *sql.DB, name string) (obj *model.Admin, err error) {
	obj = &model.Admin{}
	builder := sqlBuilder()
	query, args := builder.Select(obj.Columns()...).From(entsql.Table(obj.Table())).Where(entsql.EQ("name", name)).Query()
	err = db.QueryRowContext(c, query, args...).Scan(obj.Fields()...)
	return
}

func DeleteAdminRoleByAdminId(c *gin.Context, db *sql.DB, adminId int) (err error) {
	obj := &model.AdminRole{}
	builder := sqlBuilder()
	query, args := builder.Delete(obj.Table()).Where(entsql.EQ("admin_id", adminId)).Query()
	_, err = db.ExecContext(c, query, args...)
	return err
}

func DeleteRolePermissionByRoleId(c *gin.Context, db *sql.DB, roleId int) (err error) {
	obj := &model.RolePermission{}
	builder := sqlBuilder()
	query, args := builder.Delete(obj.Table()).Where(entsql.EQ("role_id", roleId)).Query()
	_, err = db.ExecContext(c, query, args...)
	return err
}

func GetRoleByName(c *gin.Context, db *sql.DB, name string) (obj *model.Role, err error) {
	obj = &model.Role{}
	builder := sqlBuilder()
	query, args := builder.Select(obj.Columns()...).From(entsql.Table(obj.Table())).Where(entsql.EQ("name", name)).Query()
	err = db.QueryRowContext(c, query, args...).Scan(obj.Fields()...)
	return
}

func GetPermissionByName(c *gin.Context, db *sql.DB, name string) (obj *model.Permission, err error) {
	obj = &model.Permission{}
	builder := sqlBuilder()
	query, args := builder.Select(obj.Columns()...).From(entsql.Table(obj.Table())).Where(entsql.EQ("name", name)).Query()
	err = db.QueryRowContext(c, query, args...).Scan(obj.Fields()...)
	return
}

func DeletePermissionMenuByPermissionId(c *gin.Context, db *sql.DB, id int) (err error) {
	obj := &model.PermissionMenu{}
	builder := sqlBuilder()
	query, args := builder.Delete(obj.Table()).Where(entsql.EQ("permission_id", id)).Query()
	_, err = db.ExecContext(c, query, args...)
	return err
}

func DeletePermissionOperationByPermissionId(c *gin.Context, db *sql.DB, id int) (err error) {
	obj := &model.PermissionOperation{}
	builder := sqlBuilder()
	query, args := builder.Delete(obj.Table()).Where(entsql.EQ("permission_id", id)).Query()
	_, err = db.ExecContext(c, query, args...)
	return err
}

func GetMenuByName(c *gin.Context, db *sql.DB, name string) (obj *model.Menu, err error) {
	obj = &model.Menu{}
	builder := sqlBuilder()
	query, args := builder.Select(obj.Columns()...).From(entsql.Table(obj.Table())).Where(entsql.EQ("name", name)).Query()
	err = db.QueryRowContext(c, query, args...).Scan(obj.Fields()...)
	return
}
func GetOperationByName(c *gin.Context, db *sql.DB, name string) (obj *model.Operation, err error) {
	obj = &model.Operation{}
	builder := sqlBuilder()
	query, args := builder.Select(obj.Columns()...).From(entsql.Table(obj.Table())).Where(entsql.EQ("name", name)).Query()
	err = db.QueryRowContext(c, query, args...).Scan(obj.Fields()...)
	return
}
