package internal

import (
	"database/sql"
	entsql "entgo.io/ent/dialect/sql"
	"github.com/bensema/goadmin/model"
	"github.com/gin-gonic/gin"
	"strings"
)

func CreateAdminRole(c *gin.Context, db *sql.DB, obj *model.AdminRole) (err error) {
	builder := sqlBuilder()
	query, args := builder.Insert(obj.Table()).Columns(obj.Columns()...).Values(obj.Fields()...).Query()
	_, err = db.ExecContext(c, query, args...)
	return err
}

func DeleteAdminRoleById(c *gin.Context, db *sql.DB, id int) (err error) {
	obj := &model.AdminRole{}
	builder := sqlBuilder()
	query, args := builder.Delete(obj.Table()).Where(entsql.EQ("id", id)).Query()
	_, err = db.ExecContext(c, query, args...)
	return err
}

func UpdateAdminRoleById(c *gin.Context, db *sql.DB, id int, key string, value interface{}) (err error) {
	obj := &model.AdminRole{}
	builder := sqlBuilder()
	query, args := builder.Update(obj.Table()).Set(key, value).Where(entsql.EQ("id", id)).Query()
	_, err = db.ExecContext(c, query, args...)
	return
}

func GetAdminRoleById(c *gin.Context, db *sql.DB, id int) (obj *model.AdminRole, err error) {
	obj = &model.AdminRole{}
	builder := sqlBuilder()
	query, args := builder.Select(obj.Columns()...).From(entsql.Table(obj.Table())).Where(entsql.EQ("id", id)).Query()
	err = db.QueryRowContext(c, query, args...).Scan(obj.Fields()...)
	return
}
func FindAdminRole(c *gin.Context, db *sql.DB, req *model.FindAdminRoleReq) (objs []*model.AdminRole, err error) {
	objs = []*model.AdminRole{}
	query, args := buildSqlFindAdminRole(req, SqlFind)
	rows, err := db.QueryContext(c, query, args...)
	if err != nil {
		return
	}
	defer rows.Close()
	for rows.Next() {
		obj := &model.AdminRole{}
		err = rows.Scan(obj.Fields()...)
		if err != nil {
			return
		}
		objs = append(objs, obj)
	}
	return
}

func PageFindAdminRole(c *gin.Context, db *sql.DB, req *model.FindAdminRoleReq) (objs []*model.AdminRole, err error) {
	objs = []*model.AdminRole{}
	query, args := buildSqlFindAdminRole(req, SqlPageList)
	rows, err := db.QueryContext(c, query, args...)
	if err != nil {
		return
	}
	defer rows.Close()
	for rows.Next() {
		obj := &model.AdminRole{}
		err = rows.Scan(obj.Fields()...)
		if err != nil {
			return
		}
		objs = append(objs, obj)
	}
	return
}

func PageFindAdminRoleTotal(c *gin.Context, db *sql.DB, req *model.FindAdminRoleReq) (total int, err error) {
	total = 0
	query, args := buildSqlFindAdminRole(req, SqlPageCount)
	err = db.QueryRowContext(c, query, args...).Scan(&total)
	return
}

func buildSqlFindAdminRole(req *model.FindAdminRoleReq, sqlType string) (string, []interface{}) {
	obj := &model.AdminRole{}
	builder := sqlBuilder()
	selector := &entsql.Selector{}
	switch sqlType {
	case SqlPageList:
		selector = builder.Select(obj.Columns()...)
	case SqlFind:
		selector = builder.Select(obj.Columns()...)
	case SqlPageCount:
		selector = builder.Select("Count(*)")
	}

	selector = selector.From(entsql.Table(obj.Table()))

	// if req.Id != "" {
	//	  selector = selector.Where(entsql.EQ("id", req.Id))
	// }

	if req.AdminId != "" {
		selector = selector.Where(entsql.EQ("admin_id", req.AdminId))
	}

	if req.RoleId != 0 {
		selector = selector.Where(entsql.EQ("role_id", req.RoleId))
	}

	// count 返回
	if sqlType == SqlPageCount {
		return selector.Query()
	}
	if sqlType == SqlFind {
		return selector.Query()
	}

	_sort := ""
	switch req.Sort {
	case "desc":
		_sort = entsql.Desc(req.OrderBy)
	case "asc":
		_sort = entsql.Asc(req.OrderBy)
	default:
		_sort = entsql.Asc(req.OrderBy)
	}

	orderByList := strings.Split(req.OrderBy, ",")
	for _, orderBy := range orderByList {
		if checkInStr(obj.Columns(), orderBy) {
			selector = selector.OrderBy(_sort)
		}
	}

	selector.Offset((req.Num - 1) * req.Size)
	selector.Limit(req.Size)
	return selector.Query()
}
