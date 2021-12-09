package router

import (
	"go-blog/conf"
	"go-blog/middleware"
	v1 "go-blog/router/api/v1"

	"github.com/gin-gonic/gin"
)

func RouterSetup() *gin.Engine {

	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery())
	gin.SetMode(conf.Conf.AppConf.RunMode)

	r.POST("/signin", v1.SignIn)
	r.POST("/login", v1.LogIn)
	apiV1 := r.Group("/api/v1")
	apiV1.Use(middleware.JWTAuthMiddleware())
	{
		// tags
		apiV1.GET("/tags", v1.GetTags)
		apiV1.POST("/tags", v1.AddTag)
		apiV1.PUT("/tags/:id", v1.EditTag)
		apiV1.DELETE("/tags/:id", v1.DeleteTag)

		apiV1.GET("/tags/articles/:id", v1.GetArticlesOfTag)

		// articles
		apiV1.GET("/articles", v1.GetArticles)
		apiV1.GET("/articles/:id", v1.GetArticle)
		apiV1.POST("/articles", v1.AddArticle)
		apiV1.PUT("/articles/:id", v1.EditArticle)
		apiV1.DELETE("/articles/:id", v1.DeleteArticle)

		//

	}

	return r
}
