package main

import (
	"fmt"
	"log"
	"time"

	"github.com/hpcloud/tail"
)

func main() {
	filename := "./mylog"
	config := tail.Config{
		ReOpen:    true,
		Follow:    true,
		Location:  &tail.SeekInfo{Offset: 0, Whence: 2},
		MustExist: false,
		Poll:      true,
	}

	t, err := tail.TailFile(filename, config)
	if err != nil {
		log.Fatalf("tail file: %s err: %v", filename, err)
		return
	}

	var msg *tail.Line

	var ok bool

	for {
		msg, ok = <-t.Lines
		if !ok {
			fmt.Printf("tail file close reopen, filename: %s", t.Filename)
			time.Sleep(time.Second)
			continue
		}
		fmt.Println("msg: ", msg.Text)
	}
}
