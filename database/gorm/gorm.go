package gorm

import (
	"log"

	_ "github.com/go-sql-driver/mysql"

	"github.com/anerg2046/utils-go/config"

	"github.com/jinzhu/gorm"
)

// Gorm 数据库对象
var Gorm *gorm.DB

func init() {
	db, err := gorm.Open("mysql", Config.DB.User+":"+Config.DB.Pass+"@tcp("+Config.DB.Host+":"+Config.DB.Port+")/"+Config.DB.Name+"?charset=utf8")
	if err != nil {
		log.Fatal(err.Error())
		return
	}
	db.DB().SetMaxIdleConns(8)
	db.DB().SetMaxOpenConns(32)
	db.SingularTable(true)
	db.LogMode(false)
	Gorm = db
}
