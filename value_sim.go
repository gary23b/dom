//go:build !(js && wasm)

package dom

var (
	// null   = "null"
	// object = "object"
	array = "array"
)

type valueS struct {
	jsValue string
}

var _ ValueI = valueS{}

// Equal reports whether v and w are equal according to JavaScript's === operator.
func (s valueS) Equal(w ValueI) bool {
	other := w.(valueS)
	return s.jsValue == other.jsValue
}

// IsUndefined reports whether v is the JavaScript value "undefined".
func (s valueS) IsUndefined() bool {
	return false
}

// IsNull reports whether v is the JavaScript value "null".
func (s valueS) IsNull() bool {
	return false
}

// IsNaN reports whether v is the JavaScript value "NaN".
func (s valueS) IsNaN() bool {
	return false
}

// Type returns the JavaScript type of the value v. It is similar to JavaScript's typeof operator,
// except that it returns TypeNull instead of TypeObject for null.
func (s valueS) Type() Type {
	return TypeNull
}

// Get returns the JavaScript property p of value v.
// It panics if v is not a JavaScript object.
func (s valueS) Get(p string) ValueI {
	return valueS{}
}

// Set sets the JavaScript property p of value v to ValueOf(x).
// It panics if v is not a JavaScript object.
func (s valueS) Set(p string, x any) {

}

// Delete deletes the JavaScript property p of value v.
// It panics if v is not a JavaScript object.
func (s valueS) Delete(p string) {

}

// Index returns JavaScript index i of value v.
// It panics if v is not a JavaScript object.
func (s valueS) Index(i int) ValueI {
	return valueS{}
}

// SetIndex sets the JavaScript index i of value v to ValueOf(x).
// It panics if v is not a JavaScript object.
func (s valueS) SetIndex(i int, x any) {
}

// Length returns the JavaScript property "length" of v.
// It panics if v is not a JavaScript object.
func (s valueS) Length() int {
	return 0
}

// Call does a JavaScript call to the method m of value v with the given arguments.
// It panics if v has no method m.
// The arguments get mapped to JavaScript values according to the ValueOf function.
func (s valueS) Call(m string, args ...any) ValueI {
	return valueS{}
}

// Invoke does a JavaScript call of the value v with the given arguments.
// It panics if v is not a JavaScript function.
// The arguments get mapped to JavaScript values according to the ValueOf function.
func (s valueS) Invoke(args ...any) ValueI {
	return valueS{}
}

// New uses JavaScript's "new" operator with value v as constructor and the given arguments.
// It panics if v is not a JavaScript function.
// The arguments get mapped to JavaScript values according to the ValueOf function.
func (s valueS) New(args ...any) ValueI {
	return valueS{}
}

func (s valueS) Float() float64 {
	return 0
}

// Int returns the value v truncated to an int.
// It panics if v is not a JavaScript number.
func (s valueS) Int() int {
	return 0
}

// Bool returns the value v as a bool.
// It panics if v is not a JavaScript boolean.
func (s valueS) Bool() bool {
	return false
}

// Truthy returns the JavaScript "truthiness" of the value v. In JavaScript,
// false, 0, "", null, undefined, and NaN are "falsy", and everything else is
// "truthy". See https://developer.mozilla.org/en-US/docs/Glossary/Truthy.
func (s valueS) Truthy() bool {
	return false
}

// String returns the value v as a string.
// String is a special case because of Go's String method convention. Unlike the other getters,
// it does not panic if v's Type is not TypeString. Instead, it returns a string of the form "<T>"
// or "<T: V>" where T is v's type and V is a string representation of v's value.
func (s valueS) String() string {
	return s.jsValue
}

// InstanceOf reports whether v is an instance of type t according to JavaScript's instanceof operator.
func (s valueS) InstanceOf(t ValueI) bool {
	return false
}

// Add an event listener to things that can do that such as the window and html elements
func (s valueS) AddEventListener(typ string, useCapture bool, listener func(EventI)) FuncI {
	return nil
}

// remove an event listener to things that they have been added to before
func (s valueS) RemoveEventListener(typ string, useCapture bool, listener FuncI) {

}

// Send an event to trigger event listeners for things that can have listeners
func (s valueS) DispatchEvent(event EventI) bool {
	return false
}
