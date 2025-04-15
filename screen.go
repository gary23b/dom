package dom

type ScreenI interface {
	AvailTop() int
	AvailLeft() int
	AvailHeight() int
	AvailWidth() int
	ColorDepth() int
	Height() int
	Left() int
	// Orientation() ScreenOrientationI
	PixelDepth() int
	Top() int
	Width() int
}

// type ScreenOrientationI interface {
// 	AddEventListener(typ string, useCapture bool, listener func(EventT)) js.Func
// 	RemoveEventListener(typ string, useCapture bool, listener js.Func)
// 	DispatchEvent(event EventT) bool
// }

type screen struct {
	ValueI
}

var _ ScreenI = &screen{}

func (s *screen) AvailTop() int {
	return s.Get("availTop").Int()
}

func (s *screen) AvailLeft() int {
	return s.Get("availLeft").Int()
}

func (s *screen) AvailHeight() int {
	return s.Get("availHeight").Int()
}

func (s *screen) AvailWidth() int {
	return s.Get("availWidth").Int()
}

func (s *screen) ColorDepth() int {
	return s.Get("colorDepth").Int()
}

func (s *screen) Height() int {
	return s.Get("height").Int()
}

func (s *screen) Left() int {
	return s.Get("left").Int()
}

func (s *screen) PixelDepth() int {
	return s.Get("pixelDepth").Int()
}

func (s *screen) Top() int {
	return s.Get("top").Int()
}

func (s *screen) Width() int {
	return s.Get("width").Int()
}
