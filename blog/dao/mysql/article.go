package mysql

import "blog/models"

func GetArticles(page, size int) ([]models.Article, error) {

	var articles []models.Article
	var err error

	err = db.Model(&models.Article{}).Preload("Tag").Offset(page).Limit(size).Find(&articles).Error
	if err != nil {
		return []models.Article{}, err
	}

	return articles, nil
}

func AddArticle(article *models.Article) error {

	return db.Model(&models.Article{}).Create(&article).Error

}

func ArticleExistByID(id int) (models.Article, error) {
	var article models.Article
	var err error

	err = db.Model(&models.Article{}).Where("id = ?", id).First(&article).Error

	if article.ID > 0 {
		return article, nil
	}

	return models.Article{}, err
}

func GetArticle(id int) (models.Article, error) {

	var article models.Article
	var err error

	err = db.Model(&models.Article{}).Preload("Tag").Where("id = ?", id).First(&article).Error
	if err != nil {
		return models.Article{}, err
	}

	return article, nil
}

func EditArticle(article *models.Article) error {

	// update 操作最好使用 map, 只更新所要求更新的字段.
	// 更详细的信息查看官方文档

	// 更新操作不涉及created_by, 如果尝试更新会被忽略
	var a = make(map[string]interface{})
	a["tag_id"] = article.TagID
	a["title"] = article.Title
	a["desc"] = article.Desc
	a["content"] = article.Content
	a["modified_by"] = article.ModifiedBy

	err := db.Model(&models.Article{}).Where("id = ?", article.ID).Updates(a).Error
	if err != nil {
		return err
	}

	return nil
}

func DeleteArticle(id int) error {

	return db.Delete(&models.Article{}, id).Error
}
