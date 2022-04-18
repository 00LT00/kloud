package model

type User struct {
	ID    string
	Name  string
	Phone string
	Pass  string
}

func GetModels() []interface{} {
	return []interface{}{&User{}}
}
