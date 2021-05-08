package model

import "errors"

const (
	RedisPrefixAdminSession = "admin-session"
	RedisPrefixCaptcha      = "captcha"
)

var (
	ErrNil = errors.New("nil")
)
