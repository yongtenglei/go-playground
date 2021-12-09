package main

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Article struct {
	gorm.Model
	Name   string
	TagID  int
	Author Author
}

type Tag struct {
	gorm.Model
	Name     string
	Articles []Article
}

type Author struct {
	gorm.Model
	Name      string
	Role      string `gorm:"default:user"`
	ArticleID int
}

var (
	Db  *gorm.DB
	err error
)

func init() {
	dsn := "rey:1@tcp(121.40.151.71:3306)/one2many?charset=utf8mb4&parseTime=True&loc=Local"
	Db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}
}

func main() {
	Db.AutoMigrate(&Tag{}, &Article{}, &Author{})

	//a1 := Article{Model: gorm.Model{ID: 1}, Name: "Golang1"}
	//a2 := Article{Model: gorm.Model{ID: 2}, Name: "Golang2"}

	//t := Tag{Model: gorm.Model{ID: 1}, Name: "Golang", Articles: []Article{a1, a2}}

	//auth1 := Author{Model: gorm.Model{ID: 1}, Name: "admin1", Role: "admin"}
	//auth2 := Author{Model: gorm.Model{ID: 2}, Name: "user1"}

	//Db.Model(&a1).Association("Author").Append(&auth1)
	//Db.Model(&a2).Association("Author").Append(&auth2)

	var tag Tag
	Db.Preload("Articles", func(db *gorm.DB) *gorm.DB {
		return db.Joins("Author").Where("role = ?", "admin")
	}).Find(&tag)
	fmt.Println(tag)

}

//a1 := Article{Model: gorm.Model{ID: 1}, Name: "Golang1"}
//a2 := Article{Model: gorm.Model{ID: 2}, Name: "Golang2"}
//t := Tag{Model: gorm.Model{ID: 1}, Name: "Golang", Articles: []Article{a1, a2}}

//var tag Tag

//Db.Preload("Articles", func(db *gorm.DB) *gorm.DB {
//return db.Where("id = ?", "2")
//}).Find(&tag)

//fmt.Println(tag)

//Db.Model(&t).Association("Article").Append(&a2)

//var tag Tag
//Db.Preload("Article").Where("id = ?", 1).First(&tag)
//fmt.Printf("%#v", tag)
