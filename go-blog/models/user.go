package models

import (
	"go-blog/conf"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Email    string `json:"email" binding:"omitempty"`
}

// TableName 会将 Tag 的表名重写为 `prefix+Tag`
func (User) TableName() string {
	return conf.Conf.MysqlConf.TablePrefix + "user"
}
