package mysql

import "go-blog/models"

func GetTags(pageNum int, pageSize int, maps interface{}) (tags []models.Tag) {
	tags = []models.Tag{}
	db.Where(maps).Offset(pageNum).Limit(pageSize).Find(&tags)
	return
}

func GetTagTotal(maps interface{}) (count int64) {
	db.Model(&models.Tag{}).Where(maps).Count(&count)
	return
}

func ExistTagByName(name string) bool {
	var tag models.Tag
	db.Select("id").Where("name = ?", name).First(&tag)
	if tag.ID > 0 {
		return true
	}
	return false
}

func AddTag(name string, state int, createdBy string) bool {
	db.Create(&models.Tag{
		Name:      name,
		State:     state,
		CreatedBy: createdBy,
	})
	return true
}
func ExistTagByID(id int) bool {

	var tag models.Tag
	db.Select("id").Where("id = ?", id).First(&tag)
	if tag.ID > 0 {
		return true
	}
	return false
}

func DeleteTag(id int) bool {
	db.Where("id = ?", id).Delete(&models.Tag{})
	return true
}

func EditTag(id int, data interface{}) bool {
	db.Model(&models.Tag{}).Where("id = ?", id).Updates(data)
	return true
}
