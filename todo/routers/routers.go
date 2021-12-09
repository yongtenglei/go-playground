package routers

import (
	"todo/controller"

	"github.com/gin-gonic/gin"
)

func SetRouters() *gin.Engine {
	r := gin.Default()
	r.Static("/static", "static")
	r.LoadHTMLGlob("templates/*")
	r.GET("/", controller.HandleIndex)

	// v1
	v1Group := r.Group("v1")
	{

		// add
		v1Group.POST("/todo", controller.CreateATodo)

		//query
		v1Group.GET("/todo", controller.GetAllDodo)

		v1Group.GET("/todo/:id", func(c *gin.Context) {

		})

		// change
		v1Group.PUT("/todo/:id", controller.ModifyATodo)

		// delete
		v1Group.DELETE("/todo/:id", controller.DeleteATodo)

	}
	return r

}
