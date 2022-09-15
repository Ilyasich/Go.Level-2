package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	sig := make(chan int)

	sigs := make(chan os.Signal, 1)

	done := make(chan bool, 1)

	signal.Notify(sigs, syscall.SIGTERM)

	go func() {

		for i := 1; i <= 100; i++ {
			sig <- i
		}
	}()
	go func() {
		for val := range sig {
			fmt.Println(val)
		}

	}()
	go func() {
		sig := <-sigs
		fmt.Println()
		fmt.Print(sig)
		done <- true

	}()
	time.Sleep(1 * time.Second)

	fmt.Println("awaiting signal")
	<-done
	fmt.Println("exiting")
}
