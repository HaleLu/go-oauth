package conf

import (
	"flag"

	"github.com/BurntSushi/toml"
	"github.com/HaleLu/go-oauth/lib/db/sql"
)

var (
	confPath string
	//Conf conf
	Conf = &Config{}
)

// Config config
type Config struct {
	Author     string
	HTTPServer *ServerConfig
	DB         *sql.Config
}

// DB db config.
type DB sql.Config

// ServerConfig present HTTP server conf
type ServerConfig struct {
	Addr string
}

func init() {
	flag.StringVar(&confPath, "conf", "default.toml", "config path")
}

// Init init conf
func Init() (err error) {
	_, err = toml.DecodeFile(confPath, &Conf)
	return
}
