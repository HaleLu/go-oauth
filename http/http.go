package http

import (
	"time"

	"github.com/HaleLu/go-oauth/conf"
	"github.com/HaleLu/go-oauth/errors"
	"github.com/HaleLu/go-oauth/service"
	"github.com/gin-gonic/gin"
	log "github.com/golang/glog"
)

const (
	// 保存错误码的key
	contextErrCode = "context/err/code"
)

var (
	svr *service.Service
)

// Init init http.
func Init(c *conf.Config) {
	svr = service.New(c)

	engine := gin.New()
	engine.Use(loggerHandler)
	outerRouter(engine)
	go engine.Run(c.HTTPServer.Addr)
}

// outerRouter init local router api path.
func outerRouter(e *gin.Engine) {
	group := e.Group("/api")
	{
		group.GET("ping", ping)
		group.POST("login", login)
	}
}

func loggerHandler(c *gin.Context) {
	// Start timer
	start := time.Now()
	path := c.Request.URL.Path
	raw := c.Request.URL.RawQuery
	method := c.Request.Method

	// Process request
	c.Next()

	// Stop timer
	end := time.Now()
	latency := end.Sub(start)
	statusCode := c.Writer.Status()
	ecode := c.GetInt(contextErrCode)
	clientIP := c.ClientIP()
	if raw != "" {
		path = path + "?" + raw
	}
	log.Infof("METHOD:%s | PATH:%s | CODE:%d | IP:%s | TIME:%d | ECODE:%d", method, path, statusCode, clientIP, latency/time.Millisecond, ecode)
}

// ping shows an example to handle requests.
// More detail in https://github.com/gin-gonic/gin#api-examples
func ping(c *gin.Context) {
	result(c, map[string]interface{}{
		"message": "pong",
	}, nil)
}

type resp struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
}

func result(c *gin.Context, data interface{}, err error) {
	ee := errors.Code(err)
	c.Set(contextErrCode, ee.Code())
	c.JSON(200, resp{
		Code: ee.Code(),
		Data: data,
	})
}
