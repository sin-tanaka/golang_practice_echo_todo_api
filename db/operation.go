package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/go-sql-driver/mysql"
)

func Opendb() (db *gorm.DB) {
	db, err := gorm.Open("mysql", "root:@/echo_sample?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("failed to connect database")
	}
	return db
}