package gorm

import (
	"log"

	_ "github.com/go-sql-driver/mysql"

	c "github.com/anerg2046/utils-go/config"

	"github.com/jinzhu/gorm"
)

// Gorm 数据库对象
var Gorm *gorm.DB

func init() {
	db, err := gorm.Open("mysql", c.Config.DB.User+":"+c.Config.DB.Pass+"@tcp("+c.Config.DB.Host+":"+c.Config.DB.Port+")/"+c.Config.DB.Name+"?charset=utf8")
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
