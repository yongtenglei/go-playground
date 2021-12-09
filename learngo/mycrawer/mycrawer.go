package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
)

func Get(url string) (result string, err error) {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	buf := make([]byte, 1024*4)
	for {
		n, err := resp.Body.Read(buf)
		if err != nil {
			if err == io.EOF {
				fmt.Println("文件读取完毕")
				break
			} else {
				fmt.Println("读取文件错误")
				break
			}
		}
		result += string((buf[:n]))
	}
	return
}

func Crawer(i int, page chan<- int) {
	url := "https://github.com/search?q=go&type=Repositories&p=" + strconv.Itoa(i)
	result, err := Get(url)
	if err != nil {
		log.Fatal(err)
	}

	filename := "page" + strconv.Itoa(i) + ".html"
	f, err := os.Create(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	_, err = f.WriteString(result)
	if err != nil {
		log.Fatal(err)
	}

	page <- i
}

func Run(start, end int) {
	fmt.Printf("正在为你爬取第%d页到%d页的内容\n", start, end)
	page := make(chan int)

	for i := start; i <= end; i++ {
		go Crawer(i, page)
	}

	for i := start; i <= end; i++ {
		fmt.Printf("第%d页爬取完成\n", <-page)
	}

}
func main() {
	Run(3, 4)
}
