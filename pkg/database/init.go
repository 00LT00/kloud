package database

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"kloud/model"
	"kloud/pkg/conf"
)

var db = new(gorm.DB)

func init() {
	var err error
	db, err = gorm.Open(mysql.Open(conf.GetConf().DB.Dsn()))
	if err != nil {
		panic(err)
	}
	if conf.GetConf().Mode == conf.Debug {
		db = db.Debug()
	}
	initTable()
}

func initTable() {
	err := db.AutoMigrate(model.GetModels()...)
	if err != nil {
		panic(err)
	}
}

func GetDB() *gorm.DB {
	return db
}

func Ping() {
	sql, _ := db.DB()
	err := sql.Ping()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("connect to DB")
	}
}
