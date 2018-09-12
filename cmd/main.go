package main

import (
	"flag"
	"os"
	"os/signal"
	"syscall"

	"github.com/HaleLu/go-oauth/conf"
	"github.com/HaleLu/go-oauth/http"

	log "github.com/golang/glog"
)

func main() {
	flag.Parse()
	if err := conf.Init(); err != nil {
		log.Errorf("conf.Init() error(%v)", err)
		panic(err)
	}
	http.Init(conf.Conf)
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		s := <-c
		log.Infof("Receive a signal %s", s.String())
		switch s {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			log.Info("Exit...")
			log.Flush()
			return
		case syscall.SIGHUP:
		default:
			return
		}
	}

}
