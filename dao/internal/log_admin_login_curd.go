package internal

import (
	"database/sql"
	entsql "entgo.io/ent/dialect/sql"
	"github.com/bensema/goadmin/model"
	"github.com/gin-gonic/gin"
	"strings"
)

func CreateLogAdminLogin(c *gin.Context, db *sql.DB, obj *model.LogAdminLogin) (err error) {
	builder := sqlBuilder()
	query, args := builder.Insert(obj.Table()).Columns(obj.Columns()...).Values(obj.Fields()...).Query()
	_, err = db.ExecContext(c, query, args...)
	return err
}

func DeleteLogAdminLoginById(c *gin.Context, db *sql.DB, id int) (err error) {
	obj := &model.LogAdminLogin{}
	builder := sqlBuilder()
	query, args := builder.Delete(obj.Table()).Where(entsql.EQ("id", id)).Query()
	_, err = db.ExecContext(c, query, args...)
	return err
}

func UpdateLogAdminLoginById(c *gin.Context, db *sql.DB, id int, key string, value interface{}) (err error) {
	obj := &model.LogAdminLogin{}
	builder := sqlBuilder()
	query, args := builder.Update(obj.Table()).Set(key, value).Where(entsql.EQ("id", id)).Query()
	_, err = db.ExecContext(c, query, args...)
	return
}

func GetLogAdminLoginById(c *gin.Context, db *sql.DB, id int) (obj *model.LogAdminLogin, err error) {
	obj = &model.LogAdminLogin{}
	builder := sqlBuilder()
	query, args := builder.Select(obj.Columns()...).From(entsql.Table(obj.Table())).Where(entsql.EQ("id", id)).Query()
	err = db.QueryRowContext(c, query, args...).Scan(obj.Fields()...)
	return
}
func FindLogAdminLogin(c *gin.Context, db *sql.DB, req *model.FindLogAdminLoginReq) (objs []*model.LogAdminLogin, err error) {
	objs = []*model.LogAdminLogin{}
	query, args := buildSqlFindLogAdminLogin(req, SqlFind)
	rows, err := db.QueryContext(c, query, args...)
	if err != nil {
		return
	}
	defer rows.Close()
	for rows.Next() {
		obj := &model.LogAdminLogin{}
		err = rows.Scan(obj.Fields()...)
		if err != nil {
			return
		}
		objs = append(objs, obj)
	}
	return
}

func PageFindLogAdminLogin(c *gin.Context, db *sql.DB, req *model.FindLogAdminLoginReq) (objs []*model.LogAdminLogin, err error) {
	objs = []*model.LogAdminLogin{}
	query, args := buildSqlFindLogAdminLogin(req, SqlPageList)
	rows, err := db.QueryContext(c, query, args...)
	if err != nil {
		return
	}
	defer rows.Close()
	for rows.Next() {
		obj := &model.LogAdminLogin{}
		err = rows.Scan(obj.Fields()...)
		if err != nil {
			return
		}
		objs = append(objs, obj)
	}
	return
}

func PageFindLogAdminLoginTotal(c *gin.Context, db *sql.DB, req *model.FindLogAdminLoginReq) (total int, err error) {
	total = 0
	query, args := buildSqlFindLogAdminLogin(req, SqlPageCount)
	err = db.QueryRowContext(c, query, args...).Scan(&total)
	return
}

func buildSqlFindLogAdminLogin(req *model.FindLogAdminLoginReq, sqlType string) (string, []interface{}) {
	obj := &model.LogAdminLogin{}
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

	if req.Id != 0 {
		selector = selector.Where(entsql.EQ("id", req.Id))
	}

	if req.AdminId != "" {
		selector = selector.Where(entsql.EQ("admin_id", req.AdminId))
	}
	if req.Name != "" {
		selector = selector.Where(entsql.EQ("name", req.Name))
	}

	if req.Ip != "" {
		selector = selector.Where(entsql.EQ("ip", req.Ip))
	}

	if req.Result != 0 {
		selector = selector.Where(entsql.EQ("result", req.Result))
	}

	if req.RecordAtFrom != 0 {
		selector = selector.Where(entsql.GTE("record_at", req.RecordAtFrom))
	}

	if req.RecordAtTo != 0 {
		selector = selector.Where(entsql.LT("record_at", req.RecordAtTo))
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
