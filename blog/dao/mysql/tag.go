package mysql

import (
	"blog/models"
)

func GetTags(page, size int) ([]models.Tag, error) {
	var tags []models.Tag
	var err error

	err = db.Model(&models.Tag{}).Offset(page).Limit(size).Find(&tags).Error
	if err != nil {
		return []models.Tag{}, err
	}

	return tags, nil
}

func AddTag(user *models.Tag) error {

	return db.Model(&models.Tag{}).Create(&user).Error

}

func TagExistByID(id int) (models.Tag, error) {
	var tag models.Tag
	var err error

	err = db.Model(&models.Tag{}).Where("id = ?", id).First(&tag).Error

	if tag.ID > 0 {
		return tag, nil
	}

	return models.Tag{}, err
}

func TagExistByName(name string) (models.Tag, error) {
	var tag models.Tag
	var err error

	err = db.Model(&models.Tag{}).Where("name = ?", name).First(&tag).Error

	if tag.ID > 0 {
		return tag, nil
	}

	return models.Tag{}, err
}

func GetTag(id int) (models.Tag, error) {
	var tag models.Tag
	var err error

	err = db.Model(&models.Tag{}).Where("id = ?", id).First(&tag).Error

	if tag.ID > 0 {
		return tag, nil
	}

	return models.Tag{}, err
}

func EditTag(tag *models.Tag) error {

	// update 操作最好使用 map, 只更新所要求更新的字段.
	// 更详细的信息查看官方文档

	// 更新操作不涉及created_by, 如果尝试更新会被忽略
	var t = make(map[string]interface{})
	t["name"] = tag.Name
	t["modified_by"] = tag.ModifiedBy

	err := db.Model(&models.Tag{}).Where("id = ?", tag.ID).Updates(t).Error
	if err != nil {
		return err
	}

	return nil
}

func DeleteTag(id int) error {

	return db.Delete(&models.Tag{}, id).Error
}
