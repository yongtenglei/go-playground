package v1

import (
	"blog/conf"
	"blog/logic"
	"blog/models"
	"blog/pkg/e"
	"blog/pkg/util"

	"github.com/gin-gonic/gin"
)

//获取多个文章
func GetArticles(c *gin.Context) {
	page := util.GetPage(c)
	size, err := util.StrToInt(c.DefaultQuery("size", util.IntToStr(conf.Conf.AppConf.PageSize)))
	if err != nil {
		code := e.INVALID_PARAMS
		util.ResposeWithError(c, code)
		return
	}

	var articles []models.Article

	// 业务处理
	articles, err = logic.GetArticles(page, size)
	if err != nil {
		util.ResposeWithError(c, e.ERROR)
		return
	}

	// 返回
	util.ResposeWithSuccessData(c, articles)
}

//新增文章
func AddArticle(c *gin.Context) {
	var article models.Article

	// 获取参数, 校验参数
	if err := c.ShouldBindJSON(&article); err != nil {
		util.ResposeWithError(c, e.INVALID_PARAMS)
		return
	}

	// 业务处理
	if err := logic.AddArticle(&article); err != nil {
		util.ResposeWithError(c, e.ERROR)
		return
	}

	// 返回
	util.ResposeWithSuccess(c)
}

//获取特定id的文章
func GetArticle(c *gin.Context) {
	id, err := util.StrToInt(c.Param("id"))
	if err != nil {
		util.ResposeWithError(c, e.INVALID_PARAMS)
		return
	}

	var tag models.Article
	tag, err = logic.GetArticle(id)
	if err != nil {
		util.ResposeWithError(c, e.ERROR)
		return
	}

	util.ResposeWithSuccessData(c, tag)
}

//修改文章
func EditArticle(c *gin.Context) {
	id, err := util.StrToInt(c.Param("id"))
	if err != nil {
		util.ResposeWithError(c, e.INVALID_PARAMS)
		return
	}

	var article models.Article
	err = c.ShouldBindJSON(&article)
	if err != nil {
		util.ResposeWithError(c, e.INVALID_PARAMS)
		return
	}

	article.ID = uint(id)

	err = logic.EditArticle(&article)
	if err != nil {
		util.ResposeWithError(c, e.ERROR)
		return
	}

	util.ResposeWithSuccess(c)

}

//删除文章
func DeleteArticle(c *gin.Context) {
	id, err := util.StrToInt(c.Param("id"))
	if err != nil {
		util.ResposeWithError(c, e.INVALID_PARAMS)
		return
	}

	err = logic.DeleteArticle(id)
	if err != nil {
		util.ResposeWithError(c, e.ERROR)
		return
	}

	util.ResposeWithSuccess(c)
}
