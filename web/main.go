package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func fun1(c *gin.Context) {
	fmt.Println("func1:1")
	c.Next()

	fmt.Println("func1:2")
}
func fun2(c *gin.Context) {
	fmt.Println("func2:1")
	c.Next()

	fmt.Println("func2:2")
}
func fun3(c *gin.Context) {
	fmt.Println("func3:1")
	c.Next()

	fmt.Println("func3:2")
}
func fun4(c *gin.Context) {
	fmt.Println("func4:1")
	c.Next()

	fmt.Println("func4:2")
}
func fun5(c *gin.Context) {
	fmt.Println("func5:1")
	c.Next()

	fmt.Println("func5:2")
}

func main() {
	r := gin.Default()

	shopGroup := r.Group("shop/", fun1, fun2)
	{
		shopGroup.Use(fun3, fun4)
		shopGroup.GET("hello/", fun5, func(c *gin.Context) {
			c.String(http.StatusOK, "hello")
		})
	}

	r.Run(":9020")
}
