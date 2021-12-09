package main

import (
	"fmt"
	"sync"
	"time"
)

type NewsFactory struct {
	m    sync.Mutex
	news map[string]int
}

func (nf *NewsFactory) WNews() {
	nf.m.Lock()
	defer nf.m.Unlock()
	nf.news["big thing"]++
}

func (nf *NewsFactory) RNews() {
	nf.m.Lock()
	defer nf.m.Unlock()
	fmt.Println(nf.news["big thing"])
}

func main() {
	newsfactory := &NewsFactory{
		m: sync.Mutex{},
		news: map[string]int{
			"big thing": 0,
		},
	}

	for i := 0; i < 10; i++ {
		go func() {
			newsfactory.WNews()
		}()
	}

	time.Sleep(time.Second * 3)

	newsfactory.RNews()

}

//func main() {
//ch1 := make(chan int, 3)
//ch2 := make(chan int, 5)
//t := time.After(time.Second * 4)

//go func() {
//for i := 0; i < 3; i++ {
//ch1 <- i
//time.Sleep(time.Second)
//}
//}()

//go func() {
//for i := 0; i < 3; i++ {
//ch2 <- i
//time.Sleep(time.Second)
//}
//}()

//Loop:
//for {
//select {
//case v := <-ch1:
//fmt.Println("ch1 = ", v)

//case v := <-ch2:
//fmt.Println("ch2 = ", v)

//case <-t:
//fmt.Println("time out")
//close(ch1)
//close(ch2)
//break Loop
//default:
//}
//}

//}

//func PrintCow(wg sync.WaitGroup) {
//defer wg.Done()
//for i := 0; i < 10; i++ {
//fmt.Println("🦬")
//time.Sleep(time.Second)
//}
//}

//func PrintSheep(wg sync.WaitGroup) {
//defer wg.Done()
//for i := 0; i < 10; i++ {
//fmt.Println("🐏")
//time.Sleep(time.Second)
//}
//}

//var wg sync.WaitGroup

//func main() {
//wg.Add(2)

//go PrintCow(wg)
//go PrintSheep(wg)

//wg.Wait()

//}

//func main() {
//c := make(chan int, 10)
//c <- 1
//v, ok := <-c
//fmt.Println(v, ok)
//close(c)
//v, ok = <-c
//fmt.Println(v, ok)

//}

//func main() {
//c := make(chan int)
//go func() { c <- 1 }()
//fmt.Println(<-c)
//time.Sleep(time.Second)
//}

//func main() {
//for i := 0; i < 1000000; i++ {
//go func() {
//fmt.Println(i)
//}()
//}
//}

//func Show() string {
//return "test successfully"
//}

//func main() {
//Show()
//}

//func Test() {
//defer func() {
//err := recover()

//if err != nil {
//fmt.Println(err)
//}
//}()
//fmt.Println("==========1=========")
//fmt.Println("==========2=========")
//panic("==========3=========")
//fmt.Println("==========4=========")

//}

//func main() {

//fmt.Println("=======start======")
//Test()
//fmt.Println("=======end========")
//}

//func ShowType(i interface{}) {
//switch i.(type) {
//case int:
//fmt.Println("int type")
//case float32, float64:
//fmt.Println("float type")

//default:
//fmt.Println("not a number type")
//}

//}

//func main() {
//i := 1
//f32 := float32(1.1)
//f64 := float64(1.1)
//s := "non-number"
//ShowType(i)
//ShowType(f32)
//ShowType(f64)
//ShowType(s)

//}

//type DuckyType interface {
//Swimming()
//}

//type Ducky struct {
//Name string
//}

//func (d Ducky) Swimming() {
//fmt.Println(d.Name + "is swimming")
//}

//type DoggyType interface {
//Woofing()
//}

//type Doggy struct {
//Name string
//}

//func (d Doggy) Woofing() {
//fmt.Println(d.Name + "is woofing")
//}

//func (d Doggy) Swimming() {
//fmt.Println(d.Name + "is swimming")
//}

//type DockyType interface {
//DuckyType
//DoggyType
//}

//func main() {
//var ducky DuckyType
//ducky = Ducky{Name: "TangDuck"}
//ducky.Swimming()

//var doggy DoggyType
//doggy = Doggy{Name: "pluto"}
//doggy.Woofing()

//var docky DockyType
//docky = doggy.(DockyType)
//docky.Swimming()
//docky.Woofing()

//}
