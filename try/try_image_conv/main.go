package main

import (
	"flag"
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"os"
	"strings"
)

func main() {
	flag.Parse()
	args := flag.Args()
	defaultExt := "png"

	f, err := os.Open(args[0]) // 元画像読み込み
	if err != nil {
		fmt.Println("open:", err)
		return
	}
	defer f.Close()

	img, _, err := image.Decode(f) // 元画像デコード
	if err != nil {
		fmt.Println("decode:", err)
		return
	}

	var outputFile string
	if len(args) < 2 {
		outputFile = fmt.Sprintf("%s.%s", strings.Split(args[0], ".")[0], defaultExt)
	} else {
		outputFile = args[1]
	}

	fso, err := os.Create(outputFile) // 変換後画像作成
	if err != nil {
		fmt.Println("create:", err)
		return
	}
	defer fso.Close()

	slice := strings.Split(outputFile, ".")

	switch slice[len(slice)-1] { // 出力画像の拡張子によってエンコードを変える
	case "jpeg", "jpg":
		jpeg.Encode(fso, img, &jpeg.Options{})
	case "png":
		png.Encode(fso, img)
	case "gif":
		gif.Encode(fso, img, nil)
	default:
	}
}
