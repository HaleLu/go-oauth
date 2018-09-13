package service

import (
	"github.com/HaleLu/go-oauth/conf"
	"github.com/HaleLu/go-oauth/dao"
)

// Service struct of service.
type Service struct {
	d *dao.Dao
	// conf
	c *conf.Config
}

// New create service instance and return.
func New(c *conf.Config) (s *Service) {
	s = &Service{
		c: c,
		d: dao.New(c),
	}
	return
}
