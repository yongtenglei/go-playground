package main

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type User struct {
	Id   int `db:"id"`
	name string
	age  int
}

var db *sqlx.DB

func init() {
	dns := "root:1@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True"
	db = sqlx.MustConnect("mysql", dns)
}

func main() {
	fmt.Println("Connect database successfully")

	user := &User{
		name: "syne",
		age:  10,
	}
	sql := `insert into user (name, age) values (:name, :age)`
	//db.NamedExec(sql, map[string]interface{}{
	//"name": "lili",
	//"age":  25,
	//})
	db.NamedExec(sql, &user)

}
