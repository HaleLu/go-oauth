package http

import (
	"github.com/HaleLu/go-oauth/model"
	"github.com/gin-gonic/gin"
)

func login(c *gin.Context) {
	params := new(model.LoginParams)
	c.Bind(params)

	result(c, map[string]interface{}{
		"result": true,
	}, nil)
}
