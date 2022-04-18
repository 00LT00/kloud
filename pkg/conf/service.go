package conf

type serviceConf struct {
	IP, Port string
}

func (c serviceConf) Addr() string {
	return c.IP + ":" + c.Port
}
