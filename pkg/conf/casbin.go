package conf

import "path"

type casbinConf struct {
	ModelFile string `toml:"model"`
}

func (c casbinConf) Model() string {
	return path.Join(pwd, c.ModelFile)
}
