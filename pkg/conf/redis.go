package conf

type redisConf struct {
	IP, Port string
	Pass     string
}

func (c redisConf) Addr() string {
	return c.IP + ":" + c.Port
}
