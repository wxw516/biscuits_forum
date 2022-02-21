package tool

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var Db *sql.DB

func GetDb() *sql.DB {
	return Db
}

func SqlEngine()  {
	connStr := "root:@tcp(localhost:3306)/biscuits_forum"
	db, err := sql.Open("mysql", connStr)
	if err != nil {
		log.Fatal(err.Error())
		return
	}
	Db = db

}
