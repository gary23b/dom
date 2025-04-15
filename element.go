package dom

import (
	"fmt"
	"sync"
)

type ElementI interface {
	EventTargetI
	RemoveAllEventListeners()

	// node
	Underlying() ValueI
	BaseURI() string
	ChildNodes() []ElementI
	FirstChild() ElementI
	LastChild() ElementI
	NextSibling() ElementI
	NodeName() string
	NodeType() int
	NodeValue() string
	SetNodeValue(string)
	ParentNode() ElementI
	PreviousSibling() ElementI
	TextContent() string
	SetTextContent(string)
	AppendChild(ElementI) ElementI
	NewChild(typ string) ElementI
	Clone(deep bool) ElementI
	CompareDocumentPosition(ElementI) int
	Contains(ElementI) bool
	HasChildNodes() bool
	InsertBefore(which ElementI, before ElementI)
	IsEqualNode(ElementI) bool
	LookupPrefix() string
	Normalize()
	RemoveChild(ElementI)
	ReplaceChild(newChild, oldChild ElementI)

	// element
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

type idMakerT struct {
	mutex     sync.Mutex
	idCounter int
}

// Make a global ID maker that can dispense IDs
var idMaker idMakerT

func GetNextID() string {
	idMaker.mutex.Lock()
	defer idMaker.mutex.Unlock()

	newID := fmt.Sprintf("id_%06d", idMaker.idCounter)
	idMaker.idCounter++
	return newID
}

type elementS struct {
	ValueI
	id             string
	children       map[string]ElementI
	eventListeners map[string]EventListenerI
}

var _ ElementI = &elementS{}

func NewElement(val ValueI) *elementS {
	ret := &elementS{
		ValueI:         val,
		eventListeners: map[string]EventListenerI{},
		id:             GetNextID(),
		children:       map[string]ElementI{},
	}

	val.Set("id", ret.id)
	return ret
}

func (n *elementS) Underlying() ValueI {
	return n.ValueI
}

func (n *elementS) BaseURI() string {
	return n.Get("baseURI").String()
}

func arrayToObjects(o ValueI) []ValueI {
	var out []ValueI
	for i := 0; i < o.Length(); i++ {
		out = append(out, o.Index(i))
	}
	return out
}

func nodeListToObjects(o ValueI) []ValueI {
	if o.Get("constructor").Equal(valueS{jsValue: array}) {
		// Support Polymer's DOM APIs, which uses Arrays instead of
		// NodeLists
		return arrayToObjects(o)
	}
	var out []ValueI
	length := o.Get("length").Int()
	for i := 0; i < length; i++ {
		out = append(out, o.Call("item", i))
	}
	return out
}

func nodeListToNodes(o ValueI) []ElementI {
	var out []ElementI
	for _, obj := range nodeListToObjects(o) {
		out = append(out, NewElement(obj))
	}
	return out
}

func nodeListToElements(o ValueI) []ElementI {
	var out []ElementI
	for _, obj := range nodeListToObjects(o) {
		out = append(out, NewElement(obj))
	}
	return out
}

func (n *elementS) ChildNodes() []ElementI {
	return nodeListToNodes(n.Get("childNodes"))
}

func (n *elementS) FirstChild() ElementI {
	return NewElement(n.Get("firstChild"))

}

func (n *elementS) LastChild() ElementI {
	return NewElement(n.Get("lastChild"))
}

func (n *elementS) NextSibling() ElementI {
	return NewElement(n.Get("nextSibling"))
}

func (n *elementS) NodeName() string {
	return n.Get("nodeName").String()
}

func (n *elementS) NodeType() int {
	return n.Get("nodeType").Int()
}

func (n *elementS) NodeValue() string {
	return n.Get("nodeValue").String()
}

func (n *elementS) SetNodeValue(s string) {
	n.Set("nodeValue", s)
}

func (n *elementS) ParentNode() ElementI {
	return NewElement(n.Get("parentNode"))
}

func (n *elementS) PreviousSibling() ElementI {
	return NewElement(n.Get("previousSibling"))
}

func (n *elementS) TextContent() string {
	return n.Get("textContent").String()
}

func (n *elementS) SetTextContent(s string) {
	n.Set("textContent", s)
}

func (n *elementS) AppendChild(newChild ElementI) ElementI {
	n.children[newChild.ID()] = newChild
	n.Call("appendChild", newChild.Underlying())
	return newChild
}

func (n *elementS) NewChild(typ string) ElementI {
	newElement := Doc.CreateElement(typ)
	n.AppendChild(newElement)
	return newElement
}

func (n *elementS) Clone(deep bool) ElementI {
	return NewElement(n.Call("cloneNode", deep))
}

func (n *elementS) CompareDocumentPosition(other ElementI) int {
	return n.Call("compareDocumentPosition", other.Underlying()).Int()
}

func (n *elementS) Contains(other ElementI) bool {
	return n.Call("contains", other.Underlying()).Bool()
}

func (n *elementS) HasChildNodes() bool {
	return n.Call("hasChildNodes").Bool()
}

func (n *elementS) InsertBefore(which ElementI, before ElementI) {
	var o interface{}
	if before != nil {
		o = before.Underlying()
	}
	n.Call("insertBefore", which.Underlying(), o)
}

func (n *elementS) IsEqualNode(other ElementI) bool {
	return n.Call("isEqualNode", other.Underlying()).Bool()
}

func (n *elementS) LookupPrefix() string {
	return n.Call("lookupPrefix").String()
}

func (n *elementS) Normalize() {
	n.Call("normalize")
}

func (n *elementS) RemoveChild(other ElementI) {
	n.Call("removeChild", other.Underlying())
}

func (n *elementS) ReplaceChild(newChild, oldChild ElementI) {
	n.Call("replaceChild", newChild.Underlying(), oldChild.Underlying())
}

/////////////////////
/////////////////////
/////////////////////
/////////////////////
/////////////////////
/////////////////////

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
	if e.id == "" {
		e.id = e.Get("id").String()

	}
	return e.id
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

func (s *elementS) AddEventListener(typ string, useCapture bool, listener func(EventI)) EventListenerI {
	ret := s.ValueI.AddEventListener(typ, useCapture, listener)
	s.eventListeners[ret.GetID()] = ret
	return ret
}

func (s *elementS) RemoveEventListener(listener EventListenerI) {
	s.ValueI.RemoveEventListener(listener)
	delete(s.eventListeners, listener.GetID())
}

func (s *elementS) RemoveAllEventListeners() {
	for _, eventListener := range s.eventListeners {
		s.RemoveEventListener(eventListener)
	}
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
