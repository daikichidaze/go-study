package imageconv

import (
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"os"
	"strings"
)

func Imageconv(argInput string, argOutput string) {
	f, err := os.Open(argInput) // 元画像読み込み
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

	fso, err := os.Create(argOutput) // 変換後画像作成
	if err != nil {
		fmt.Println("create:", err)
		return
	}
	defer fso.Close()

	slice := strings.Split(argOutput, ".")

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
