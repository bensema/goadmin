package sql

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

type Config struct {
	DSN         string
	ReadDSN     []string
	Active      int
	Idle        int
	IdleTimeout int
}

func New(c *Config) (db *sql.DB) {
	db, _ = sql.Open("mysql", c.DSN)
	db.SetMaxOpenConns(c.Active)
	db.SetMaxIdleConns(c.Idle)
	return
}
