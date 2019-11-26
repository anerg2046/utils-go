package sqlx

import (
	"log"

	"github.com/jmoiron/sqlx"

	c "github.com/anerg2046/utils-go/config"

	_ "github.com/go-sql-driver/mysql"
)

// Mysql 数据库对象
var Mysql *sqlx.DB

func init() {
	db, err := sqlx.Connect("mysql", c.Config.DB.User+":"+c.Config.DB.Pass+"@tcp("+c.Config.DB.Host+":"+c.Config.DB.Port+")/"+c.Config.DB.Name+"?charset=utf8")
	if err != nil {
		log.Fatal(err.Error())
		return
	}
	db.SetMaxIdleConns(8)
	db.SetMaxOpenConns(32)
	Mysql = db
}
