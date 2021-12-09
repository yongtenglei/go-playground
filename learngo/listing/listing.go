package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	numberGoroutine = 4
	taskLoad        = 10
)

var wg sync.WaitGroup

func init() {
	rand.Seed(time.Now().Unix())
}

func main() {

	tasks := make(chan string, taskLoad)

	wg.Add(numberGoroutine)

	for i := 1; i <= numberGoroutine; i++ {
		go worker(tasks, i)
	}

	for i := 1; i <= taskLoad; i++ {
		tasks <- fmt.Sprintf("Task: %d", i)
	}

	close(tasks)

	wg.Wait()
}

func worker(tasks chan string, id int) {
	defer wg.Done()

	for {
		task, ok := <-tasks
		if !ok {
			fmt.Printf("Worker %d: Shunting Down\n", id)
			return
		}

		fmt.Printf("Worker %d Start his task %s\n", id, task)
		n := rand.Int63n(3)
		time.Sleep(time.Duration(n) * time.Second)

		fmt.Printf("Worker %d Finished his task %s\n", id, task)
	}

}
