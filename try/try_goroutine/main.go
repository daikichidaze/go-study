package main

import (
	"fmt"
	"time"
)

func main() {
	defer fmt.Println("main done")
	go func() {
		defer fmt.Println("go routine 1 done")
		time.Sleep(5 * time.Second)
	}()
	go func() {
		defer fmt.Println("go routine 2 done")
		time.Sleep(3 * time.Second)
	}()
	time.Sleep(1 * time.Second)

}
