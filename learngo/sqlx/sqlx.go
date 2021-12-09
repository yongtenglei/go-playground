package main

import (
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var db *sqlx.DB

func Init() (err error) {
	dsn := "rey:1@tcp(121.40.151.71:3306)/test?charset=utf8mb4&parseTime=True"
	//dsn := "root:1@tcp(localhost:3306)/test?charset=utf8mb4&parseTime=True"

	if db, err = sqlx.Connect("mysql", dsn); err != nil {
		log.Println("Connect error")
		return
	}

	db.SetMaxOpenConns(299)
	db.SetMaxIdleConns(50)

	return
}

func Close() {
	_ = db.Close()

}

type Todo struct {
	Id int `db:"id"`
}

func main() {
	err := Init()
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	var todo []Todo
	sql := `select * from todos`
	rows, err := db.Query(sql)
	if err != nil {
		log.Fatal(err)
	}

	for rows.Next() {
		_ = rows.Scan(&todo)
		fmt.Println(todo)
	}
}
