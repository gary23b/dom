//go:build js && wasm

package dom

import "syscall/js"

var Window WindowI

func init() {
	Window = &window{
		ValueI: NewValue(js.Global()),
	}
}

var Doc DocumentI

func init() {
	window := &window{
		ValueI: NewValue(js.Global()),
	}
	Doc = window.Document()
}
