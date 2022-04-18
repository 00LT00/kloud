package conf

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"os"
	"path"
)

const (
	Info  = "info"
	Debug = "debug"
)

type config struct {
	Service serviceConf  `toml:"service"`
	DB      databaseConf `toml:"database"`
	Casbin  casbinConf   `toml:"casbin"`
	Redis   redisConf    `toml:"redis"`
	Mode    string
}

var c = new(config)
var pwd string

func init() {
	var err error
	pwd, err = os.Getwd()
	if err != nil {
		panic(err)
	}
	_, err = toml.DecodeFile(path.Join(pwd, "config", "kloud.toml"), c)
	if err != nil {
		panic(err)
	}

}

func GetConf() *config {
	return c
}

func ShowAllConfig() {
	fmt.Printf("conf: %+v\n", *c)
	fmt.Printf("Dsn: %s\n", c.DB.Dsn())
	fmt.Printf("model file: %s\n", c.Casbin.Model())
	fmt.Printf("redis addr: %s\n", c.Redis.Addr())
	fmt.Printf("service addr: %s\n", c.Service.Addr())
}
