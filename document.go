package dom

import "time"

// https://developer.mozilla.org/en-US/docs/Web/API/Document/activeElement
type DocumentI interface {
	EventTargetI

	Underlying() ValueI

	DocumentURI() string                           // https://developer.mozilla.org/en-US/docs/Web/API/Document/documentURI
	CreateElement(name string) ElementI            // https://developer.mozilla.org/en-US/docs/Web/API/Document/createElement
	ElementFromPoint(x, y int) ElementI            // https://developer.mozilla.org/en-US/docs/Web/API/Document/elementFromPoint
	GetElementsByClassName(name string) []ElementI // https://developer.mozilla.org/en-US/docs/Web/API/Document/getElementsByClassName
	GetElementsByTagName(name string) []ElementI   // https://developer.mozilla.org/en-US/docs/Web/API/Document/getElementsByTagName
	GetElementByID(id string) ElementI             // https://developer.mozilla.org/en-US/docs/Web/API/Document/getElementById
	QuerySelector(sel string) ElementI             // https://developer.mozilla.org/en-US/docs/Web/API/Document/querySelector
	QuerySelectorAll(sel string) []ElementI        // https://developer.mozilla.org/en-US/docs/Web/API/Document/querySelectorAll

	// HTMLDocument
	ActiveElement() ElementI // https://developer.mozilla.org/en-US/docs/Web/API/Document/activeElement
	Body() ElementI          // https://developer.mozilla.org/en-US/docs/Web/API/Document/body
	Cookie() string          // https://developer.mozilla.org/en-US/docs/Web/API/Document/cookie
	SetCookie(string)        // https://developer.mozilla.org/en-US/docs/Web/API/Document/cookie
	LastModified() time.Time // https://developer.mozilla.org/en-US/docs/Web/API/Document/lastModified
	ReadyState() string      // https://developer.mozilla.org/en-US/docs/Web/API/Document/readyState
	Referrer() string        // https://developer.mozilla.org/en-US/docs/Web/API/Document/referrer
	Title() string           // https://developer.mozilla.org/en-US/docs/Web/API/Document/title
	SetTitle(string)         // https://developer.mozilla.org/en-US/docs/Web/API/Document/titleq
	URL() string             // https://developer.mozilla.org/en-US/docs/Web/API/Document/URL
}

type documentS struct {
	ValueI
	ElementI
}

var _ DocumentI = &documentS{}

func NewDocument(val ValueI) *documentS {
	ret := &documentS{
		ValueI:   val,
		ElementI: NewElement(val),
	}
	return ret
}

func (s *documentS) Underlying() ValueI {
	return s.ValueI
}

func (d *documentS) DocumentURI() string {
	return d.Get("documentURI").String()
}

func (d documentS) CreateElement(name string) ElementI {
	return NewElement(d.Call("createElement", name))
}

func (d documentS) ElementFromPoint(x, y int) ElementI {
	return NewElement(d.Call("elementFromPoint", x, y))
}

func (d documentS) GetElementsByClassName(name string) []ElementI {
	return d.ElementI.GetElementsByClassName(name)
}

func (d documentS) GetElementsByTagName(name string) []ElementI {
	return d.ElementI.GetElementsByTagName(name)
}

func (d documentS) GetElementByID(id string) ElementI {
	return NewElement(d.Call("getElementById", id))
}

func (d documentS) QuerySelector(sel string) ElementI {
	return d.ElementI.QuerySelector(sel)
}

func (d documentS) QuerySelectorAll(sel string) []ElementI {
	return d.ElementI.QuerySelectorAll(sel)
}

////
////
////

func (s *documentS) AddEventListener(typ string, useCapture bool, listener func(EventI)) EventListenerI {
	return s.ValueI.AddEventListener(typ, useCapture, listener)
}

func (s *documentS) RemoveEventListener(listener EventListenerI) {
	s.ValueI.RemoveEventListener(listener)
}

func (s *documentS) DispatchEvent(event EventI) bool {
	return s.ValueI.DispatchEvent(event)
}

////
////
////

func (s *documentS) ActiveElement() ElementI {
	return NewElement(s.Get("activeElement"))
}

func (s *documentS) Body() ElementI {
	return NewElement(s.Get("body"))
}

func (s *documentS) Cookie() string {
	return s.Get("cookie").String()
}

func (s *documentS) SetCookie(in string) {
	s.Set("cookie", in)
}

func (s *documentS) LastModified() time.Time {
	return time.Unix(0, int64(s.Get("lastModified").Call("getTime").Int())*1000000)
}

func (s *documentS) ReadyState() string {
	return s.Get("readyState").String()
}

func (s *documentS) Referrer() string {
	return s.Get("referrer").String()
}

func (s *documentS) Title() string {
	return s.Get("title").String()
}

func (s *documentS) SetTitle(title string) {
	s.Set("title", title)
}

func (s *documentS) URL() string {
	return s.Get("URL").String()
}
