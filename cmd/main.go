package main

import (
	"flag"
	"github.com/bensema/goadmin/config"
	"github.com/bensema/goadmin/server/http"
	"github.com/bensema/goadmin/service"
	log "github.com/sirupsen/logrus"
	"io"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"
	"time"
)

func init() {
	log.SetFormatter(&log.TextFormatter{})
	log.SetOutput(os.Stdout)
	file, err := os.OpenFile(filepath.Join(".", "goadmin.log"), os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	writers := []io.Writer{file, os.Stdout}
	fileAndStdoutWriter := io.MultiWriter(writers...)
	if err == nil {
		log.SetOutput(fileAndStdoutWriter)
	} else {
		log.Info("failed to log to file.")
	}
}

func main() {
	flag.Parse()
	if err := config.Init(); err != nil {
		panic(err)
	}
	srv := service.New(config.Conf)
	log.Error("test")
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
