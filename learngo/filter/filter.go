package main

import "fmt"

func IntergerGenerator() chan int {
	ch := make(chan int)

	go func() {
		for i := 2; ; i++ {
			ch <- i
		}
	}()

	return ch
}

// Filter 输入整数队列numbers, 筛选出number的倍数, 保留不是number倍数的数
func Filter(numbers chan int, number int) chan int {
	remainNumbers := make(chan int)
	go func() {
		for {
			i := <-numbers

			if i%number != 0 {
				remainNumbers <- i
			}
		}
	}()

	return remainNumbers
}

func main() {
	const MAX = 100
	numbers := IntergerGenerator()
	number := <-numbers

	for number <= MAX {
		fmt.Printf("%d ", number)
		numbers = Filter(numbers, number)
		number = <-numbers
	}
}
