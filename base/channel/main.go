package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)

	go func() {
		ch <- 2
	}()

	go func() {
		time.Sleep(time.Second * 5)
		fmt.Println("finish")
		ch <- 1
	}()

	select {
	case <-ch:
		fmt.Println("main start")
	case <-time.After(3 * time.Second):
		fmt.Println("Timeout")
	}

	// 没有数据会一直阻塞
	<-ch
	println("done")
}
