package model

type BBAdminApi struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	T       int64       `json:"t"`
	R       string      `json:"r"`
}

type AdminApiReply struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	T       int64       `json:"t"`
	R       string      `json:"r"`
}
