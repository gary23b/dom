package dom

import (
	"time"
)

type EventI interface {
	Bubbles() bool
	Cancelable() bool
	CurrentTarget() ElementI
	DefaultPrevented() bool
	EventPhase() int
	Target() ElementI
	Timestamp() time.Time
	Type() string
	PreventDefault()
	StopImmediatePropagation()
	StopPropagation()
	Underlying() ValueI
}

type EventTargetI interface {
	// AddEventListener adds a new event listener and returns the
	// wrapper function it generated. If using RemoveEventListener,
	// that wrapper has to be used.
	AddEventListener(typ string, useCapture bool, listener func(EventI)) EventListenerI
	RemoveEventListener(listener EventListenerI)
	DispatchEvent(event EventI) bool
}

// Type BasicEvent implements the Event interface and is embedded by
// concrete eventS types.
type eventS struct {
	ValueI
}

var _ EventI = eventS{}

func CreateEvent(window WindowI, typ string, bubbles, cancelable bool) eventS {
	var event = window.Underlying().Get("Event").New(
		typ,
		map[string]interface{}{
			"bubbles":    bubbles,
			"cancelable": cancelable,
		},
	)
	return eventS{ValueI: event}
}

func (ev eventS) Bubbles() bool {
	return ev.Get("bubbles").Bool()
}

func (ev eventS) Cancelable() bool {
	return ev.Get("cancelable").Bool()
}

func (ev eventS) CurrentTarget() ElementI {
	val := ev.Get("currentTarget")

	return &elementS{ValueI: val}
}

func (ev eventS) DefaultPrevented() bool {
	return ev.Get("defaultPrevented").Bool()
}

func (ev eventS) EventPhase() int {
	return ev.Get("eventPhase").Int()
}

func (ev eventS) Target() ElementI {
	val := ev.Get("target")
	return &elementS{ValueI: val}
}

func (ev eventS) Timestamp() time.Time {
	ms := ev.Get("timeStamp").Int()
	s := ms / 1000
	ns := (ms % 1000 * 1e6)
	return time.Unix(int64(s), int64(ns))
}

func (ev eventS) Type() string {
	return ev.Get("type").String()
}

func (ev eventS) PreventDefault() {
	ev.Call("preventDefault")
}

func (ev eventS) StopImmediatePropagation() {
	ev.Call("stopImmediatePropagation")
}

func (ev eventS) StopPropagation() {
	ev.Call("stopPropagation")
}

func (ev eventS) Underlying() ValueI {
	return ev.ValueI
}

////
////
////
////

type EventListenerI interface {
	Underlying() FuncI
	GetID() string
	GetType() string
	GetCapture() bool
}

type EventTargetS struct {
	FuncI
	id      string
	typ     string
	capture bool
}

var _ EventListenerI = EventTargetS{}

func NewEventListener(fn FuncI, typ string, capture bool) EventTargetS {
	ret := EventTargetS{
		FuncI:   fn,
		id:      GetNextID(),
		typ:     typ,
		capture: capture,
	}
	return ret
}

func (s EventTargetS) Underlying() FuncI {
	return s.FuncI
}

func (s EventTargetS) GetID() string {
	return s.id
}

func (s EventTargetS) GetType() string {
	return s.typ
}

func (s EventTargetS) GetCapture() bool {
	return s.capture
}
