//go:build js && wasm

package dom

import (
	"reflect"
	"syscall/js"
)

var (
	null   = js.ValueOf(nil)
	object = js.Global().Get("Object")
	array  = js.Global().Get("Array")
)

type valueS struct {
	jsValue js.Value
}

var _ ValueI = valueS{}

func NewValue(jsValue js.Value) valueS {
	ret := valueS{
		jsValue: jsValue,
	}
	return ret
}

// // JSValue returns the JS value.
func (v valueS) JSValue() js.Value {
	return v.jsValue
}

// Equal reports whether v and w are equal according to JavaScript's === operator.
func (s valueS) Equal(w ValueI) bool {
	other := w.(valueS)
	return s.jsValue.Equal(other.jsValue)
}

// IsUndefined reports whether v is the JavaScript value "undefined".
func (s valueS) IsUndefined() bool {
	return s.jsValue.IsUndefined()
}

// IsNull reports whether v is the JavaScript value "null".
func (s valueS) IsNull() bool {
	return s.jsValue.IsNull()
}

// IsNaN reports whether v is the JavaScript value "NaN".
func (s valueS) IsNaN() bool {
	return s.jsValue.IsNaN()
}

// Type returns the JavaScript type of the value v. It is similar to JavaScript's typeof operator,
// except that it returns TypeNull instead of TypeObject for null.
func (s valueS) Type() Type {
	return Type(s.jsValue.Type())
}

// Get returns the JavaScript property p of value v.
// It panics if v is not a JavaScript object.
func (s valueS) Get(p string) ValueI {
	got := s.jsValue.Get(p)
	// fmt.Println("js.Value.Get()", p, got)
	return valueS{jsValue: got}
}

// Set sets the JavaScript property p of value v to ValueOf(x).
// It panics if v is not a JavaScript object.
func (s valueS) Set(p string, x any) {
	valConverted := ValueOf(x)
	s.jsValue.Set(p, valConverted.jsValue)
}

// Delete deletes the JavaScript property p of value v.
// It panics if v is not a JavaScript object.
func (s valueS) Delete(p string) {
	s.jsValue.Delete(p)
}

// Index returns JavaScript index i of value v.
// It panics if v is not a JavaScript object.
func (s valueS) Index(i int) ValueI {
	got := s.jsValue.Index(i)
	return valueS{jsValue: got}
}

// SetIndex sets the JavaScript index i of value v to ValueOf(x).
// It panics if v is not a JavaScript object.
func (s valueS) SetIndex(i int, x any) {
	s.jsValue.SetIndex(i, x)
}

// Length returns the JavaScript property "length" of v.
// It panics if v is not a JavaScript object.
func (s valueS) Length() int {
	return s.jsValue.Length()
}

// Call does a JavaScript call to the method m of value v with the given arguments.
// It panics if v has no method m.
// The arguments get mapped to JavaScript values according to the ValueOf function.
func (s valueS) Call(m string, args ...any) ValueI {
	convertedArgs := convertArgsToJsValue(args)
	// fmt.Println(m, args, convertedArgs)
	got := s.jsValue.Call(m, convertedArgs...)
	return valueS{jsValue: got}
}

// Invoke does a JavaScript call of the value v with the given arguments.
// It panics if v is not a JavaScript function.
// The arguments get mapped to JavaScript values according to the ValueOf function.
func (s valueS) Invoke(args ...any) ValueI {
	convertedArgs := convertArgsToJsValue(args)
	got := s.jsValue.Invoke(convertedArgs...)
	return valueS{jsValue: got}
}

// New uses JavaScript's "new" operator with value v as constructor and the given arguments.
// It panics if v is not a JavaScript function.
// The arguments get mapped to JavaScript values according to the ValueOf function.
func (s valueS) New(args ...any) ValueI {
	convertedArgs := convertArgsToJsValue(args)
	// fmt.Println("new", args, convertedArgs)
	got := s.jsValue.New(convertedArgs...)
	return valueS{jsValue: got}
}

func (s valueS) Float() float64 {
	return s.jsValue.Float()
}

// Int returns the value v truncated to an int.
// It panics if v is not a JavaScript number.
func (s valueS) Int() int {
	return s.jsValue.Int()
}

// Bool returns the value v as a bool.
// It panics if v is not a JavaScript boolean.
func (s valueS) Bool() bool {
	return s.jsValue.Bool()
}

// Truthy returns the JavaScript "truthiness" of the value v. In JavaScript,
// false, 0, "", null, undefined, and NaN are "falsy", and everything else is
// "truthy". See https://developer.mozilla.org/en-US/docs/Glossary/Truthy.
func (s valueS) Truthy() bool {
	return s.jsValue.Truthy()
}

// String returns the value v as a string.
// String is a special case because of Go's String method convention. Unlike the other getters,
// it does not panic if v's Type is not TypeString. Instead, it returns a string of the form "<T>"
// or "<T: V>" where T is v's type and V is a string representation of v's value.
func (s valueS) String() string {
	if s.jsValue.IsNull() || s.jsValue.IsUndefined() {
		return ""
	}
	return s.jsValue.String()
}

// InstanceOf reports whether v is an instance of type t according to JavaScript's instanceof operator.
func (s valueS) InstanceOf(t ValueI) bool {
	other := t.(valueS)
	return s.jsValue.InstanceOf(other.jsValue)
}

func (s valueS) AddEventListener(typ string, useCapture bool, listener func(EventI)) EventListenerI {
	wrapperJsFunc := NewFuncForJavascript(func(this ValueI, args []ValueI) any {
		arg := args[0]
		var e *eventS
		if !arg.IsNull() && !arg.IsUndefined() {
			jsArg := arg.(valueS)
			e = &eventS{ValueI: valueS{jsValue: jsArg.jsValue}}
		}
		listener(e)
		return nil
	})

	s.Call("addEventListener", typ, wrapperJsFunc, useCapture)

	ret := NewEventListener(wrapperJsFunc, typ, useCapture)
	return ret
}

func (s valueS) RemoveEventListener(listener EventListenerI) {
	fn := listener.Underlying()
	value := fn.(funcS)
	s.Call("removeEventListener", listener.GetType(), value.Func, listener.GetCapture())
	fn.Release()
}

func (s valueS) DispatchEvent(event EventI) bool {
	return s.Call("dispatchEvent", event).Bool()
}

//
//
//

func convertArgsToJsValue(args []any) []any {
	ret := make([]any, 0, len(args))

	// fmt.Printf("args, %d, %#v\n", len(args), args)

	for _, arg := range args {
		switch v := arg.(type) {
		case funcS:
			ret = append(ret, v.Func)
		case *funcS:
			if v == nil {
				ret = append(ret, valueS{jsValue: null})
			} else {
				ret = append(ret, v.Func)
			}

		default:
			// fmt.Printf("arg, %#v\n", arg)
			val := ValueOf(arg)
			// fmt.Printf("arg, %#v = %#v\n", arg, val.jsValue)
			ret = append(ret, val.jsValue)
		}
	}

	return ret
}

// ValueOf returns the Go value as a new value.
func ValueOf(i any) valueS {
	switch v := i.(type) {
	case nil:
		return valueS{jsValue: null}
	case js.Value:
		return valueS{jsValue: v}
	case valueS:
		return v
	case *valueS:
		if v == nil {
			return valueS{jsValue: null}
		}
		return *v
	default:
		rv := reflect.ValueOf(i)
		return valueS{jsValue: valueOf(rv)}
	}
}

// valueOf recursively returns a new value.
func valueOf(v reflect.Value) js.Value {
	switch v.Kind() {
	case reflect.Ptr, reflect.Interface:
		return valueOfPointerOrInterface(v)
	case reflect.Slice, reflect.Array:
		return valueOfSliceOrArray(v)
	case reflect.Map:
		return valueOfMap(v)
	case reflect.Struct:
		return valueOfStruct(v)
	default:
		// fmt.Printf("choosing default\n")
		return js.ValueOf(v.Interface())
	}
}

// valueOfPointerOrInterface returns a new value.
func valueOfPointerOrInterface(v reflect.Value) js.Value {
	if v.IsNil() {
		return null
	}
	return valueOf(v.Elem())
}

// valueOfSliceOrArray returns a new array object value.
func valueOfSliceOrArray(v reflect.Value) js.Value {
	if v.IsNil() {
		return null
	}
	a := array.New()
	n := v.Len()
	for i := range n {
		e := v.Index(i)
		a.SetIndex(i, valueOf(e))
	}
	return a
}

// valueOfMap returns a new object value.
// Map keys must be of type string.
func valueOfMap(v reflect.Value) js.Value {
	if v.IsNil() {
		return null
	}
	m := object.New()
	i := v.MapRange()
	for i.Next() {
		k := i.Key().Interface().(string)
		m.Set(k, valueOf(i.Value()))
	}
	return m
}

// valueOfStruct returns a new object value.
func valueOfStruct(v reflect.Value) js.Value {
	t := v.Type()
	s := object.New()
	n := v.NumField()
	for i := range n {
		if f := v.Field(i); f.CanInterface() {
			k := nameOf(t.Field(i))
			s.Set(k, valueOf(f))
		}
	}
	return s
}

// nameOf returns the JS tag name, otherwise the field name.
func nameOf(sf reflect.StructField) string {
	name := sf.Tag.Get("js")
	if name == "" {
		name = sf.Tag.Get("json")
	}
	if name == "" {
		return sf.Name
	}
	return name
}
