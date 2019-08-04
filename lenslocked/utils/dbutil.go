package utils

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

const (
	host     = "127.0.0.1"
	port     = 3306
	user     = "root"
	password = "root"
	dbname   = "lenslocked_dev"
)

type Db struct {
	Db *gorm.DB
}

var internalDb *Db

func init() {
	internalDb = new(Db)
	mysqlInfo := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		user, password, host, port, dbname)
	db, err := gorm.Open("mysql", mysqlInfo)
	if err != nil {
		panic(err)
	}
	internalDb.Db = db
	db.LogMode(true)
}

func NewDB() *Db {
	return internalDb
}
