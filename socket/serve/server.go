package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"net"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	conns    []net.Conn
	username = "root"
	password = "0000"
	hostname = "127.0.0.1:3306"
	dbname   = "user"
)
var db *gorm.DB

type Infor struct {
	gorm.Model
	LatsName string `json:"last_name"`
	Age      int    `json:"age"`
}

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
func main() {
	server, err := net.Listen("tcp", ":3000")
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := server.Accept()
		if err != nil {
			log.Fatal(err)
		}
		conns = append(conns, conn)
		Connect()
		db.AutoMigrate(&Infor{})
		for {
			reader := bufio.NewReader(conn)
			msg, err := reader.ReadString('\n')
			if err != nil {
				break
			}
			data := Infor{}
			json.Unmarshal([]byte(msg), &data)
			fmt.Println("data Infor: ", data)
			db.Create(&data)
		}
	}

}
