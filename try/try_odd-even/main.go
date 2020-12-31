package main

func ifVer(arg int) {
	if arg%2 == 1 {
		println("iOdd  - ", arg)
	} else if arg%2 == 0 {
		println("iEven - ", arg)
	}

}
func switchVer(arg int) {
	switch {
	case arg%2 == 1:
		println("sOdd  - ", arg)
	case arg%2 == 0:
		println("sEven - ", arg)
	}
}

func main() {
	for i := 0; i < 100; i++ {
		switchVer(i + 1)
	}
}
