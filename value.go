package dom

/*
Doing a similar thing to: https://github.com/maxence-charriere/go-app/blob/master/pkg/app/js.go
ValueI is an exact wrapping of the capabilities of the "syscall/js" package.
*/

type ValueI interface {
	// Equal reports whether v and w are equal according to JavaScript's === operator.
	Equal(w ValueI) bool

	// IsUndefined reports whether v is the JavaScript value "undefined".
	IsUndefined() bool

	// IsNull reports whether v is the JavaScript value "null".
	IsNull() bool

	// IsNaN reports whether v is the JavaScript value "NaN".
	IsNaN() bool

	// Type returns the JavaScript type of the value v. It is similar to JavaScript's typeof operator,
	// except that it returns TypeNull instead of TypeObject for null.
	Type() Type

	// Get returns the JavaScript property p of value v.
	// It panics if v is not a JavaScript object.
	Get(p string) ValueI

	// Set sets the JavaScript property p of value v to ValueOf(x).
	// It panics if v is not a JavaScript object.
	Set(p string, x any)

	// Delete deletes the JavaScript property p of value v.
	// It panics if v is not a JavaScript object.
	Delete(p string)

	// Index returns JavaScript index i of value v.
	// It panics if v is not a JavaScript object.
	Index(i int) ValueI

	// SetIndex sets the JavaScript index i of value v to ValueOf(x).
	// It panics if v is not a JavaScript object.
	SetIndex(i int, x any)

	// Length returns the JavaScript property "length" of v.
	// It panics if v is not a JavaScript object.
	Length() int

	// Call does a JavaScript call to the method m of value v with the given arguments.
	// It panics if v has no method m.
	// The arguments get mapped to JavaScript values according to the ValueOf function.
	Call(m string, args ...any) ValueI

	// Invoke does a JavaScript call of the value v with the given arguments.
	// It panics if v is not a JavaScript function.
	// The arguments get mapped to JavaScript values according to the ValueOf function.
	Invoke(args ...any) ValueI

	// New uses JavaScript's "new" operator with value v as constructor and the given arguments.
	// It panics if v is not a JavaScript function.
	// The arguments get mapped to JavaScript values according to the ValueOf function.
	New(args ...any) ValueI

	// Float returns the value v as a float64.
	// It panics if v is not a JavaScript number.
	Float() float64

	// Int returns the value v truncated to an int.
	// It panics if v is not a JavaScript number.
	Int() int

	// Bool returns the value v as a bool.
	// It panics if v is not a JavaScript boolean.
	Bool() bool

	// Truthy returns the JavaScript "truthiness" of the value v. In JavaScript,
	// false, 0, "", null, undefined, and NaN are "falsy", and everything else is
	// "truthy". See https://developer.mozilla.org/en-US/docs/Glossary/Truthy.
	Truthy() bool

	// String returns the value v as a string.
	// String is a special case because of Go's String method convention. Unlike the other getters,
	// it does not panic if v's Type is not TypeString. Instead, it returns a string of the form "<T>"
	// or "<T: V>" where T is v's type and V is a string representation of v's value.
	String() string

	// InstanceOf reports whether v is an instance of type t according to JavaScript's instanceof operator.
	InstanceOf(t ValueI) bool

	// Add an event listener to things that can do that such as the window and html elements
	AddEventListener(typ string, useCapture bool, listener func(EventI)) FuncI

	// remove an event listener to things that they have been added to before
	RemoveEventListener(typ string, useCapture bool, listener FuncI)

	// Send an event to trigger event listeners for things that can have listeners
	DispatchEvent(event EventI) bool
}

// Type represents the JavaScript type of a Value.
type Type int

const (
	TypeUndefined Type = iota
	TypeNull
	TypeBoolean
	TypeNumber
	TypeString
	TypeSymbol
	TypeObject
	TypeFunction
)

func (t Type) String() string {
	switch t {
	case TypeUndefined:
		return "undefined"
	case TypeNull:
		return "null"
	case TypeBoolean:
		return "boolean"
	case TypeNumber:
		return "number"
	case TypeString:
		return "string"
	case TypeSymbol:
		return "symbol"
	case TypeObject:
		return "object"
	case TypeFunction:
		return "function"
	default:
		panic("bad type")
	}
}
