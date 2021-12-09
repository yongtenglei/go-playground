package router

import (
	"blog/conf"
	v1 "blog/controllers/v1"

	"github.com/gin-gonic/gin"
)

func RouterSetup() *gin.Engine {

	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery())
	gin.SetMode(conf.Conf.AppConf.RunMode)

	apiV1 := r.Group("/api/v1")
	{
		// tags
		apiV1.GET("/tags", v1.GetTags)          // 获取多个tag
		apiV1.GET("/tags/:id", v1.GetTag)       // 获取特定id的tag
		apiV1.POST("/tags", v1.AddTag)          // 创建一个tag
		apiV1.PUT("/tags/:id", v1.EditTag)      // 修改tag
		apiV1.DELETE("/tags/:id", v1.DeleteTag) // 删除tag

		// articles
		apiV1.GET("/articles", v1.GetArticles)          // 获取多个tag
		apiV1.GET("/articles/:id", v1.GetArticle)       // 获取特定id的tag
		apiV1.POST("/articles", v1.AddArticle)          // 创建一个tag
		apiV1.PUT("/articles/:id", v1.EditArticle)      // 修改tag
		apiV1.DELETE("/articles/:id", v1.DeleteArticle) // 删除tag

	}
	return r
}
