package dom

type NodeI interface {
	EventTargetI

	Underlying() ValueI
	BaseURI() string
	ChildNodes() []NodeI
	FirstChild() NodeI
	LastChild() NodeI
	NextSibling() NodeI
	NodeName() string
	NodeType() int
	NodeValue() string
	SetNodeValue(string)
	ParentNode() NodeI
	ParentElement() ElementI
	PreviousSibling() NodeI
	TextContent() string
	SetTextContent(string)
	AppendChild(NodeI)
	CloneNode(deep bool) NodeI
	CompareDocumentPosition(NodeI) int
	Contains(NodeI) bool
	HasChildNodes() bool
	InsertBefore(which NodeI, before NodeI)
	IsDefaultNamespace(string) bool
	IsEqualNode(NodeI) bool
	LookupPrefix() string
	LookupNamespaceURI(string) string
	Normalize()
	RemoveChild(NodeI)
	ReplaceChild(newChild, oldChild NodeI)
}

type ChildNodeI interface {
	PreviousElementSibling() ElementI
	NextElementSibling() ElementI
}

// Type nodeS implements the Node interface and is embedded by
// concrete node types and element types.
type nodeS struct {
	ValueI
}

var _ NodeI = &nodeS{}

func NewNode(val ValueI) *nodeS {
	ret := &nodeS{
		ValueI: val,
	}
	return ret
}

func (n *nodeS) Underlying() ValueI {
	return n.ValueI
}

func (n *nodeS) BaseURI() string {
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

func nodeListToNodes(o ValueI) []NodeI {
	var out []NodeI
	for _, obj := range nodeListToObjects(o) {
		out = append(out, NewNode(obj))
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

func (n *nodeS) ChildNodes() []NodeI {
	return nodeListToNodes(n.Get("childNodes"))
}

func (n *nodeS) FirstChild() NodeI {
	return NewNode(n.Get("firstChild"))

}

func (n *nodeS) LastChild() NodeI {
	return NewNode(n.Get("lastChild"))

}

func (n *nodeS) NextSibling() NodeI {

	return NewNode(n.Get("nextSibling"))

}

func (n *nodeS) NodeName() string {
	return n.Get("nodeName").String()
}

func (n *nodeS) NodeType() int {
	return n.Get("nodeType").Int()
}

func (n *nodeS) NodeValue() string {
	return n.Get("nodeValue").String()
}

func (n *nodeS) SetNodeValue(s string) {
	n.Set("nodeValue", s)
}

func (n *nodeS) ParentNode() NodeI {
	val := n.Get("parentNode")
	return &nodeS{ValueI: val}
}

func (n *nodeS) ParentElement() ElementI {
	val := n.Get("parentElement")
	return NewElement(val)
}

func (n *nodeS) PreviousSibling() NodeI {
	return NewNode(n.Get("previousSibling"))
}

func (n *nodeS) TextContent() string {
	return n.Get("textContent").String()
}

func (n *nodeS) SetTextContent(s string) {
	n.Set("textContent", s)
}

func (n *nodeS) AppendChild(newchild NodeI) {
	n.Call("appendChild", newchild.Underlying())
}

func (n *nodeS) CloneNode(deep bool) NodeI {
	return NewNode(n.Call("cloneNode", deep))
}

func (n *nodeS) CompareDocumentPosition(other NodeI) int {
	return n.Call("compareDocumentPosition", other.Underlying()).Int()
}

func (n *nodeS) Contains(other NodeI) bool {
	return n.Call("contains", other.Underlying()).Bool()
}

func (n *nodeS) HasChildNodes() bool {
	return n.Call("hasChildNodes").Bool()
}

func (n *nodeS) InsertBefore(which NodeI, before NodeI) {
	var o interface{}
	if before != nil {
		o = before.Underlying()
	}
	n.Call("insertBefore", which.Underlying(), o)
}

func (n *nodeS) IsDefaultNamespace(s string) bool {
	return n.Call("isDefaultNamespace", s).Bool()
}

func (n *nodeS) IsEqualNode(other NodeI) bool {
	return n.Call("isEqualNode", other.Underlying()).Bool()
}

func (n *nodeS) LookupPrefix() string {
	return n.Call("lookupPrefix").String()
}

func (n *nodeS) LookupNamespaceURI(s string) string {
	return n.Call("lookupNamespaceURI", s).String()
}

func (n *nodeS) Normalize() {
	n.Call("normalize")
}

func (n *nodeS) RemoveChild(other NodeI) {
	n.Call("removeChild", other.Underlying())
}

func (n *nodeS) ReplaceChild(newChild, oldChild NodeI) {
	n.Call("replaceChild", newChild.Underlying(), oldChild.Underlying())
}
