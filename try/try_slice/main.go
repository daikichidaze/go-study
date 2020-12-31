package main

import "fmt"

func main() {
	s := []int{19, 86, 1, 12}
	var sum int
	for _, v := range s {
		sum += v
	}
	fmt.Println(sum)
}
