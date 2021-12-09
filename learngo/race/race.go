package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	baton := make(chan int)

	wg.Add(1)

	go Player(baton)

	baton <- 1

	wg.Wait()
}

func Player(baton chan int) {
	var newRunner int

	runner := <-baton

	fmt.Printf("Runner %d Running with Baton\n", runner)

	if runner != 4 {
		newRunner = runner + 1
		fmt.Printf("Runner %d To the line\n", newRunner)
		go Player(baton)
	}

	time.Sleep(time.Millisecond)

	if runner == 4 {
		fmt.Printf("Runner %d finished, race over\n", runner)
		wg.Done()
		return
	}

	fmt.Printf("Runner %d exchange runner %d\n", runner, newRunner)

	baton <- newRunner
}
