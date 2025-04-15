//go:build !(js && wasm)

package dom

var Window WindowI
var Doc DocumentI
var Body ElementI

func init() {
	Window = &window{
		ValueI: valueS{jsValue: "window"},
	}

	Doc = Window.Document()

	Body = Doc.Body()
}
