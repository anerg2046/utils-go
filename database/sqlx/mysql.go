package sqlx

import (
	"log"

	"github.com/jmoiron/sqlx"

	. "anerg/config"

	_ "github.com/go-sql-driver/mysql"
)

// Mysql 数据库对象
var Mysql *sqlx.DB

func init() {
	db, err := sqlx.Connect("mysql", Config.DB.User+":"+Config.DB.Pass+"@tcp("+Config.DB.Host+":"+Config.DB.Port+")/"+Config.DB.Name+"?charset=utf8")
	if err != nil {
		log.Fatal(err.Error())
		return
	}
	db.SetMaxIdleConns(8)
	db.SetMaxOpenConns(32)
	Mysql = db
}
