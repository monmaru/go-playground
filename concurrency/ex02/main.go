package main

import (
	"fmt"
	"time"
)

func EsampleBufferedChannel() {
	ch := make(chan int, 5)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println("Writing", i)
			ch <- i
		}
		close(ch)
	}()

	time.Sleep(2 * time.Second)

	for {
		v, ok := <-ch
		if !ok {
			break
		}
		fmt.Println("Read", v)
	}
}

func main() {
	EsampleBufferedChannel()
}
