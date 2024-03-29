package service

import (
	"database/sql"
	"errors"
	"github.com/bensema/gcurd"
	"github.com/bensema/goadmin/dao"
	"github.com/bensema/goadmin/model"
	"github.com/bensema/goadmin/utils"
	"github.com/fatih/structs"
	"github.com/gin-gonic/gin"
	"library/xtime"
	"time"
)

func gPage[T gcurd.Model](c *gin.Context, db *sql.DB, obj T, req *gcurd.Request, f func() T) (reply *model.PageReply[T], err error) {

	var count int
	objs := make([]T, 0)
	reply = &model.PageReply[T]{}

	reply.Rows = objs
	reply.Page = req.Pagination.Page
	reply.PageSize = req.Pagination.PageSize

	if count, err = dao.PageTotal(c, db, obj, req); err != nil {
		return
	}

	if count <= 0 {
		return
	}

	if objs, err = dao.PageFind(c, db, obj, req, f); err != nil {
		return
	}

	reply.Rows = objs
	reply.RowsTotal = count

	return
}

func gCreate[T gcurd.Model](c *gin.Context, db *sql.DB, obj T, op model.OpCode, mosaicsColumns []string) error {
	res, err := dao.Create(c, db, obj)
	if err != nil {
		return err
	}
	id, _ := res.LastInsertId()
	obj.SetID(int(id))

	content := ContentNew(obj, mosaicsColumns)
	result := model.SUCCESS
	return logAction(c, db, string(op), content, result)
}

func gDelete[T gcurd.Model](c *gin.Context, db *sql.DB, obj T, id int, op model.OpCode, mosaicsColumns []string) error {
	old, err := dao.Get(c, db, obj, id)
	if err != nil {
		return errors.New("ID not found[-1]")
	}

	err = dao.Delete(c, db, old, id)
	if err != nil {
		return err
	}
	content := ContentNew(old, mosaicsColumns)
	result := model.SUCCESS
	return logAction(c, db, string(op), content, result)
}

func gGet[T gcurd.Model](c *gin.Context, db *sql.DB, obj T, id int) (T, error) {
	return dao.Get(c, db, obj, id)
}

func gFind[T gcurd.Model](c *gin.Context, db *sql.DB, obj T, wvs []*gcurd.WhereValue, f func() T) ([]T, error) {
	return dao.Find(c, db, obj, wvs, f)
}

func gUpdate[T gcurd.Model](c *gin.Context, db *sql.DB, obj T, id int, op model.OpCode, ignoreColumns []string, mosaicsColumns []string) error {
	n := structs.New(obj)
	n.TagName = "json"
	_new := n.Map()

	old, err := dao.Get(c, db, obj, id)
	if err != nil {
		return errors.New("id not exist")
	}

	o := structs.New(old)
	o.TagName = "json"
	_old := o.Map()

	var kvs []gcurd.KeyValue
	var wvs []*gcurd.WhereValue
	for _, col := range obj.Columns() {
		if utils.CheckIn(ignoreColumns, col) {
			continue
		}
		if _new[col] != _old[col] {
			kvs = append(kvs, gcurd.KeyValue{Key: col, Value: _new[col]})
		}
	}

	if kvs == nil {
		return errors.New("nothing has change")
	}

	wvs = append(wvs, gcurd.EQ("id", id))
	err = dao.UpdateWhereKV(c, db, obj, kvs, wvs)

	if err != nil {
		return err
	}

	content := logFieldTemp("id", id, nil, true, false) + ";" + ContentDiff(obj, old, mosaicsColumns)
	result := model.SUCCESS
	return logAction(c, db, string(op), content, result)
}

func logAction(c *gin.Context, db *sql.DB, opCode string, content string, result string) error {
	operatorInfo, err := GetAdminFromContext(c)
	if err != nil {
		return err
	}
	recordLog := &model.LogAdminOperation{
		AdminId:  operatorInfo.Id,
		Name:     operatorInfo.Name,
		Action:   opCode,
		Content:  content,
		Result:   result,
		Ip:       c.ClientIP(),
		RecordAt: xtime.Time(time.Now().Unix()),
	}
	_, err = dao.Create(c, db, recordLog)
	return err
}
