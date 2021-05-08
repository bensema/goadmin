package redis

import (
	"library/xtime"
)

type Config struct {
	Addr         string
	Db           int
	Password     string
	DialTimeout  xtime.Duration
	ReadTimeout  xtime.Duration
	WriteTimeout xtime.Duration

	MaxIdle         int
	MaxActive       int
	IdleTimeout     xtime.Duration
	MaxConnLifetime xtime.Duration
}
