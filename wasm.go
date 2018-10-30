// +build wasm,!arm,!avr

package main

import (
	"bytes"
)

type Message int

const (
	Draw Message = iota
)

var back []byte
var page *bytes.Buffer

func main() {
	back = make([]byte, 256)
	page = bytes.NewBuffer(back)
	updatePage(0, 0)
}

// CommonWA: message
func _Cfunc_message(Message, *byte, int)

//go:export mouse
func updatePage(x, y int) {
	page.Truncate(0)
	page.WriteString("<div style=\"")
	page.WriteString("position:absolute;")
	page.WriteString("width:50px;")
	page.WriteString("height:50px;")
	page.WriteString("background:red;")
	page.WriteString("left:")
	page.WriteString("0")
	page.WriteString("px;top:")
	page.WriteString("0")
	page.WriteString("px;\"></div>")
	_Cfunc_message(Draw, &(back[0]), page.Len())
}
