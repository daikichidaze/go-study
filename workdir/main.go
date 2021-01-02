package main

func main() {
	h := 2.
	b := 0.9
	w := 1.5

	var cnt int
	if b < 1 && h > w {
		for h > w {
			println(h)
			cnt += 2
			h *= b
		}
		cnt--

	} else {
		cnt = -1
	}
	println(cnt)

}
