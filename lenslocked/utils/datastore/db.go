package datastore

import (
	"fmt"

	"github.com/imeraj/go_playground/lenslocked/utils/paths"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/spf13/viper"
)

var (
	configdir  = "configs"
	configfile = "dbconfig"
)

type DbConfig struct {
	Host     string
	Port     uint
	Username string
	Password string
	DbName   string
}

type Db struct {
	Db *gorm.DB
}

var internalDb *Db

func init() {
	internalDb = new(Db)
	dbConfig := readDbConfig()
	mysqlInfo := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		dbConfig.Username, dbConfig.Password, dbConfig.Host, dbConfig.Port, dbConfig.DbName)
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

func readDbConfig() DbConfig {
	var dbConfig DbConfig
	var config string
	config = paths.BuildPath(configfile)

	viper.SetConfigName(config)
	viper.AddConfigPath(configdir)
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	if err := viper.Unmarshal(&dbConfig); err != nil {
		panic(err)
	}

	return dbConfig
}
