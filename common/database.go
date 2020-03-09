package common

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"xietong.me/ginessential/model"
)

//var DB *gorm.DB

func InitDB() *gorm.DB {
	driverName := "mysql"
	host := "localhost"
	port := "3306"
	database := "ginessential"
	username := "root"
	password := ""
	charset := "utf8"
	args := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true",
		username,
		password,
		host,
		port,
		database,
		charset)
	db, err := gorm.Open(driverName, args)
	if err != nil {
		panic("fail to connect databse,err:" + err.Error())
	}
	db.AutoMigrate(&model.User{})
	DB = db
	return db
}

//func GetDB() *gorm.DB {
//	return DB
//}
