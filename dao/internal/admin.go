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
