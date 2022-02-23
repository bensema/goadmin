package dao

import (
	"database/sql"
	"github.com/bensema/gcurd"
	"github.com/gin-gonic/gin"
)

func init() {
	gcurd.Level = gcurd.Debug
}

func Create[T gcurd.Model](c *gin.Context, db *sql.DB, obj T) (sql.Result, error) {
	return gcurd.Create(c, db, obj)
}

func Delete[T gcurd.Model](c *gin.Context, db *sql.DB, obj T, id int) error {
	return gcurd.Delete(c, db, obj, id)
}

func Update[T gcurd.Model](c *gin.Context, db *sql.DB, obj T, id int, key string, value interface{}) error {
	return gcurd.Update(c, db, obj, id, key, value)
}

func UpdateWhere[T gcurd.Model](c *gin.Context, db *sql.DB, obj T, key string, value interface{}, wvs []*gcurd.WhereValue) error {
	return gcurd.UpdateWhere(c, db, obj, key, value, wvs)
}

func UpdateWhereKV[T gcurd.Model](c *gin.Context, db *sql.DB, obj T, kvs []gcurd.KeyValue, wvs []*gcurd.WhereValue) error {
	return gcurd.UpdateWhereKV(c, db, obj, kvs, wvs)
}

func First[T gcurd.Model](c *gin.Context, db *sql.DB, obj T) (T, error) {
	return gcurd.First(c, db, obj)
}

func Last[T gcurd.Model](c *gin.Context, db *sql.DB, obj T) (T, error) {
	return gcurd.Last(c, db, obj)
}

func Get[T gcurd.Model](c *gin.Context, db *sql.DB, obj T, id int) (T, error) {
	return gcurd.Get(c, db, obj, id)
}

func GetWhere[T gcurd.Model](c *gin.Context, db *sql.DB, obj T, wvs []*gcurd.WhereValue) (T, error) {
	return gcurd.GetWhere(c, db, obj, wvs)
}

func Find[T gcurd.Model](c *gin.Context, db *sql.DB, obj T, wvs []*gcurd.WhereValue, f func() T) ([]T, error) {
	return gcurd.Find(c, db, obj, wvs, f)
}

func PageTotal[T gcurd.Model](c *gin.Context, db *sql.DB, obj T, req *gcurd.Request) (int, error) {
	return gcurd.PageTotal(c, db, obj, req)
}

func PageFind[T gcurd.Model](c *gin.Context, db *sql.DB, obj T, req *gcurd.Request, f func() T) ([]T, error) {
	return gcurd.PageFind(c, db, obj, req, f)
}
