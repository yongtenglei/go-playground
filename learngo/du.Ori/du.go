package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"time"
)

func WalkDir(dir string, fileSizes chan<- int64) {
	for _, entry := range dirents(dir) {
		if entry.IsDir() {
			subdir := filepath.Join(dir, entry.Name())
			WalkDir(subdir, fileSizes)
		} else {
			fileSizes <- entry.Size()
		}
	}
}

func dirents(dir string) []os.FileInfo {
	entries, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}

	return entries

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

	start := time.Now()

	go func() {
		for _, root := range roots {
			WalkDir(root, fileSizes)
		}
		close(fileSizes)
	}()

	var nfiles, nbytes int64
loop:
	for {
		select {
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
