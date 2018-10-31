// +build wasm,!arm,!avr

package main

import (
	"bytes"
	"image"
	"image/color"
	"image/jpeg"
)

type Message int

var back []byte
var buf *bytes.Buffer

func main() {
	back = make([]byte, 0)
	buf = bytes.NewBuffer(back)
	img := image.NewRGBA(image.Rect(0, 0, 100, 50))
	img.Set(2, 3, color.RGBA{255, 0, 0, 255})
	jpeg.Encode(buf, img, nil)
	println(buf.Bytes())
}

type JsData struct {
	len   int
	start *byte
}

//go:export image
func imageData() JsData {
	return JsData{
		buf.Len(),
		&(back[0]),
	}
}
