package main

func swap(arg1 int, arg2 int) (int, int) {
	return arg2, arg1
}

func swap2(arg1 *int, arg2 *int) {
	v1, v2 := *arg1, *arg2

	*arg1 = v2
	*arg2 = v1
}

func main() {
	n, m := swap(10, 20)
	println(n, m)

	n2, m2 := 11, 22
	swap2(&n2, &m2)
	println(n2, m2)

}
