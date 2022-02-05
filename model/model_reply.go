package model

type PageReply[T any] struct {
	Data []T `json:"data"`
	PaginationReply
}
