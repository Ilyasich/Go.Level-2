package main

import (
	"fmt"
	"time"
)

func main() {
	var ch = make(chan int)

	go func() {
		for val := range ch {
			fmt.Println(val)
		}
	}()

	go func() {
		for i := 1; i <= 1000; i++ {
			ch <- i

		}
	}()
	time.Sleep(1 * time.Second)
}
