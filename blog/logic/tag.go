package logic

import (
	"blog/dao/mysql"
	"blog/models"
	"errors"
)

func GetTags(page, size int) ([]models.Tag, error) {
	var tags []models.Tag
	var err error

	tags, err = mysql.GetTags(page, size)
	if err != nil {
		return tags, err
	}

	return tags, nil
}

func AddTag(tag *models.Tag) error {
	_, err := mysql.TagExistByName(tag.Name)
	if err == nil {
		return errors.New("tag already exist")
	}

	err = mysql.AddTag(tag)
	if err != nil {
		return err
	}

	return nil
}

func GetTag(id int) (models.Tag, error) {
	var tag models.Tag
	var err error

	tag, err = mysql.GetTag(id)
	if err != nil {
		return models.Tag{}, err
	}

	return tag, nil
}

func EditTag(tag *models.Tag) error {

	// 查看是否存在
	_, err := mysql.TagExistByID(int(tag.ID))
	if err != nil {
		return err
	}

	return mysql.EditTag(tag)
}

func DeleteTag(id int) error {
	// 查看是否存在
	_, err := mysql.TagExistByID(id)
	if err != nil {
		return err
	}

	return mysql.DeleteTag(id)

}
