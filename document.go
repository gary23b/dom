package dom

import "time"

type DocumentI interface {
	NodeI

	Async() bool
	SetAsync(bool)
	DocumentElement() ElementI
	DocumentURI() string
	AdoptNode(node NodeI) NodeI
	ImportNode(node NodeI, deep bool) NodeI
	CreateElement(name string) ElementI
	ElementFromPoint(x, y int) ElementI
	EnableStyleSheetsForSet(name string)
	GetElementsByClassName(name string) []ElementI
	GetElementsByTagName(name string) []ElementI
	GetElementByID(id string) ElementI
	QuerySelector(sel string) ElementI
	QuerySelectorAll(sel string) []ElementI

	// HTMLDocument
	ActiveElement() ElementI
	Body() ElementI
	Cookie() string
	SetCookie(string)
	Domain() string
	SetDomain(string)
	LastModified() time.Time
	Links() []ElementI
	ReadyState() string
	Referrer() string
	Title() string
	SetTitle(string)
	URL() string
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

func (d *documentS) Async() bool {
	return d.Get("async").Bool()
}

func (d *documentS) SetAsync(b bool) {
	d.Set("async", b)
}

func (d *documentS) DocumentElement() ElementI {
	val := d.Get("documentElement")
	return NewElement(val)
}

func (d *documentS) DocumentURI() string {
	return d.Get("documentURI").String()
}

func (d documentS) LastStyleSheetSet() string {
	return d.Get("lastStyleSheetSet").String()
}

func (d documentS) PreferredStyleSheetSet() string {
	return d.Get("preferredStyleSheetSet").String()
}

func (d documentS) SelectedStyleSheetSet() string {
	return d.Get("selectedStyleSheetSet").String()
}

func (d documentS) AdoptNode(node NodeI) NodeI {
	val := d.Call("adoptNode", node.Underlying())
	return NewNode(val)
}

func (d documentS) ImportNode(node NodeI, deep bool) NodeI {
	val := d.Call("importNode", node.Underlying(), deep)
	return NewNode(val)
}

func (d documentS) CreateElement(name string) ElementI {
	return NewElement(d.Call("createElement", name))
}

func (d documentS) ElementFromPoint(x, y int) ElementI {
	return NewElement(d.Call("elementFromPoint", x, y))
}

func (d documentS) EnableStyleSheetsForSet(name string) {
	d.Call("enableStyleSheetsForSet", name)
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

func (s *documentS) AddEventListener(typ string, useCapture bool, listener func(EventI)) FuncI {
	return s.ValueI.AddEventListener(typ, useCapture, listener)
}

func (s *documentS) RemoveEventListener(typ string, useCapture bool, listener FuncI) {
	s.ValueI.RemoveEventListener(typ, useCapture, listener)
}

func (s *documentS) DispatchEvent(event EventI) bool {
	return s.ValueI.DispatchEvent(event)
}

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

func (s *documentS) DefaultView() WindowI {
	return &window{s.Get("defaultView")}
}

func (s *documentS) DesignMode() bool {
	val := s.Get("designMode").String()
	return val != "off"
}

func (s *documentS) SetDesignMode(b bool) {
	val := "off"
	if b {
		val = "on"
	}
	s.Set("designMode", val)
}

func (s *documentS) Domain() string {
	return s.Get("domain").String()
}

func (s *documentS) SetDomain(in string) {
	s.Set("domain", in)
}

func (s *documentS) LastModified() time.Time {
	return time.Unix(0, int64(s.Get("lastModified").Call("getTime").Int())*1000000)
}

func (s *documentS) Links() []ElementI {
	var els []ElementI
	links := s.Get("links")
	length := links.Get("length").Int()
	for i := 0; i < length; i++ {
		els = append(els, NewElement(links.Call("item", i)))
	}
	return els
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
