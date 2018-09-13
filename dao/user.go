package dao

import (
	"context"
	xsql "database/sql"

	"github.com/HaleLu/go-oauth/errors"
	"github.com/HaleLu/go-oauth/model"
	log "github.com/golang/glog"
)

//   CREATE TABLE `user` (
// 	`id` int(11) NOT NULL AUTO_INCREMENT,
// 	`username` varchar(32) CHARACTER SET latin1 NOT NULL,
// 	`password` varchar(32) CHARACTER SET latin1 NOT NULL,
// 	`nickname` varchar(64) NOT NULL DEFAULT '',
// 	`is_two_step` binary(1) NOT NULL DEFAULT '0',
// 	`secret` varchar(32) CHARACTER SET latin1 NOT NULL DEFAULT '',
// 	PRIMARY KEY (`id`),
// 	UNIQUE KEY `uk_username` (`username`) USING BTREE
//   ) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4;

const (
	_addUser    = "INSERT INTO `user` (`username`,`password`,`nickname`,`is_two_step`,`secret`)VALUES(?, ?, ?, ?, ?)"
	_updateUser = "UPDATE `user` SET `password`=?,`nickname`=?,`is_two_step`=?,`secret`=? WHERE `id`=?"
)

// AddUser add a new user
func (d *Dao) AddUser(c context.Context, u *model.User) (id int64, err error) {
	var res xsql.Result
	if res, err = d.db.Exec(c, _addUser, u.Username, u.Password, u.Nickname, u.IsTwoStep, u.Secret); err != nil {
		log.Errorf("add user error: %v", err)
		err = errors.NotModified
		return
	}
	return res.LastInsertId()
}

// UpdateUser update a user
func (d *Dao) UpdateUser(c context.Context, u *model.User) (rows int64, err error) {
	var res xsql.Result
	if res, err = d.db.Exec(c, _updateUser, u.Password, u.Nickname, u.IsTwoStep, u.Secret, u.ID); err != nil {
		log.Errorf("update user error: %v", err)
		err = errors.NotModified
		return
	}
	return res.RowsAffected()
}
