package conf

import "path"

type casbinConf struct {
	ModelFilePath string `toml:"model"`
}

func (c casbinConf) Model() string {
	return path.Join(Pwd, c.ModelFilePath)
}
