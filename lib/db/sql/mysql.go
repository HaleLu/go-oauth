package sql

import (
	"github.com/HaleLu/go-oauth/lib/time"
	// database driver
	_ "github.com/go-sql-driver/mysql"
	log "github.com/golang/glog"
)

// Config mysql config.
type Config struct {
	DSN          string        // write data source name.
	Active       int           // pool
	Idle         int           // pool
	IdleTimeout  time.Duration // connect max life time.
	QueryTimeout time.Duration // query sql timeout
	ExecTimeout  time.Duration // execute sql timeout
	TranTimeout  time.Duration // transaction sql timeout
}

// NewMySQL new db and retry connection when has error.
func NewMySQL(c *Config) (db *DB) {
	if c.QueryTimeout == 0 || c.ExecTimeout == 0 || c.TranTimeout == 0 {
		panic("mysql must be set query/execute/transction timeout")
	}
	db, err := Open(c)
	if err != nil {
		log.Errorf("open mysql error(%v)", err)
		panic(err)
	}
	return
}
