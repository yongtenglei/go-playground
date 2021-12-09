package logic

import (
	"blog/dao/mysql"
	"blog/models"
)

func GetArticles(page, size int) ([]models.Article, error) {
	var articles []models.Article
	var err error

	articles, err = mysql.GetArticles(page, size)
	if err != nil {
		return articles, err
	}

	return articles, nil
}

func AddArticle(article *models.Article) error {
	err := mysql.AddArticle(article)
	if err != nil {
		return err
	}

	return nil
}

func GetArticle(id int) (models.Article, error) {
	var article models.Article
	var err error

	article, err = mysql.GetArticle(id)
	if err != nil {
		return models.Article{}, err
	}

	return article, nil
}

func EditArticle(article *models.Article) error {

	// 查看是否存在
	_, err := mysql.ArticleExistByID(int(article.ID))
	if err != nil {
		return err
	}

	return mysql.EditArticle(article)
}

func DeleteArticle(id int) error {
	// 查看是否存在
	_, err := mysql.ArticleExistByID(id)
	if err != nil {
		return err
	}

	return mysql.DeleteArticle(id)

}
