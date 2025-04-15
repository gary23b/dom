package dom

type TokenListI interface {
	Length() int
	Item(idx int) string
	Contains(token string) bool
	Add(token string)
	Remove(token string)
	Toggle(token string)
	String() string
	Slice() []string
	// SetString(s string)
	// Set(s []string)
}

func NewTokenList(val ValueI) tokenListS {
	ret := tokenListS{
		ValueI: val,
	}
	return ret
}

type tokenListS struct {
	ValueI
}

var _ TokenListI = tokenListS{}

func (s tokenListS) Length() int { return s.Get("length").Int() }

func (s tokenListS) Item(idx int) string {
	return s.Call("item", idx).String()
}

func (s tokenListS) Contains(token string) bool {
	return s.Call("contains", token).Bool()
}

func (s tokenListS) Add(token string) {
	s.Call("add", token)
}

func (s tokenListS) Remove(token string) {
	s.Call("remove", token)
}

func (s tokenListS) Toggle(token string) {
	s.Call("toggle", token)
}

func (s tokenListS) Slice() []string {
	var out []string
	length := s.Get("length").Int()
	for i := 0; i < length; i++ {
		out = append(out, s.Call("item", i).String())
	}
	return out
}
