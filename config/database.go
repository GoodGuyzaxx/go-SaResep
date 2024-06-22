package config

import (
	"fmt"

	"gorm.io/drive/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func loadDataBase() {

	dbConnectionStr := fmt.Sprintf("%v:%v@tcp(%v)/%v?%v", ENV.DB_USERNAME, ENV.DB_PASSWORD, ENV.DB_URL, ENV.DB_DATABASE, "charset=utf8mb4&parseTime=True&loc=Local")
	db, err := gorm.Open(mysql.Open(dbConnectionStr), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	BD = db
}
