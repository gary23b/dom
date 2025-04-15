//go:build !(js && wasm)

package dom

var Window WindowI

func init() {
	Window = &window{
		ValueI: valueS{jsValue: "window"},
	}
}

var Doc DocumentI

func init() {
	window := &window{
		ValueI: valueS{jsValue: "window"},
	}
	Doc = window.Document()
}
