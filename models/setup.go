package models

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

const DSN = "root:root123qq@tcp(sdqh-database.ckn7e9y6eoez.us-east-1.rds.amazonaws.com:3306)/sdqh?charset=utf8&parseTime=True&loc=Local"

func ConnectDatabase() {
	var err error
	DB, err = gorm.Open(mysql.Open(DSN), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	DB.AutoMigrate(&Users{}, &Admin{}, &Transaksi{}, &Waqaf_Sedeqah{})

}
