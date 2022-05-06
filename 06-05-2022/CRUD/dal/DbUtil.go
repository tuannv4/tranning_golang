package dal

import (
	"database/sql"
	"fmt"
	"time"
	_ "github.com/go-sql-driver/mysql"
)

const (
	username = "root"
	password = "0000"
	hostname = "127.0.0.1:3306"
	dbname   = "product"
)

var db *sql.DB

func InitializeMySQL() {
	dBConnection, err := sql.Open("mysql", dsn())
	if err != nil {
		fmt.Println("connect fail")
	}
	err = dBConnection.Ping()
	if err != nil {
		fmt.Println("Ping Flase")
	}
	db = dBConnection
	dBConnection.SetMaxOpenConns(10)
	dBConnection.SetMaxIdleConns(5)
	dBConnection.SetConnMaxLifetime(time.Second * 10)
}

// get connect is get MySQL Connection

func GetConnection() *sql.DB {
	if db == nil {
		InitializeMySQL()
	}
	return db
}

//CloseStmt after run stmt
func CloseStmt(stmt *sql.Stmt) {
	if stmt != nil {
		stmt.Close()
	}
}

//CloseRows when select
func CloseRows(rows *sql.Rows) {
	if rows != nil {
		rows.Close()
	}
}

// format data connect
func dsn() string {
	return fmt.Sprintf("%s:%s@tcp(%s)/%s", username, password, hostname, dbname)
}
