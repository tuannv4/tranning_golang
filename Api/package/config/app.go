package config

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const (
	username = "root"
	password = "0000"
	hostname = "127.0.0.1:3306"
	dbname   = "product"
)

var db *gorm.DB

func dsn() string {
	return fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", username, password, hostname, dbname)
}

func Connect() {
	inforDB := dsn()
	dbCon, err := gorm.Open(mysql.Open(inforDB), &gorm.Config{})

	if err != nil {
		log.Println("error connect!")
	}
	db = dbCon
	log.Println("connect succesfully!")

}

func GetDB() *gorm.DB {
	return db
}
