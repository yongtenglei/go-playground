package routers

import (
	"app/controllers"
	"app/middlewares"

	"github.com/gin-gonic/gin"
)

func Setup(mode string) *gin.Engine {
	if mode == gin.ReleaseMode {
		gin.SetMode(gin.ReleaseMode)

	}
	r := gin.New()
	r.Use(middlewares.GinLogger(), middlewares.GinRecovery(true))

	v1 := r.Group("/api/v1")

	// 注册
	v1.POST("/signup", controllers.SignUpHandler)
	// 登录
	v1.POST("/login", controllers.LogInHandler)

	v1.Use(middlewares.JWTAuthMiddleware())

	{
		v1.GET("/community", controllers.CommunityHandler)
		v1.GET("/community/:id", controllers.CommunityDetailHandler)

		// Create A post
		v1.POST("/post", controllers.CreatePostHandler)
		// Get post detail with specific id
		v1.GET("/post/:id", controllers.PostDetileHandler)
		// Get All posts
		v1.GET("/post", controllers.PostListHandler)
		// Vote one post
		v1.POST("/post/vote", controllers.PostVoteHandler)

		// Get All posts by time/score order
		// Get All posts in specific community_id
		v1.GET("/post2", controllers.PostListHandler2)

	}

	return r
}
