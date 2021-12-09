package main

import (
	"errors"
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
	fmt.Println("=========查询单条数据===========")
	var u1 User
	var u2 User
	var u3 User
	// 获取第一条记录（主键升序）
	db.First(&u1)
	// SELECT * FROM users ORDER BY id LIMIT 1;
	fmt.Println(u1)

	// 获取一条记录，没有指定排序字段
	db.Take(&u2)
	// SELECT * FROM users LIMIT 1;
	fmt.Println(u2)

	// 获取最后一条记录（主键降序）
	db.Last(&u3)
	// SELECT * FROM users ORDER BY id DESC LIMIT 1;
	fmt.Println(u3)

	result := db.First(&u3)
	fmt.Println(result.RowsAffected) // 返回找到的记录数
	fmt.Println(result.Error)        // returns error or nil

	// 检查 ErrRecordNotFound 错误
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		fmt.Println("record not found")
	} else {
		fmt.Println("record found successfully")
	}

	fmt.Println("=========查询多条数据===========")
	// 查询所有数据
	var users []User
	result = db.Find(&users)
	fmt.Println(result.RowsAffected) // 返回找到的记录数
	fmt.Println(users)

	result = db.Where("id in ?", []int64{1, 2}).Find(&users)
	fmt.Println(result.RowsAffected) // 返回找到的记录数
	fmt.Println(users)

}
