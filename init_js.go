//go:build js && wasm

package dom

import (
	"syscall/js"
)

var Window WindowI
var Doc DocumentI
var Body ElementI

func init() {
	Window = &window{
		ValueI: NewValue(js.Global()),
	}

	Doc = Window.Document()

	Body = Doc.Body()
}
