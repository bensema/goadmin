package main

import (
	"flag"
	"github.com/bensema/goadmin/config"
	"github.com/bensema/goadmin/server/http"
	"github.com/bensema/goadmin/service"
	"github.com/bensema/library/log"
	"github.com/bensema/library/net/trace"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	flag.Parse()
	if err := config.Init(); err != nil {
		panic(err)
	}
	log.Init(config.Conf.Log)

	flush := trace.Init(config.Conf.Trace)
	defer flush()

	srv := service.New(config.Conf)
	http.Init(config.Conf, srv)
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		s := <-c
		switch s {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGSTOP, syscall.SIGINT:
			time.Sleep(time.Second * 0)
			srv.Close()
			return
		case syscall.SIGHUP:
			// TODO reload
		default:
			return
		}
	}
}
