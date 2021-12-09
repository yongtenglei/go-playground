package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	Id   int
	name string
	age  int
}

var db *sql.DB

func initDB() (err error) {
	dns := "root:1@tcp(127.0.0.1:3306)/test"
	db, err = sql.Open("mysql", dns)
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	return nil
}

func queryDemo() {
	user := &User{}
	sql := "select * from user where id = ?"
	if err := db.QueryRow(sql, "1").Scan(&user.Id, &user.name, &user.age); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("query successfully, %#v\n", user)
}

func queryMutiDemo() {
	user := &User{}
	sql := "select * from user"
	rows, err := db.Query(sql)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		if err := rows.Scan(&user.Id, &user.name, &user.age); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("query successfully, %#v\n", user)
	}
}

func insertDemo() {
	sql := `insert into user (name, age) values ("hello", 1)`
	ret, err := db.Exec(sql)
	if err != nil {
		log.Fatal(err)
	}
	n, err := ret.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d affected\n", n)
}

func deleteDemo() {
	sql := `delete from user where id = ?`
	ret, err := db.Exec(sql, 4)
	if err != nil {
		log.Fatal(err)
	}
	n, err := ret.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d affected\n", n)
}

func updateDemo() {
	sql := `update user set age=30 where id = ?`
	ret, err := db.Exec(sql, 3)
	if err != nil {
		log.Fatal(err)
	}
	n, err := ret.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d affected\n", n)
}

func main() {
	err := initDB()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connect database successfully")

	//queryDemo()
	//insertDemo()
	//queryMutiDemo()
	//fmt.Println()
	//deleteDemo()
	//fmt.Println()
	//queryMutiDemo()

	updateDemo()
	queryMutiDemo()

}
