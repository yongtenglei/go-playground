package mysql

import (
	"go-blog/models"
)

func ExistArticleByID(id int) bool {
	var article models.Article
	db.Select("id").Where("id = ?", id).First(&article)
	if article.ID > 0 {
		return true
	}
	return false
}

func GetArticleTotal(maps interface{}) (count int64) {
	db.Preload("Tag").Model(&models.Article{}).Where(maps).Count(&count)
	return
}

// 所有的文章
func GetArticles(pageNum int, pageSize int) (articles []models.Article) {
	db.Model(models.Article{}).Preload("Tag").Offset(pageNum).Limit(pageSize).Find(&articles)
	//db.Offset(pageNum).Limit(pageSize).Find(&articles)
	return
}

// 特定id的文章
func GetArticle(id int) (article models.Article) {
	db.Model(models.Article{}).Preload("Tag").Where("id = ?", id).First(&article)
	return
}

func EditArticle(id int, data interface{}) error {
	//switch data.(type) {
	//case map[string]interface{}:
	//return db.Model(&models.Article{}).Updates(data).Error
	//}

	return db.Model(&models.Article{}).Where("id = ?", id).Updates(data).Error
	//return db.Model(&models.Article{}).Where("id = ?", id).Updates(map[string]interface{}{
	//"title":       "art3edited",
	//"modified_by": "auth3",
	//}).Error

	//m, ok := data.(map[string]interface{})
	//if ok {
	//return db.Model(&models.Article{}).Where("id = ?", id).Updates(m).Error
	//}
	//return errors.New("Edit failed")
}

func AddArticle(article models.Article) error {
	return db.Create(&article).Error
}

func DeleteArticle(id int) error {
	return db.Delete(&models.Article{}, id).Error
}

func GetArticlesOfTag(id int) ([]models.Article, error) {
	//err = db.Model(models.Article{}).Preload("Tag", "id = ?", id).Find(&articles).Error
	var articles []models.Article
	err := db.Model(models.Article{}).Preload("Tag").Where("tag_id = ?", id).Find(&articles).Error
	//log.Println("2", id)
	//err := db.Model(models.Article{}).Joins("Tag", "id = ?", id).Find(&articles).Error

	return articles, err
}
