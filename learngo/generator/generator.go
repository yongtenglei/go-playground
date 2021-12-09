package main

import "fmt"

func IntergerGenerator() chan int {
	ch := make(chan int)

	go func() {
		for i := 0; ; i++ {
			ch <- i
		}
	}()

	return ch
}

func main() {
	generator := IntergerGenerator()
	for i := 0; i < 100; i++ {
		fmt.Println("i = ", <-generator)

	}

}
