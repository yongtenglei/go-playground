package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// r.LoadHTMLGlob("templates/**/*") 加载templates文件夹中的所有HTML文件
	r.LoadHTMLFiles("index.html")

	// load css file
	r.Static("/static", "./static")

	r.GET("/index", GetIndex)

	r.Run()
}

func GetIndex(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{
		"name": "rey",
	})
}
