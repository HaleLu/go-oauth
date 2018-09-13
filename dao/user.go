package dao

import (
	"context"
	xsql "database/sql"
	"errors"

	"github.com/HaleLu/go-oauth/model"
)

const (
	_addUser = "INSERT INTO user (`username`,`password`,`nickname`,`is_two_step`)VALUES(?, ?, ?, 0)"
)

// AddUser add a new user
func (d *Dao) AddUser(c context.Context, u model.User) (id int64, err error) {
	var res xsql.Result
	if res, err = d.db.Exec(c, _addUser, u.Username, u.Password, u.Nickname); err != nil {
		err = errors.New("")
		return
	}
	return res.LastInsertId()
}
