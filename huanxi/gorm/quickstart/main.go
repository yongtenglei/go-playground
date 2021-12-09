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

var db *gorm.DB
var err error

func init() {
	// 参考 https://github.com/go-sql-driver/mysql#dsn-data-source-name 获取详情

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer（日志输出的目标，前缀和日志包含的内容——译者注）
		logger.Config{
			SlowThreshold:             time.Second, // 慢 SQL 阈值
			LogLevel:                  logger.Info, // 日志级别
			IgnoreRecordNotFoundError: true,        // 忽略ErrRecordNotFound（记录未找到）错误
			Colorful:                  true,        // 启用彩色打印
		},
	)

	dsn := "root:1@tcp(127.0.0.1:3306)/huanxi?charset=utf8mb4&parseTime=True&loc=Local"
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		panic(err)
	}

	err = db.AutoMigrate(Tag{})
	if err != nil {
		panic("AutoMigrate Failed")
	}

}

func main() {
	fmt.Println("=========单行插入===========")
	fmt.Println("==========结构体============")
	t1 := Tag{Name: "t1"}
	result := db.Create(&t1)
	fmt.Println(result.Error)        // 返回 error
	fmt.Println(result.RowsAffected) // 返回插入记录的条数
	fmt.Println("==========MAP============")
	t2 := map[string]interface{}{"name": "t2"}
	result = db.Model(&Tag{}).Create(t2)
	fmt.Println(result.Error)        // 返回 error
	fmt.Println(result.RowsAffected) // 返回插入记录的条数

	fmt.Println("=========多行插入===========")
	fmt.Println("==========结构体============")
	t3s := []Tag{
		{Name: "t3"},
		{Name: "t4"},
		{Name: "t5"},
	}
	result = db.Create(&t3s)
	fmt.Println(result.Error)        // 返回 error
	fmt.Println(result.RowsAffected) // 返回插入记录的条数
	fmt.Println("==========MAP============")
	t4s := []map[string]interface{}{
		{"name": "t6"},
		{"name": "t7"},
		{"name": "t8"},
	}
	result = db.Model(&Tag{}).Create(&t4s)
	fmt.Println(result.Error)        // 返回 error
	fmt.Println(result.RowsAffected) // 返回插入记录的条数

	fmt.Println("=========多行插入InBatch===========")
	fmt.Println("==========结构体============")
	t5s := []Tag{
		{Name: "t9"},
		{Name: "t10"},
		{Name: "t11"},
	}
	result = db.CreateInBatches(&t5s, 2) // 一次2条数据进行插入
	fmt.Println(result.Error)            // 返回 error
	fmt.Println(result.RowsAffected)     // 返回插入记录的条数
	fmt.Println("==========MAP============")
	t6s := []map[string]interface{}{
		{"name": "t12"},
		{"name": "t13"},
		{"name": "t14"},
	}
	result = db.Model(&Tag{}).CreateInBatches(t6s, 3) // 一次三条数据插入
	fmt.Println(result.Error)                         // 返回 error
	fmt.Println(result.RowsAffected)                  // 返回插入记录的条数

}
