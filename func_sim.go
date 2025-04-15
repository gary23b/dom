//go:build !(js && wasm)

package dom

type funcS struct {
	Func string
}

var _ FuncI = funcS{}

func (s funcS) Release() {

}

func NewFuncForJavascript(fn func(this ValueI, args []ValueI) any) funcS {
	ret := funcS{}

	return ret
}
