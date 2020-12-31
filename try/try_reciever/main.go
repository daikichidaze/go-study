package main

type MyInt int

func (n *MyInt) Inc() { *n++ }

func main() {
	var n MyInt
	println(n)
	n.Inc()
	n.Inc()
	println(n)

}
