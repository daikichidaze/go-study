package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	resp, err := http.Get("http://localhost:8080/?p=test")
	if err != nil {
		fmt.Println("error!")
		os.Exit(0)
	}
	defer resp.Body.Close()

	fmt.Println(resp.Body)

}
