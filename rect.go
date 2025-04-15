package dom

type RectI interface {
	X() float64
	Y() float64
	Width() float64
	Height() float64
	Top() float64
	Right() float64
	Bottom() float64
	Left() float64
	SetX(v float64)
	SetY(v float64)
	SetWidth(v float64)
	SetHeight(v float64)
	SetTop(v float64)
	SetRight(v float64)
	SetBottom(v float64)
	SetLeft(v float64)
}

func NewRect(val ValueI) rectS {
	ret := rectS{
		ValueI: val,
	}
	return ret
}

type rectS struct {
	ValueI
}

var _ RectI = rectS{}

func (s rectS) X() float64      { return s.Get("x").Float() }
func (s rectS) Y() float64      { return s.Get("y").Float() }
func (s rectS) Width() float64  { return s.Get("width").Float() }
func (s rectS) Height() float64 { return s.Get("height").Float() }
func (s rectS) Top() float64    { return s.Get("top").Float() }
func (s rectS) Right() float64  { return s.Get("right").Float() }
func (s rectS) Bottom() float64 { return s.Get("bottom").Float() }
func (s rectS) Left() float64   { return s.Get("left").Float() }

func (s rectS) SetX(v float64)      { s.Set("x", v) }
func (s rectS) SetY(v float64)      { s.Set("y", v) }
func (s rectS) SetWidth(v float64)  { s.Set("width", v) }
func (s rectS) SetHeight(v float64) { s.Set("height", v) }
func (s rectS) SetTop(v float64)    { s.Set("top", v) }
func (s rectS) SetRight(v float64)  { s.Set("right", v) }
func (s rectS) SetBottom(v float64) { s.Set("bottom", v) }
func (s rectS) SetLeft(v float64)   { s.Set("left", v) }
