package models

import (
	"blog/conf"
	"time"

	"gorm.io/gorm"
)

type Tag struct {
	gorm.Model
	Name       string `json:"name"`
	CreatedBy  string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
}

// TableName 会将 Tag 的表名重写为 `prefix+Tag`
func (Tag) TableName() string {
	return conf.Conf.MysqlConf.TablePrefix + "tag"
}

func (tag *Tag) BeforeCreate(tx *gorm.DB) error {
	tx.Set("CreatedOn", time.Now().Unix())
	return nil
}

func (*Tag) BeforeUpdate(tx *gorm.DB) (err error) {
	tx.Set("ModifiedOn", time.Now().Unix())
	return nil
}
