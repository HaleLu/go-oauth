package service

import (
	"context"

	"github.com/HaleLu/go-authenticator/encrypt"
	"github.com/HaleLu/go-oauth/errors"
	"github.com/HaleLu/go-oauth/model"
	log "github.com/golang/glog"
)

// Login check user's username and password and return a token
func (s *Service) Login(c context.Context, params *model.LoginParams) (token string, err error) {
	var (
		user *model.User
		code int
	)
	if user, err = s.d.RawUserByNameAndPwd(c, params.Username, params.Password); err != nil {
		log.Errorf("RawUserByNameAndPwd err:%v", err)
	}
	if user == nil {
		err = errors.Unauthorized
		return
	}
	if !user.IsTwoStep {
		if token, err = s.genToken(c, user.ID); err != nil {
			log.Errorf("genToken err:%v", err)
			err = errors.Unauthorized
			return
		}
		return
	}
	if code, err = encrypt.GetCodeNow(user.Secret); err != nil {
		err = errors.Unauthorized
		return
	}
	if params.Code != code {
		err = errors.Unauthorized
		return
	}
	if token, err = s.genToken(c, user.ID); err != nil {
		log.Errorf("genToken err:%v", err)
		err = errors.Unauthorized
		return
	}
	return
}

func (s *Service) genToken(c context.Context, id int64) (token string, err error) {
	return "token", nil
}
