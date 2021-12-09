package main

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Article struct {
	gorm.Model
	Name  string
	TagID int
}

type Tag struct {
	gorm.Model
	Name    string
	Article Article
}

var (
	Db  *gorm.DB
	err error
)

func init() {
	dsn := "rey:1@tcp(121.40.151.71:3306)/one2one?charset=utf8mb4&parseTime=True&loc=Local"
	Db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}
}

func main() {
	Db.AutoMigrate(&Tag{}, &Article{})

	a := Article{Model: gorm.Model{ID: 1}, Name: "Practical Golang"}
	t := Tag{Model: gorm.Model{ID: 1}, Name: "Golang", Article: a}
	a2 := Article{Model: gorm.Model{ID: 2}, Name: "Leaning Golang"}

	//Db.Create(&a2)

	Db.Model(&t).Association("Article").Append(&a2)

	var tag Tag
	Db.Preload("Article").Where("id = ?", 1).First(&tag)
	fmt.Printf("%#v", tag)
}

//t2 := Tag{Model: gorm.Model{ID: 2}, Name: "Golang  for beginner"}

//Db.Model(&a).Association("Tag").Append(&t2)

//var art Article
//Db.Preload("Tag").Where("id = ?", 1).First(&art)
//fmt.Printf("%#v", art)
