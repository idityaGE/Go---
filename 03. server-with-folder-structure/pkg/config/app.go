package config

import (
	 "gorm.io/driver/mysql"
	 "gorm.io/gorm"
)

var (
	db *gorm.DB
)

func Connect() {
	dsn := "gorm:gorm@tcp(localhost:9910)/gorm?charset=utf8&parseTime=True&loc=Local"
	d, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db = d
}

func GetDB() *gorm.DB {
	return db
}


