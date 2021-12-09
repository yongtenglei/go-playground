package main

import (
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	url := "http://www.baidu.com"

	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	f, err := os.Create("file")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	dest := io.MultiWriter(os.Stdout, f)
	io.Copy(dest, resp.Body)
	defer resp.Body.Close()
}
