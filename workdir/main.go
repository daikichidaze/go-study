package main

import "fmt"

func main() {
	/*
		n := 100 + 200
		m := n + 10
		msg := "huga" + "hoge"

		println(msg, m < 1000)
	*/
	/*	for i, v := range []int{1, 2, 4} {
			println("No", i, v)
		}
	*/
	//	arr := [...]int{1, 4, 5, 66, 73, 4}
	/*
		m := map[string]string{"x": "10", "y": "11"}
		m["x"] = "test"
		x, ok := m["x"]
		println(x, ok)
	*/
	/*var n, m int
	n = 10
	m = n
	n += 2
	println(n, m)
	*/
	p := struct {
		age  int
		name string
	}{age: 10, name: "Gopher"}

	p2 := &p // コピー
	p2.age = 20

	fmt.Println(p)
	fmt.Println(p2)
}
