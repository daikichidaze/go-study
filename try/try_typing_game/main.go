//https://qiita.com/TsubasaSato/items/92f12af9d770bc93951f

package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"time"
)

var t int

func init() {
	flag.IntVar(&t, "t", 1, "Playing time (min)") // option
	flag.Parse()
}

func main() {
	var (
		files     = flag.Args()
		path, err = os.Executable()
		chRcv     = myinput(os.Stdin)
		tm        = time.After(time.Duration(t) * time.Minute)
		words     []string
		score     = 0
	)
	if err != nil {
		fmt.Fprintln(os.Stderr, "File read error", err)
		os.Exit((0))
	} else if len(files) != 1 {
		fmt.Fprintln(os.Stderr, "Select only one file")
		os.Exit(0)
	}

	path = filepath.Dir(path)
	words, err = getwords(filepath.Join(path, files[0]))
	if err != nil {
		fmt.Fprintln(os.Stderr, "Faild in file  reading", err)
		os.Exit(0)
	}

	fmt.Println("Start the typing game. Time limit:", t, "min. 1 point / 1 word.", len(words), "full points")

	// for i := true; i && score < len(words); {
Label:
	for score < len(words) {
		question := words[score]
		fmt.Print(question)
		fmt.Print(">")
		select {
		case x := <-chRcv:
			if question == x {
				score++
			}
		case <-tm:
			fmt.Println("Time out")
			break Label
			// i = false
		}
	}
	fmt.Println("Your score:", score, "/", len(words))
}
func myinput(r io.Reader) <-chan string {
	ch1 := make(chan string)
	go func() {
		s := bufio.NewScanner(r)
		for s.Scan() {
			ch1 <- s.Text()
		}
	}()
	return ch1
}
func getwords(txtPath string) ([]string, error) {
	var words []string
	sf, err := os.Open(txtPath)
	if err != nil {
		return nil, err
	} else {
		scanner := bufio.NewScanner(sf)
		for scanner.Scan() {
			words = append(words, scanner.Text())
		}
	}
	return words, err
}
