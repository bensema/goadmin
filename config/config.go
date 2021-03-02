package config

import (
	"flag"
	"github.com/BurntSushi/toml"
	"github.com/bensema/library/cache/redis"
	"github.com/bensema/library/database/sql"
	"github.com/bensema/library/image/captcha"
	"github.com/bensema/library/log"
	"github.com/bensema/library/net/trace"
	"github.com/bensema/library/time"
)

var (
	confPath string
	Conf     = &Config{}
)

func init() {
	flag.StringVar(&confPath, "conf", "admin.toml", "default config path")
}

type Config struct {
	Web       *Web
	MySQL     *sql.Config
	Captcha   *captcha.Config
	Trace     *trace.Config
	Log       *log.Config
	Redis     *Redis
	Ip2Region *Ip2Region
}

func local() (err error) {
	_, err = toml.DecodeFile(confPath, &Conf)
	return
}

func Init() error {
	return local()
}

// Redis .
type Redis struct {
	*redis.Config
	AdminSessionExpire time.Duration
}

type Ip2Region struct {
	Path string
}

type Web struct {
	Port     int
	Dir      string
	Template string
	Static   string
}
