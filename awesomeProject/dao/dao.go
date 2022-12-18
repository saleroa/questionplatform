package dao

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var DB *sql.DB

func InitDb() {
	db, err := sql.Open("mysql", "root:301044825@tcp(127.0.0.1:3306)/question")
	if err != nil {
		log.Printf("connect mysql error :%v", err)
	}

	DB = db
	fmt.Println(DB.Ping())
}
