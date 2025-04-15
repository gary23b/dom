package dom

type WindowI interface {
	EventTargetI

	Underlying() ValueI
	Document() DocumentI
	FrameElement() ElementI
	Location() LocationI
	Name() string
	SetName(string)
	InnerHeight() int
	InnerWidth() int
	Length() int
	Opener() WindowI
	OuterHeight() int
	OuterWidth() int
	ScrollX() int
	ScrollY() int
	Parent() WindowI
	ScreenX() int
	ScreenY() int
	ScrollMaxX() int
	ScrollMaxY() int
	Top() WindowI
	History() HistoryI
	Screen() ScreenI
	Alert(string)
	Back()
	Blur()
	Close()
	Confirm(string) bool
	Focus()
	Forward()
	Home()
	MoveBy(dx, dy int)
	MoveTo(x, y int)
	Open(url, name, features string) WindowI
	OpenDialog(url, name, features string, args []any) WindowI
	PostMessage(message string, target string, transfer []any)
	Print()
	Prompt(prompt string, initial string) string
	ResizeBy(dw, dh int)
	ResizeTo(w, h int)
	Scroll(x, y int)
	ScrollBy(dx, dy int)
	ScrollByLines(int)
	ScrollTo(x, y int)
	SetCursor(name string)
	Stop()
}

type LocationI interface {
	ValueI

	Href() string
	Protocol() string
	Host() string
	Hostname() string
	Port() string
	Pathname() string
	Search() string
	Hash() string
	Username() string
	Password() string
	Origin() string
	SetHref(v string)
	SetProtocol(v string)
	SetHost(v string)
	SetHostname(v string)
	SetPort(v string)
	SetPathname(v string)
	SetSearch(v string)
	SetHash(v string)
	SetUsername(v string)
	SetPassword(v string)
}

type HistoryI interface {
	Length() int
	State() any
	Back()
	Forward()
	Go(offset int)
	PushState(state any, title string, url string)
	ReplaceState(state any, title string, url string)
}

type window struct {
	ValueI
}

var _ WindowI = &window{}

//////
//////
//////
//////

func (n *window) Underlying() ValueI {
	return n.ValueI
}

func (w *window) Document() DocumentI {
	val := w.Get("document")
	return NewDocument(val)
}

func (w *window) FrameElement() ElementI {
	val := w.Get("frameElement")
	return NewElement(val)
}

func (w *window) Location() LocationI {
	o := w.Get("location")
	return &location{ValueI: o}
}

func (w *window) Name() string {
	return w.Get("name").String()
}

func (w *window) SetName(s string) {
	w.Set("name", s)
}

func (w *window) InnerHeight() int {
	return w.Get("innerHeight").Int()
}

func (w *window) InnerWidth() int {
	return w.Get("innerWidth").Int()
}

func (w *window) Length() int {
	return w.Get("length").Int()
}

func (w *window) Opener() WindowI {
	return &window{w.Get("opener")}
}

func (w *window) OuterHeight() int {
	return w.Get("outerHeight").Int()
}

func (w *window) OuterWidth() int {
	return w.Get("outerWidth").Int()
}

func (w *window) ScrollX() int {
	return w.Get("scrollX").Int()
}

func (w *window) ScrollY() int {
	return w.Get("scrollY").Int()
}

func (w *window) Parent() WindowI {
	return &window{ValueI: w.Get("parent")}
}

func (w *window) ScreenX() int {
	return w.Get("screenX").Int()
}

func (w *window) ScreenY() int {
	return w.Get("screenY").Int()
}

func (w *window) ScrollMaxX() int {
	return w.Get("scrollMaxX").Int()
}

func (w *window) ScrollMaxY() int {
	return w.Get("scrollMaxY").Int()
}

func (w *window) Top() WindowI {
	return &window{w.Get("top")}
}

func (w *window) History() HistoryI {
	// FIXME implement
	return nil
}

func (w *window) Screen() ScreenI {
	return &screen{ValueI: w.Get("screen")}
}

func (w *window) Alert(msg string) {
	w.Call("alert", msg)
}

func (w *window) Back() {
	w.Call("back")
}

func (w *window) Blur() {
	w.Call("blur")
}

func (w *window) Close() {
	w.Call("close")
}

func (w *window) Confirm(prompt string) bool {
	return w.Call("confirm", prompt).Bool()
}

func (w *window) Focus() {
	w.Call("focus")
}

func (w *window) Forward() {
	w.Call("forward")
}

func (w *window) Home() {
	w.Call("home")
}

func (w *window) MoveBy(dx, dy int) {
	w.Call("moveBy", dx, dy)
}

func (w *window) MoveTo(x, y int) {
	w.Call("moveTo", x, y)
}

func (w *window) Open(url, name, features string) WindowI {
	return &window{w.Call("open", url, name, features)}
}

func (w *window) OpenDialog(url, name, features string, args []any) WindowI {
	return &window{w.Call("openDialog", url, name, features, args)}
}

func (w *window) PostMessage(message string, target string, transfer []any) {
	w.Call("postMessage", message, target, transfer)
}

func (w *window) Print() {
	w.Call("print")
}

func (w *window) Prompt(prompt string, initial string) string {
	return w.Call("prompt", prompt, initial).String()
}

func (w *window) ResizeBy(dw, dh int) {
	w.Call("resizeBy", dw, dh)
}

func (w *window) ResizeTo(width, height int) {
	w.Call("resizeTo", width, height)
}

func (w *window) Scroll(x, y int) {
	w.Call("scroll", x, y)
}

func (w *window) ScrollBy(dx, dy int) {
	w.Call("scrollBy", dx, dy)
}

func (w *window) ScrollByLines(i int) {
	w.Call("scrollByLines", i)
}

func (w *window) ScrollTo(x, y int) {
	w.Call("scrollTo", x, y)
}

func (w *window) SetCursor(name string) {
	w.Call("setCursor", name)
}

func (w *window) Stop() {
	w.Call("stop")
}

func (s *window) AddEventListener(typ string, useCapture bool, listener func(EventI)) EventListenerI {
	return s.Underlying().AddEventListener(typ, useCapture, listener)
}

func (s *window) RemoveEventListener(listener EventListenerI) {
	s.Underlying().RemoveEventListener(listener)
}

func (s *window) DispatchEvent(event EventI) bool {
	return s.Underlying().DispatchEvent(event)
}

type location struct {
	ValueI
}

var _ LocationI = &location{}

func (u *location) Href() string     { return u.Get("href").String() }
func (u *location) Protocol() string { return u.Get("protocol").String() }
func (u *location) Host() string     { return u.Get("host").String() }
func (u *location) Hostname() string { return u.Get("hostname").String() }
func (u *location) Port() string     { return u.Get("port").String() }
func (u *location) Pathname() string { return u.Get("pathname").String() }
func (u *location) Search() string   { return u.Get("search").String() }
func (u *location) Hash() string     { return u.Get("hash").String() }
func (u *location) Username() string { return u.Get("username").String() }
func (u *location) Password() string { return u.Get("password").String() }
func (u *location) Origin() string   { return u.Get("origin").String() }

func (u *location) SetHref(v string)     { u.Set("href", v) }
func (u *location) SetProtocol(v string) { u.Set("protocol", v) }
func (u *location) SetHost(v string)     { u.Set("host", v) }
func (u *location) SetHostname(v string) { u.Set("hostname", v) }
func (u *location) SetPort(v string)     { u.Set("port", v) }
func (u *location) SetPathname(v string) { u.Set("pathname", v) }
func (u *location) SetSearch(v string)   { u.Set("search", v) }
func (u *location) SetHash(v string)     { u.Set("hash", v) }
func (u *location) SetUsername(v string) { u.Set("username", v) }
func (u *location) SetPassword(v string) { u.Set("password", v) }
