package model

func GetModels() []interface{} {
	return []interface{}{&User{}, &App{}, &Resource{}, &Flow{}}
}
