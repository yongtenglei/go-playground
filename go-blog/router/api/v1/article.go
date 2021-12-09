package v1

import (
	"fmt"
	"go-blog/dao/mysql"
	"go-blog/models"
	"go-blog/pkg/e"
	"go-blog/pkg/util"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

//获取单个文章
func GetArticle(c *gin.Context) {
	id, _ := util.Convertoi(c.Param("id"))

	var article models.Article
	article = mysql.GetArticle(id)

	code := e.SUCCESS
	if article.ID == 0 {
		code = e.ERROR_NOT_EXIST_ARTICLE

	}
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": article,
	})
}

//获取多个文章
func GetArticles(c *gin.Context) {
	page := util.GetPage(c)
	size, _ := util.Convertoi(c.Query("size"))

	var articles []models.Article
	articles = mysql.GetArticles(page, size)
	code := e.SUCCESS
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": articles,
	})

}

//新增文章
func AddArticle(c *gin.Context) {
	var article models.Article
	err := c.ShouldBindJSON(&article)
	if err != nil {
		fmt.Println(err)
		code := e.INVALID_PARAMS
		c.JSON(http.StatusOK, gin.H{
			"code": code,
			"msg":  e.GetMsg(code),
			"data": make(map[string]struct{}),
		})
		return
	}

	err = mysql.AddArticle(article)
	if err != nil {
		code := e.ERROR_NOT_EXIST_TAG
		c.JSON(http.StatusOK, gin.H{
			"code": code,
			"msg":  e.GetMsg(code),
			"data": make(map[string]struct{}),
		})
	}

	code := e.SUCCESS
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": make(map[string]struct{}),
	})
}

//修改文章
func EditArticle(c *gin.Context) {
	id, _ := util.Convertoi(c.Param("id"))

	if !mysql.ExistArticleByID(id) {
		code := e.ERROR_NOT_EXIST_ARTICLE
		c.JSON(http.StatusOK, gin.H{
			"code": code,
			"msg":  e.GetMsg(code),
			"data": make(map[string]struct{}),
		})

	}

	var article = make(map[string]interface{})

	c.ShouldBindJSON(&article)

	fmt.Printf("%#v", article)

	mysql.EditArticle(id, article)
	code := e.SUCCESS
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": make(map[string]struct{}),
	})

}

//删除文章
func DeleteArticle(c *gin.Context) {
	id, _ := util.Convertoi(c.Param("id"))

	if err := mysql.DeleteArticle(id); err != nil {
		code := e.ERROR_NOT_EXIST_ARTICLE

		c.JSON(http.StatusOK, gin.H{
			"code": code,
			"msg":  e.GetMsg(code),
			"data": make(map[string]struct{}),
		})
	}

	code := e.SUCCESS
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": make(map[string]struct{}),
	})
}

func GetArticlesOfTag(c *gin.Context) {
	tagId, _ := util.Convertoi(c.Param("id"))

	log.Println("1", tagId)
	articles, err := mysql.GetArticlesOfTag(tagId)
	log.Println(err)
	log.Println(articles)
	if err != nil {
		code := e.ERROR_NOT_EXIST_TAG
		c.JSON(http.StatusOK, gin.H{
			"code": code,
			"msg":  e.GetMsg(code),
			"data": make(map[string]struct{}),
		})
	}

	code := e.SUCCESS
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": articles,
	})

}
