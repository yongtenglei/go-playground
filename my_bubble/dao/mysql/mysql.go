package mysql

import (
	"fmt"
	"my_bubble/configs"
	"time"

	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB
var dbx *sqlx.DB

func Init(config *configs.MysqlConf) (err error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True",
		config.User,
		config.Password,
		config.Host,
		config.Port,
		config.DbName,
	)

	if db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{}); err != nil {
		zap.L().Error("Connect mysql failed", zap.Error(err))
		return
	}

	dbx, err := db.DB()
	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	dbx.SetMaxIdleConns(config.MaxIdleConns)

	// SetMaxOpenConns sets the maximum number of open connections to the database.
	dbx.SetMaxOpenConns(config.MaxOpenConns)

	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	dbx.SetConnMaxLifetime(time.Duration(config.MaxLifeTime))

	return
}
