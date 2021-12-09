package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"sync"
	"time"
)

func WalkDir(dir string, fileSizes chan<- int64, wg *sync.WaitGroup) {
	defer wg.Done()

	if cancelled() {
		return
	}

	for _, entry := range dirents(dir) {
		if entry.IsDir() {
			wg.Add(1)
			subdir := filepath.Join(dir, entry.Name())
			go WalkDir(subdir, fileSizes, wg)
		} else {
			fileSizes <- entry.Size()
		}
	}
}

var sema = make(chan struct{}, 20)

func dirents(dir string) []os.FileInfo {
	select {
	case sema <- struct{}{}:
	case <-done:
		return nil
	}
	defer func() { <-sema }()

	entries, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}

	return entries

}

var done = make(chan struct{})

func cancelled() bool {
	select {
	case <-done:
		return true
	default:
		return false
	}
}

var verbose = flag.Bool("v", false, "show verbose progress messages")

func main() {
	flag.Parse()
	var tiker <-chan time.Time
	if *verbose {
		tiker = time.Tick(time.Millisecond)
	}

	roots := flag.Args()
	if len(roots) == 0 {
		roots = []string{"."}
	}

	fileSizes := make(chan int64)
	var wg sync.WaitGroup

	start := time.Now()

	go func() {
		os.Stdin.Read(make([]byte, 1))
		close(done)
	}()

	for _, root := range roots {
		wg.Add(1)
		go WalkDir(root, fileSizes, &wg)
	}

	go func() {
		wg.Wait()
		close(fileSizes)
	}()

	var nfiles, nbytes int64
loop:
	for {
		select {
		case <-done:
			for range fileSizes {
				//
			}
			return
		case size, ok := <-fileSizes:
			if !ok {
				break loop
			}
			nfiles++
			nbytes += size
		case <-tiker:
			fmt.Printf("%d files\t%.1f GB\n", nfiles, float64(nbytes)/1e9)
		}
	}
	fmt.Printf("%d files\t%.1f GB\n", nfiles, float64(nbytes)/1e9)
	fmt.Println(time.Since(start))
}
