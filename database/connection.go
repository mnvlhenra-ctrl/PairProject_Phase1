package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func ConnectDatabase() {

	dsn := "root:Blackhell11!@tcp(localhost:3306)/Game_Store?parseTime=true"

	db, err := sql.Open("mysql", dsn)

	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()

	if err != nil {
		log.Fatal(err)
	}

	DB = db

	fmt.Println("===================================")
	fmt.Println(" Connected to Game Store Database")
	fmt.Println("===================================")

}
