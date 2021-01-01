package main

import (
	"bufio"
	"fmt"
	"os"
)

func catFile(path string, rowFlg bool, rowIdx *int) {
	data, _ := os.Open(path)
	defer data.Close()
	scanner := bufio.NewScanner(data)
	for scanner.Scan() {
		*rowIdx++
		var output string
		if rowFlg {
			output = fmt.Sprintf("%v: %v\n", *rowIdx, scanner.Text())
		} else {
			output = fmt.Sprintf("%v\n", scanner.Text())
		}
		fmt.Printf(output)
	}
}

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func main() {
	args := os.Args
	var rowIdxFlg bool
	var idx int

	for k, v := range args {
		if k == 0 {
			continue
		}
		if v == "-n" {
			rowIdxFlg = true
		}
		if fileExists(v) {
			catFile(v, rowIdxFlg, &idx)
		}
	}

}
