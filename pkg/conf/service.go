package conf

type serviceConf struct {
	IP, Port string
	Secret   string
}

func (c serviceConf) Addr() string {
	return c.IP + ":" + c.Port
}
