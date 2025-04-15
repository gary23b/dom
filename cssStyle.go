package dom

import (
	"fmt"
	"image/color"
)

type CSSStyleI interface {
	ToMap() map[string]string
	RemoveProperty(name string)
	GetPropertyValue(name string) string
	GetPropertyPriority(name string) string
	SetProperty(name, value, priority string)
	Index(idx int) string
	Length() int

	//
	//
	Set(name, value string) CSSStyleI
	BackgroundColorStr(c string) CSSStyleI
	BackgroundColor(c color.Color) CSSStyleI
	ColorStr(c string) CSSStyleI
	Color(c color.Color) CSSStyleI
	TextAlign(in TextAlign) CSSStyleI
	Padding(in string) CSSStyleI
	FlexBox() CSSStyleFlexBoxI
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

// //
// //
func (s cssStyleS) Set(name, value string) CSSStyleI {
	s.ValueI.Set(name, value)
	return s
}

/*
https://www.w3schools.com/jsref/prop_style_backgroundcolor.asp
transparent
initial
inherit
*/
func (s cssStyleS) BackgroundColorStr(c string) CSSStyleI {
	s.Set("backgroundColor", c)
	return s
}

func (s cssStyleS) BackgroundColor(c color.Color) CSSStyleI {
	rgba, _ := color.RGBAModel.Convert(c).(color.RGBA)
	rbgColorString := fmt.Sprintf("#%02x%02x%02x%02x", rgba.R, rgba.G, rgba.B, rgba.A)
	return s.BackgroundColorStr(rbgColorString)
}

/*
https://www.w3schools.com/jsref/prop_style_color.asp
initial
inherit
*/
func (s cssStyleS) ColorStr(c string) CSSStyleI {
	s.Set("color", c)
	return s
}

func (s cssStyleS) Color(c color.Color) CSSStyleI {
	rgba, _ := color.RGBAModel.Convert(c).(color.RGBA)
	rbgColorString := fmt.Sprintf("#%02x%02x%02x%02x", rgba.R, rgba.G, rgba.B, rgba.A)
	return s.ColorStr(rbgColorString)
}

/*
https://www.w3schools.com/jsref/prop_style_textalign.asp
*/
type TextAlign string

const (
	TextAlign_Left    TextAlign = "left"
	TextAlign_Right   TextAlign = "right"
	TextAlign_Center  TextAlign = "center"
	TextAlign_Justify TextAlign = "justify"
	TextAlign_Initial TextAlign = "initial"
	TextAlign_Inherit TextAlign = "inherit"
)

func (s cssStyleS) TextAlign(in TextAlign) CSSStyleI {
	s.Set("textAlign", string(in))
	return s
}

/*
https://www.w3schools.com/jsref/prop_style_padding.asp
*/
func (s cssStyleS) Padding(in string) CSSStyleI {
	s.Set("padding", in)
	return s
}

func (s cssStyleS) FlexBox() CSSStyleFlexBoxI {
	ret := NewCssStyleFlexBox(s.ValueI)
	ret.DisplayFlex()
	return ret
}
