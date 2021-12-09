package v1

import (
	"blog/conf"
	"blog/logic"
	"blog/models"
	"blog/pkg/e"
	"blog/pkg/util"

	"github.com/gin-gonic/gin"
)

//获取多个文章标签
func GetTags(c *gin.Context) {
	page := util.GetPage(c)
	size, err := util.StrToInt(c.DefaultQuery("size", util.IntToStr(conf.Conf.AppConf.PageSize)))
	if err != nil {
		code := e.INVALID_PARAMS
		util.ResposeWithError(c, code)
		return
	}

	var tags []models.Tag

	// 业务处理
	tags, err = logic.GetTags(page, size)
	if err != nil {
		util.ResposeWithError(c, e.ERROR)
		return
	}

	// 返回
	util.ResposeWithSuccessData(c, tags)
}

//新增文章标签
func AddTag(c *gin.Context) {
	var tag models.Tag

	// 获取参数, 校验参数
	if err := c.ShouldBindJSON(&tag); err != nil {
		util.ResposeWithError(c, e.INVALID_PARAMS)
		return
	}

	// 业务处理
	if err := logic.AddTag(&tag); err != nil {
		util.ResposeWithError(c, e.ERROR)
		return
	}

	// 返回
	util.ResposeWithSuccess(c)
}

//获取特定id的文章标签
func GetTag(c *gin.Context) {
	id, err := util.StrToInt(c.Param("id"))
	if err != nil {
		util.ResposeWithError(c, e.INVALID_PARAMS)
		return
	}

	var tag models.Tag
	tag, err = logic.GetTag(id)
	if err != nil {
		util.ResposeWithError(c, e.ERROR)
		return
	}

	util.ResposeWithSuccessData(c, tag)
}

//修改文章标签
func EditTag(c *gin.Context) {
	id, err := util.StrToInt(c.Param("id"))
	if err != nil {
		util.ResposeWithError(c, e.INVALID_PARAMS)
		return
	}

	var tag models.Tag
	err = c.ShouldBindJSON(&tag)
	if err != nil {
		util.ResposeWithError(c, e.INVALID_PARAMS)
		return
	}

	tag.ID = uint(id)

	err = logic.EditTag(&tag)
	if err != nil {
		util.ResposeWithError(c, e.ERROR)
		return
	}

	util.ResposeWithSuccess(c)

}

//删除文章标签
func DeleteTag(c *gin.Context) {
	id, err := util.StrToInt(c.Param("id"))
	if err != nil {
		util.ResposeWithError(c, e.INVALID_PARAMS)
		return
	}

	err = logic.DeleteTag(id)
	if err != nil {
		util.ResposeWithError(c, e.ERROR)
		return
	}

	util.ResposeWithSuccess(c)
}
