package main

import "fmt"

func main() {
	msg := "!!!"
	defer fmt.Println(msg)
	msg = "world"
	defer fmt.Println(msg)
	fmt.Println("hello")
}
