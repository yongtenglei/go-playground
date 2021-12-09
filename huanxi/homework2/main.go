package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	wg.Add(1)
	go func() {
		defer wg.Done()
	loop:
		for {
			select {
			case <-ctx.Done():
				fmt.Println("timeout")
				break loop
			default:
				fmt.Println("still working")
				time.Sleep(time.Second)
			}
		}
	}()

	wg.Wait()
}

//func main() {
//cond := sync.NewCond(new(sync.Mutex))
//condition := 0

//// 消费者
//go func() {
//for {
//// 消费者开始消费时，锁住
//cond.L.Lock()
//// 如果没有可消费的值，则等待
//for condition == 0 {
//cond.Wait()
//}
//// 消费
//condition--
//fmt.Printf("Consumer: %d\n", condition)

//// 唤醒一个生产者
//cond.Signal()
//// 解锁
//cond.L.Unlock()
//}
//}()

//// 生产者
//for {
//// 生产者开始生产
//cond.L.Lock()

//// 当生产太多时，等待消费者消费
//for condition == 100 {
//cond.Wait()
//}
//// 生产
//condition++
//fmt.Printf("Producer: %d\n", condition)

//// 通知消费者可以开始消费了
//cond.Signal()
//// 解锁
//cond.L.Unlock()
//}
//}

//func main() {
//c := make(chan int, 2)
//c <- 1
//c <- 2
//close(c)
//fmt.Println(<-c)
//fmt.Println(<-c)
//fmt.Println(<-c)
//fmt.Println(<-c)
//fmt.Println(<-c)
//}

//func DeferSequence() {
//defer fmt.Println("==========1==========")
//defer fmt.Println("==========2==========")
//defer fmt.Println("==========3==========")
//defer fmt.Println("==========4==========")
//defer fmt.Println("==========5==========")
//fmt.Println("==========6==========")
//}

//func main() {
//DeferSequence()
//}

//func MutiSummation(nums ...int) (sum int) {
//for _, v := range nums {
//sum += v
//}
//return
//}

//func main() {
//sum := MutiSummation(1, 2, 3)
//fmt.Println(sum)
//// output: 6
//}

//func makeSuffix(suffix string) func(str string) string {
//return func(str string) string {
//if !strings.HasSuffix(str, suffix) {
//return str + suffix
//}
//return str
//}
//}

//func main() {
//checkSuffix := makeSuffix(".txt")
//file := checkSuffix("rey")
//fmt.Println(file)
//// output: rey.txt
//}

//func Add1(s *[5]int) {
//for i := 0; i < len(s); i++ {
//s[i] += 1
//}

//}
//func main() {
//s := [5]int{1, 2, 3, 4, 5}
//fmt.Println(s)
//Add1(&s)
//fmt.Println(s)

//}
