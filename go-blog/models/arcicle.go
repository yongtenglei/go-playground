package models

import (
	"go-blog/conf"
	"time"

	"gorm.io/gorm"
)

type Article struct {
	gorm.Model
	TagID      int    `json:"tag_id" gorm:"index"`
	Tag        Tag    `json:"tag" gorm:"foreignkey:TagID"`
	Title      string `json:"title" binding:"required"`
	Name       string `json:"name" binding:"required"`
	Desc       string `json:"desc" binding:"required"`
	Content    string `json:"content" binding:"required"`
	CreatedBy  string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
	State      int    `json:"state"`
}

// TableName 会将 Tag 的表名重写为 `prefix+Tag`
func (Article) TableName() string {
	return conf.Conf.MysqlConf.TablePrefix + "article"
}

func (*Article) BeforeCreate(tx *gorm.DB) error {
	tx.Set("CreatedOn", time.Now().Unix())
	return nil
}

func (*Article) BeforeUpdate(tx *gorm.DB) (err error) {
	tx.Set("ModifiedOn", time.Now().Unix())
	return nil
}
