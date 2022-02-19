package model

type PageReply[T any] struct {
	Rows []T `json:"rows"`
	PaginationReply
}
