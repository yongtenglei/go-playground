package main

import (
	"fmt"
	"strings"
)

func main() {
	a := "localhost:9093, 127.0.0.1:20432"
	b := strings.Split(a, ",")
	fmt.Println(a)
	fmt.Println(b)
	fmt.Printf("v: %v\t t: %T", b, b)
}
