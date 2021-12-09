package main

import (
	"fmt"
	"strconv"
	"time"
)

func Greeting() {
	index := 0
	for {
		index++
		fmt.Println("hello k8s" + strconv.Itoa(index))
		time.Sleep(time.Second)
	}
}

func main() {
	Greeting()
}
