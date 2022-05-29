package conf

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"k8s.io/client-go/util/homedir"
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
	K8s     k8sConf      `toml:"k8s"`
	Mode    string
}

var c = new(config)
var Pwd string

func init() {
	var err error
	Pwd, err = os.Getwd()
	if err != nil {
		panic(err)
	}
	_, err = toml.DecodeFile(path.Join(Pwd, "config", "kloud.toml"), c)
	if err != nil {
		panic(err)
	}
	if c.K8s.ConfigPath == "" {
		home := homedir.HomeDir()
		if home == "" {
			panic("kubeconfig not found")
		}
		c.K8s.ConfigPath = path.Join(home, ".kube", "config")
	}
	if c.K8s.Namespace == "" {
		c.K8s.Namespace = "default"
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
	fmt.Printf("k8s Config: %s, namespace: %s\n", c.K8s.ConfigPath, c.K8s.Namespace)
}
