package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func Middlewares1(c *gin.Context) {
	fmt.Println("=====start 1=====")
	c.Next()
	fmt.Println("=====end 1=====")

}

func Middlewares2(c *gin.Context) {
	fmt.Println("=====start 2=====")
	// Middlewares Chain exit

	_, ok := c.Get("id")
	if !ok {
		c.Abort()
	}
	c.Next()
	fmt.Println("=====end 2=====")

}
func Middlewares3(c *gin.Context) {
	fmt.Println("=====start 3=====")
	c.Next()
	fmt.Println("=====end 3=====")

}
func Middlewares4(c *gin.Context) {
	fmt.Println("=====start 4=====")
	c.Next()
	fmt.Println("=====end 4=====")

}
func Middlewares5(c *gin.Context) {
	fmt.Println("=====start 5=====")
	c.Next()
	fmt.Println("=====end 5=====")

}

func main() {
	r := gin.Default()

	r.GET("/ding", Middlewares1, Middlewares2, Middlewares3, Middlewares4, Middlewares5, Dong)

	r.Run()
}

func Dong(c *gin.Context) {
	fmt.Println("dong")
}
