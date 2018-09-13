package dao

import (
	"github.com/HaleLu/go-oauth/conf"
	"github.com/HaleLu/go-oauth/lib/db/sql"
)

// Dao struct answer history of Dao
type Dao struct {
	c *conf.Config
	// db
	db *sql.DB
}

// New new a Dao and return.
func New(c *conf.Config) (d *Dao) {
	d = &Dao{
		c:  c,
		db: sql.NewMySQL(c.DB),
	}
	return
}
