package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// 現在時刻
	t := time.Now().UnixNano()
	rand.Seed(t)
	n := rand.Intn(6) // 0-5

	switch n + 1 {
	case 6:
		fmt.Println("大吉")
	case 5, 4:
		fmt.Println("中吉")
	case 3, 2:
		fmt.Println("小吉")
	case 1:
		fmt.Println("凶")
	}
}
