package model

type Result string

const (
	SUCCESS = "SUCCESS"
	FAIL    = "FAIL"
)

type PageReply[T any] struct {
	Rows      []T `json:"rows"`
	PageSize  int `json:"page_size"`
	Page      int `json:"page"`
	RowsTotal int `json:"rows_total"`
}

type GinSession struct {
	AdminId int
	Name    string
}
