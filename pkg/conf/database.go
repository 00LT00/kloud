package conf

import "fmt"

type databaseConf struct {
	User     string
	Pass     string
	Port     string
	Addr     string
	DBName   string
	Protocol string
	Params   []struct {
		Key   string
		Value string
	}
}

func (c databaseConf) Dsn() string {
	dsn := fmt.Sprintf("%s:%s@%s(%s:%s)/%s", c.User, c.Pass, c.Protocol, c.Addr, c.Port, c.DBName)
	if len(c.Params) != 0 {
		dsn += fmt.Sprintf("?%s=%s", c.Params[0].Key, c.Params[0].Value)
		for i := 1; i < len(c.Params); i++ {
			dsn += fmt.Sprintf("&%s=%s", c.Params[i].Key, c.Params[i].Value)
		}
	}
	return dsn
}
