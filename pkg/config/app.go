package config

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/subosito/gotenv"
	"os"
)

var (
	db *gorm.DB
)

func init() {
	err := gotenv.Load()
	if err != nil {
		fmt.Println("Fatal error: cannot load environment variables")
		return
	}
}

func Connect() {
	// use your local configuration for the testing db (this case is mysql)
	d, err := gorm.Open("mysql", os.Getenv("MYSQL_USERNAME")+":"+os.Getenv("MYSQL_PASSWORD")+"@("+os.Getenv("MYSQL_HOST")+")/"+os.Getenv("MYSQL_DATABASE")+"?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	db = d
}

func GetDB() *gorm.DB {
	return db
}
