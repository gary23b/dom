package dom

type ElementI interface {
	NodeI
	ChildNodeI

	Attributes() map[string]string
	Class() TokenListI
	Closest(string) ElementI
	ID() string
	SetID(string)
	TagName() string
	GetAttribute(string) string
	GetBoundingClientRect() RectI
	GetElementsByClassName(string) []ElementI
	GetElementsByTagName(string) []ElementI
	HasAttribute(string) bool
	Matches(string) bool
	QuerySelector(string) ElementI
	QuerySelectorAll(string) []ElementI
	Remove()
	RemoveAttribute(string)
	SetAttribute(name string, value string)
	InnerHTML() string
	SetInnerHTML(string)
	OuterHTML() string
	SetOuterHTML(string)

	// HTML Element
	AccessKey() string
	Dataset() map[string]string
	SetAccessKey(string)
	AccessKeyLabel() string
	SetAccessKeyLabel(string)
	ContentEditable() string
	SetContentEditable(string)
	IsContentEditable() bool
	Dir() string
	SetDir(string)
	Draggable() bool
	SetDraggable(bool)
	Lang() string
	SetLang(string)
	OffsetHeight() float64
	OffsetLeft() float64
	OffsetParent() ElementI
	OffsetTop() float64
	OffsetWidth() float64
	Style() CSSStyleI
	Title() string
	SetTitle(string)
	Blur()
	Click()
	Focus()
}

type CSSStyleI interface {
	ToMap() map[string]string
	RemoveProperty(name string)
	GetPropertyValue(name string) string
	GetPropertyPriority(name string) string
	SetProperty(name, value, priority string)
	Index(idx int) string
	Length() int
}

type elementS struct {
	ValueI
	NodeI
}

var _ ElementI = &elementS{}

func NewElement(val ValueI) *elementS {
	ret := &elementS{
		ValueI: val,
		NodeI:  &nodeS{ValueI: val},
	}
	return ret
}

func (n *elementS) Underlying() ValueI {
	return n.ValueI
}

func (e *elementS) Attributes() map[string]string {
	o := e.Get("attributes")
	attrs := map[string]string{}
	length := o.Get("length").Int()
	for i := 0; i < length; i++ {
		item := o.Call("item", i)
		attrs[item.Get("name").String()] = item.Get("value").String()
	}
	return attrs
}

func (e *elementS) GetBoundingClientRect() RectI {
	return NewRect(e.Call("getBoundingClientRect"))
}

func (e *elementS) PreviousElementSibling() ElementI {
	return NewElement(e.Get("previousElementSibling"))
}

func (e *elementS) NextElementSibling() ElementI {
	return NewElement(e.Get("nextElementSibling"))
}

func (e *elementS) Class() TokenListI {
	return NewTokenList(e.Get("classList"))
}

// SetClass sets the element's className attribute to s. Consider
// using the Class method instead.
func (e *elementS) SetClass(s string) {
	e.Set("className", s)
}

func (e *elementS) Closest(s string) ElementI {
	return NewElement(NewElement(e.Call("closest", s)))
}

func (e *elementS) ID() string {
	return e.Get("id").String()
}

func (e *elementS) SetID(s string) {
	e.Set("id", s)
}

func (e *elementS) TagName() string {
	return e.Get("tagName").String()
}

func (e *elementS) GetAttribute(name string) string {
	return e.Call("getAttribute", name).String()
}

func (e *elementS) GetElementsByClassName(s string) []ElementI {
	return nodeListToElements(e.Call("getElementsByClassName", s))
}

func (e *elementS) GetElementsByTagName(s string) []ElementI {
	return nodeListToElements(e.Call("getElementsByTagName", s))
}

func (e *elementS) HasAttribute(s string) bool {
	return e.Call("hasAttribute", s).Bool()
}

func (e *elementS) Matches(s string) bool {
	return e.Call("matches", s).Bool()
}

func (e *elementS) QuerySelector(s string) ElementI {
	return NewElement(e.Call("querySelector", s))
}

func (e *elementS) QuerySelectorAll(s string) []ElementI {
	return nodeListToElements(e.Call("querySelectorAll", s))
}

func (e *elementS) Remove() {
	e.Call("remove")
}

func (e *elementS) RemoveAttribute(s string) {
	e.Call("removeAttribute", s)
}

func (e *elementS) SetAttribute(name string, value string) {
	e.Call("setAttribute", name, value)
}

func (e *elementS) InnerHTML() string {
	return e.Get("innerHTML").String()
}

func (e *elementS) SetInnerHTML(s string) {
	e.Set("innerHTML", s)
}

func (e *elementS) OuterHTML() string {
	return e.Get("outerHTML").String()
}

func (e *elementS) SetOuterHTML(s string) {
	e.Set("outerHTML", s)
}

func (s *elementS) AddEventListener(typ string, useCapture bool, listener func(EventI)) FuncI {
	return s.ValueI.AddEventListener(typ, useCapture, listener)
}

func (s *elementS) RemoveEventListener(typ string, useCapture bool, listener FuncI) {
	s.ValueI.RemoveEventListener(typ, useCapture, listener)
}

func (s *elementS) DispatchEvent(event EventI) bool {
	return s.ValueI.DispatchEvent(event)
}

/////
/////

func (e *elementS) AccessKey() string {
	return e.Get("accessKey").String()
}

func (e *elementS) Dataset() map[string]string {
	o := e.Get("dataset")
	data := map[string]string{}
	keys := jsKeys(o)
	for _, key := range keys {
		data[key] = o.Get(key).String()
	}
	return data
}

// jsKeys returns the keys of the given JavaScript object.
func jsKeys(o ValueI) []string {
	if o.IsNull() || o.IsUndefined() {
		return nil
	}
	a := Window.Underlying().Get("Object").Call("keys", o)
	s := make([]string, a.Length())
	for i := 0; i < a.Length(); i++ {
		s[i] = a.Index(i).String()
	}
	return s
}

func (e *elementS) SetAccessKey(s string) {
	e.Set("accessKey", s)
}

func (e *elementS) AccessKeyLabel() string {
	return e.Get("accessKeyLabel").String()
}

func (e *elementS) SetAccessKeyLabel(s string) {
	e.Set("accessKeyLabel", s)
}

func (e *elementS) ContentEditable() string {
	return e.Get("contentEditable").String()
}

func (e *elementS) SetContentEditable(s string) {
	e.Set("contentEditable", s)
}

func (e *elementS) IsContentEditable() bool {
	return e.Get("isContentEditable").Bool()
}

func (e *elementS) Dir() string {
	return e.Get("dir").String()
}

func (e *elementS) SetDir(s string) {
	e.Set("dir", s)
}

func (e *elementS) Draggable() bool {
	return e.Get("draggable").Bool()
}

func (e *elementS) SetDraggable(b bool) {
	e.Set("draggable", b)
}

func (e *elementS) Lang() string {
	return e.Get("lang").String()
}

func (e *elementS) SetLang(s string) {
	e.Set("lang", s)
}

func (e *elementS) OffsetHeight() float64 {
	return e.Get("offsetHeight").Float()
}

func (e *elementS) OffsetLeft() float64 {
	return e.Get("offsetLeft").Float()
}

func (e *elementS) OffsetParent() ElementI {
	return NewElement(e.Get("offsetParent"))
}

func (e *elementS) OffsetTop() float64 {
	return e.Get("offsetTop").Float()
}

func (e *elementS) OffsetWidth() float64 {
	return e.Get("offsetWidth").Float()
}

func (e *elementS) Style() CSSStyleI {
	return NewCssStyle(e.Get("style"))
}

func (e *elementS) TabIndex() int {
	return e.Get("tabIndex").Int()
}

func (e *elementS) SetTabIndex(i int) {
	e.Set("tabIndex", i)
}

func (e *elementS) Title() string {
	return e.Get("title").String()
}

func (e *elementS) SetTitle(s string) {
	e.Set("title", s)
}

func (e *elementS) Blur() {
	e.Call("blur")
}

func (e *elementS) Click() {
	e.Call("click")
}

func (e *elementS) Focus() {
	e.Call("focus")
}

func NewCssStyle(val ValueI) cssStyleS {
	ret := cssStyleS{
		ValueI: val,
	}
	return ret
}

type cssStyleS struct {
	ValueI
}

var _ CSSStyleI = cssStyleS{}

func (s cssStyleS) ToMap() map[string]string {
	m := make(map[string]string)
	N := s.Get("length").Int()
	for i := 0; i < N; i++ {
		name := s.Call("item", i).String()
		value := s.Call("getPropertyValue", name).String()
		m[name] = value
	}

	return m
}

func (s cssStyleS) RemoveProperty(name string) { s.Call("removeProperty", name) }

func (s cssStyleS) GetPropertyValue(name string) string {
	return s.Call("getPropertyValue", name).String()
}

func (s cssStyleS) GetPropertyPriority(name string) string {
	return s.Call("getPropertyPriority", name).String()
}

func (s cssStyleS) SetProperty(name, value, priority string) {
	s.Call("setProperty", name, value, priority)
}

func (s cssStyleS) Index(idx int) string {
	return s.Call("index", idx).String()
}

func (s cssStyleS) Length() int {
	return s.Get("length").Int()
}
