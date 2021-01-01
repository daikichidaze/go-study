package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	mymodule "work/try/try_image_conv/module"
)

func main() {
	flag.Parse()
	args := flag.Args()

	defaultExt := "png"

	var ext string
	if len(args) < 1 {
		ext = defaultExt
	} else {
		ext = args[0]
	}

	// Goファイルを探し出す
	err := filepath.Walk("./",
		func(path string, info os.FileInfo, err error) error {
			if filepath.Ext(path) == ".jpg" || filepath.Ext(path) == ".jpeg" {
				fmt.Println(path)

				outputFile := fmt.Sprintf("%s.%s", strings.Split(path, ".")[0], ext)
				mymodule.Imageconv(path, outputFile)
			}
			return nil
		})
	if err != nil {
		println(err)
	}

}
