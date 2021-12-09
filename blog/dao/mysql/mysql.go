package mysql

import (
	"blog/conf"
	"blog/models"
	"fmt"
	"log"
	"time"

	"github.com/jmoiron/sqlx"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB
var dbx *sqlx.DB

func Init(config *conf.MysqlConf) (err error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True",
		config.User,
		config.Password,
		config.Host,
		config.DbName,
	)

	if db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{}); err != nil {
		log.Println("connect to mysql failed")
		return
	}

	dbx, err := db.DB()
	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	dbx.SetMaxIdleConns(config.MaxIdleConns)

	// SetMaxOpenConns sets the maximum number of open connections to the database.
	dbx.SetMaxOpenConns(config.MaxOpenConns)

	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	dbx.SetConnMaxLifetime(time.Duration(config.MaxLifeTime))

	db.AutoMigrate(models.Tag{}, models.Article{})

	return
}
