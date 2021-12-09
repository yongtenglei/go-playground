package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Tag struct {
	gorm.Model
	Name string `json:"name"`
}

/* 等价与
type Tag struct {
	ID        uint `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt DeletedAt `gorm:"index"`
	Name string `json:"name"`
}*/

// TableName 会将 Tag 的表名重写为 `my_tag`
func (Tag) TableName() string {
	return "my_tag"
}

type User struct {
	gorm.Model
	Name     string `json:"name"`
	Password string `json:"password"`
}

// TableName 会将 Tag 的表名重写为 `my_user`
func (User) TableName() string {
	return "my_user"
}

var db *gorm.DB
var err error

func init() {
	// 参考 https://github.com/go-sql-driver/mysql#dsn-data-source-name 获取详情

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer（日志输出的目标，前缀和日志包含的内容——译者注）
		logger.Config{
			SlowThreshold:             time.Second,  // 慢 SQL 阈值
			LogLevel:                  logger.Error, // 日志级别
			IgnoreRecordNotFoundError: true,         // 忽略ErrRecordNotFound（记录未找到）错误
			Colorful:                  true,         // 启用彩色打印
		},
	)

	dsn := "root:1@tcp(127.0.0.1:3306)/huanxi?charset=utf8mb4&parseTime=True&loc=Local"
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		panic(err)
	}

	err = db.AutoMigrate(Tag{}, User{})
	if err != nil {
		panic("AutoMigrate Failed")
	}

}

func main() {
	// 创建User
	//u1 := User{Name: "u1", Password: "passwd1"}
	//u2 := User{Name: "u2", Password: "passwd2"}
	//us := []User{u1, u2}
	//db.Create(&us)

	fmt.Println("=========Save===========")
	// Save 会保存所有的字段，即使字段是零值
	var t1 Tag
	db.First(&t1)
	fmt.Println(t1)

	t1.Name = "t1 modified"
	db.Save(&t1)
	fmt.Println(t1)

	fmt.Println("=========Update 更新单列===========")
	var t2 Tag
	db.First(&t2, 2) // id = 2
	fmt.Println(t2)

	db.Model(&Tag{}).Where("id = ?", 2).Update("name", "t2 modified")
	db.First(&t2, 2) // id = 2
	fmt.Println(t2)

	fmt.Println("=========Updates 更新多列===========")
	// Updates 方法支持 struct 和 map[string]interface{} 参数。
	// 当使用 struct 更新时，默认情况下，GORM 只会更新非零值的字段

	fmt.Println("=========结构体===========")
	var u1 User
	db.First(&u1)
	fmt.Println(u1)

	db.Model(&User{}).Where("id = ?", 1).Updates(User{Name: "u1 modified", Password: ""})
	db.First(&u1, 1)
	fmt.Println(u1)

	fmt.Println("=========MAP===========")
	var u2 User
	db.First(&u2, 2)
	fmt.Println(u2)

	db.Model(&User{}).Where("id = ?", 2).Updates(map[string]interface{}{"name": "u2 modified", "password": ""})
	db.First(&u2, 2)
	fmt.Println(u2)
}
