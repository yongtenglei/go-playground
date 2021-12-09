package main

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func Ding() {
	time.Sleep(time.Second * 3)
	fmt.Println("Ding")
}

func Dong() {
	time.Sleep(time.Second * 2)
	fmt.Println("Dong")
}

func DingDong(c *gin.Context) {
	Ding()
	Dong()
}

func TimerMiddleWares(c *gin.Context) {
	start := time.Now()
	c.Next()
	end := time.Now().Sub(start)
	fmt.Println("time elapsed ", end)
}

func TimerMiddleWares2() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Next()
		end := time.Now().Sub(start)
		fmt.Println("time elapsed ", end)

	}
}

func main() {
	r := gin.Default()

	//r.Use(TimerMiddleWares2())

	r.GET("/ding", TimerMiddleWares2(), DingDong)

	r.Run()
}
