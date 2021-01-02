package main

import (
	"fmt"
	"time"
)

func tryGo(idx int) {
	var v int
	go func() {
		time.Sleep(1 * time.Second)
		v = 100
	}()
	go func() {
		time.Sleep(1 * time.Second)
		fmt.Println(idx, v)
	}()
	time.Sleep(2 * time.Second)

}

func main() {

	for i := 0; i < 10; i++ {
		tryGo(i)
	}

}
