package dao

import (
	"context"
	"sync"
	"time"

	"github.com/HaleLu/go-oauth/conf"
)

var (
	once sync.Once
	d    *Dao
	ctx  = context.TODO()
)

func initConf() {
	if err := conf.Init(); err != nil {
		panic(err)
	}
}

func startService() {
	initConf()
	d = New(conf.Conf)
	time.Sleep(time.Second * 2)
}
