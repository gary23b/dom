//go:build js && wasm

package dom

import "syscall/js"

type funcS struct {
	js.Func
}

var _ FuncI = funcS{}

// Release frees up resources allocated for the function.
// The function must not be invoked after calling Release.
// It is allowed to call Release while the function is still running.
func (s funcS) Release() {
	s.Func.Release()
}

// FuncOf returns a function to be used by JavaScript.
//
// The Go function fn is called with the value of JavaScript's "this" keyword and the
// arguments of the invocation. The return value of the invocation is
// the result of the Go function mapped back to JavaScript according to ValueOf.
//
// Invoking the wrapped Go function from JavaScript will
// pause the event loop and spawn a new goroutine.
// Other wrapped functions which are triggered during a call from Go to JavaScript
// get executed on the same goroutine.
//
// As a consequence, if one wrapped function blocks, JavaScript's event loop
// is blocked until that function returns. Hence, calling any async JavaScript
// API, which requires the event loop, like fetch (http.Client), will cause an
// immediate deadlock. Therefore a blocking function should explicitly start a
// new goroutine.
//
// Func.Release must be called to free up resources when the function will not be invoked any more.
func NewFuncForJavascript(fn func(this ValueI, args []ValueI) any) funcS {
	wrapper := js.FuncOf(func(this js.Value, args []js.Value) any {
		thisConverted := NewValue(this)
		argsConverted := make([]ValueI, 0, len(args))
		for _, arg := range args {
			argsConverted = append(argsConverted, NewValue(arg))
		}

		result := fn(thisConverted, argsConverted)
		return ValueOf(result)
	})

	ret := funcS{
		Func: wrapper,
	}

	return ret
}
