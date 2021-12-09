package main

import (
	"fmt"
	"math/rand"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	pingPongGroup := r.Group("/pingpong")
	{
		pingPongGroup.GET("/ping", Pong)
	}

	dingdongGroup := r.Group("/dingdong")
	{
		dingdongGroup.GET("/ding", Dong)

		// : 后的参数
		dingdongGroup.GET("/ding/:times", Dong2)
		dingdongGroup.GET("/ding/:times/:name", Dong3)

		// query 参数
		dingdongGroup.GET("/ding/query", Dong4)

		// * 参数
		dingdongGroup.POST("/file/*all", File)
		// post 表单数据
		dingdongGroup.POST("/ding/add", Dong5)
	}

	r.Run()

}

func Pong(c *gin.Context) {
	c.String(http.StatusOK, "pong\n")
}

func Dong(c *gin.Context) {
	c.String(http.StatusOK, "dong\n")
}

func Dong2(c *gin.Context) {
	times := c.Param("times")

	if times == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "bad request",
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"msg": fmt.Sprintf("Ding %s Dong", times),
	})
}

func Dong3(c *gin.Context) {
	var dingdong DingDong
	if err := c.ShouldBindUri(&dingdong); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "bad request",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"msg": fmt.Sprintf("Ding %d-%s Dong", dingdong.Times, dingdong.Name),
	})
}

func Dong4(c *gin.Context) {
	times := c.DefaultQuery("times", "1")

	c.JSON(http.StatusOK, gin.H{
		"msg": fmt.Sprintf("Ding %s Dong", times),
	})
}

func Dong5(c *gin.Context) {
	id := rand.Intn(10000)
	name := c.DefaultPostForm("name", "defaultName")

	c.JSON(http.StatusOK, gin.H{
		"id":   id,
		"name": name,
	})
}

func File(c *gin.Context) {
	all := c.Param("all")
	c.JSON(http.StatusOK, gin.H{
		"msg": all,
	})
}

type DingDong struct {
	Times int    `uri:"times" binding:"required"`
	Name  string `uri:"name" binding:"required"`
}
