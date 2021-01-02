package main

import (
	"fmt"
	"time"
)

func channel() {
	ch1 := make(chan int)
	ch2 := make(chan string)
	go func() {
		ch1 <- 100
	}()

	go func() {
		ch2 <- "hi"
	}()

	var v1 int
	select {
	case v1 = <-ch1:
		fmt.Println(v1)
	case v2 := <-ch2:
		fmt.Println(v1, v2)
	}
	time.Sleep(1 * time.Second)

}
func main() {
	for i := 0; i < 10; i++ {
		channel()
	}
}
