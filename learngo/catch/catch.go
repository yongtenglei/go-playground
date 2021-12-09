package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wg sync.WaitGroup

type Player struct {
	Name string
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	court := make(chan int)

	p1 := Player{Name: "rey"}
	p2 := Player{Name: "charlotte"}

	wg.Add(2)

	go p1.Play(court)
	go p2.Play(court)

	court <- 1

	wg.Wait()

}

func (p *Player) Play(court chan int) {
	defer wg.Done()

	for {
		ball, ok := <-court
		if !ok {
			fmt.Printf("Player %s win the game\n", p.Name)
			return
		}

		n := rand.Intn(100)
		if n%13 == 0 {
			fmt.Printf("Player %s missed the ball\n", p.Name)
			close(court)
			return
		}

		fmt.Printf("Player %s Hit the ball %d\n", p.Name, ball)
		ball++

		court <- ball
	}
}
