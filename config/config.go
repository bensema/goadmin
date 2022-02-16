package config

import (
	"flag"
	"github.com/BurntSushi/toml"
	"library/cache/redis"
	"library/database/sql"
	"library/image/captcha"
)

var (
	confPath string
	Conf     = &Config{}
)

func init() {
	flag.StringVar(&confPath, "c", "goadmin.toml", "default config path")
}

type Config struct {
	Port      int
	MySQL     *sql.Config
	Redis     *redis.Config
	Captcha   *captcha.Config
	Ip2Region *Ip2Region
}

func local() (err error) {
	_, err = toml.DecodeFile(confPath, &Conf)
	return
}

func Init() error {
	return local()
}

type Ip2Region struct {
	Path string
}
