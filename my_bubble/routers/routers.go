package routers

import (
	"my_bubble/controllers"
	"my_bubble/middlewares"

	"github.com/gin-gonic/gin"
)

func Setup(mode string) (r *gin.Engine) {
	if mode == gin.ReleaseMode {
		gin.SetMode(gin.ReleaseMode)
	}

	r = gin.New()
	r.Use(middlewares.GinLogger(), middlewares.GinRecovery(true))

	v1 := r.Group("/v1/todo")
	{
		v1.GET("/", controllers.TodoList)
		v1.POST("/", controllers.CreateTodo)
		v1.PUT("/:id", controllers.UpdateTodo)
		v1.DELETE("/:id", controllers.DeleteTodo)
	}

	return
}
