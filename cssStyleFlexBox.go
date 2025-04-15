package dom

import (
	"fmt"
	"log"
)

type CSSStyleFlexBoxI interface {
	DisplayFlex() CSSStyleFlexBoxI
	FlexDirection(direction FlexBoxFlexDirection) CSSStyleFlexBoxI
	FlexWrap(wrap FlexBoxFlexWrap) CSSStyleFlexBoxI
	JustifyContent(in FlexBoxJustifyContent) CSSStyleFlexBoxI
	AlignItems(in FlexBoxJustifyContent) CSSStyleFlexBoxI
	AlignContent(in FlexBoxAlignContent) CSSStyleFlexBoxI
	Order(in int) CSSStyleFlexBoxI
	OrderDefault(in FlexBoxOrder) CSSStyleFlexBoxI
	FlexGrow(in float64) CSSStyleFlexBoxI
	FlexGrowDefault(in FlexBoxFlexGrow) CSSStyleFlexBoxI
	FlexShrink(in float64) CSSStyleFlexBoxI
	FlexShrinkDefault(in FlexBoxFlexShrink) CSSStyleFlexBoxI
	FlexFlexBasis(in string) CSSStyleFlexBoxI
	FlexSFlexBasisDefault(in FlexBoxFlexBasis) CSSStyleFlexBoxI
	AlignSelf(in FlexBoxAlignSelf) CSSStyleFlexBoxI
}

func NewCssStyleFlexBox(val ValueI) cssStyleFlexBoxS {
	ret := cssStyleFlexBoxS{
		ValueI: val,
	}
	return ret
}

type cssStyleFlexBoxS struct {
	ValueI
}

var _ CSSStyleFlexBoxI = cssStyleFlexBoxS{}

/*
https://www.w3schools.com/css/css3_flexbox.asp
*/

/*
https://www.w3schools.com/jsref/prop_style_display.asp
*/
func (s cssStyleFlexBoxS) DisplayFlex() CSSStyleFlexBoxI {
	s.Set("display", "flex")
	return s
}

/*
https://www.w3schools.com/jsref/prop_style_flexdirection.asp
*/
type FlexBoxFlexDirection string

const (
	FlexBoxFlexDirection_Row           FlexBoxFlexDirection = "row"
	FlexBoxFlexDirection_Column        FlexBoxFlexDirection = "column"
	FlexBoxFlexDirection_RowReverse    FlexBoxFlexDirection = "row-reverse"
	FlexBoxFlexDirection_ColumnReverse FlexBoxFlexDirection = "column-reverse"
	FlexBoxFlexDirection_Initial       FlexBoxFlexDirection = "initial"
	FlexBoxFlexDirection_Inherit       FlexBoxFlexDirection = "inherit"
)

func (s cssStyleFlexBoxS) FlexDirection(direction FlexBoxFlexDirection) CSSStyleFlexBoxI {
	s.Set("flexDirection", string(direction))
	return s
}

/*
https://www.w3schools.com/jsref/prop_style_flexwrap.asp
*/
type FlexBoxFlexWrap string

const (
	FlexBoxFlexWrap_NoWrap      FlexBoxFlexWrap = "nowrap"
	FlexBoxFlexWrap_Wrap        FlexBoxFlexWrap = "wrap"
	FlexBoxFlexWrap_WrapReverse FlexBoxFlexWrap = "wrap-reverse"
	FlexBoxFlexWrap_Initial     FlexBoxFlexWrap = "initial"
	FlexBoxFlexWrap_Inherit     FlexBoxFlexWrap = "inherit"
)

func (s cssStyleFlexBoxS) FlexWrap(wrap FlexBoxFlexWrap) CSSStyleFlexBoxI {
	s.Set("flexWrap", string(wrap))
	return s
}

/*
https://www.w3schools.com/jsref/prop_style_justifycontent.asp
*/
type FlexBoxJustifyContent string

const (
	FlexBoxJustifyContent_FlexStart    FlexBoxJustifyContent = "flex-start"
	FlexBoxJustifyContent_FlexEnd      FlexBoxJustifyContent = "flex-end"
	FlexBoxJustifyContent_Center       FlexBoxJustifyContent = "center"
	FlexBoxJustifyContent_SpaceBetween FlexBoxJustifyContent = "space-between"
	FlexBoxJustifyContent_SpaceAround  FlexBoxJustifyContent = "space-around"
	FlexBoxJustifyContent_SpaceEvenly  FlexBoxJustifyContent = "space-evenly"
	FlexBoxJustifyContent_Initial      FlexBoxJustifyContent = "initial"
	FlexBoxJustifyContent_Inherit      FlexBoxJustifyContent = "inherit"
)

// This is the left right alignment fo the content
func (s cssStyleFlexBoxS) JustifyContent(in FlexBoxJustifyContent) CSSStyleFlexBoxI {
	s.Set("justifyContent", string(in))
	return s
}

/*
https://www.w3schools.com/jsref/prop_style_alignitems.asp
*/
type FlexBoxAlignItems string

const (
	FlexBoxAlignItems_Stretch   FlexBoxAlignItems = "stretch"
	FlexBoxAlignItems_Center    FlexBoxAlignItems = "center"
	FlexBoxAlignItems_FlexStart FlexBoxAlignItems = "flex-start"
	FlexBoxAlignItems_FlexEnd   FlexBoxAlignItems = "flex-end"
	FlexBoxAlignItems_BaseLine  FlexBoxAlignItems = "baseline"
	FlexBoxAlignItems_Initial   FlexBoxAlignItems = "initial"
	FlexBoxAlignItems_Inherit   FlexBoxAlignItems = "inherit"
)

// This is the up/down alignment of the content inside the flex items
// property is used to align the flex items when they do not use all available space on the cross-axis (vertically).
func (s cssStyleFlexBoxS) AlignItems(in FlexBoxJustifyContent) CSSStyleFlexBoxI {
	s.Set("alignItems", string(in))
	return s
}

/*
https://www.w3schools.com/jsref/prop_style_alignitems.asp
*/
type FlexBoxAlignContent string

const (
	FlexBoxAlignContent_Stretch      FlexBoxAlignContent = "stretch"
	FlexBoxAlignContent_Center       FlexBoxAlignContent = "center"
	FlexBoxAlignContent_FlexStart    FlexBoxAlignContent = "flex-start"
	FlexBoxAlignContent_FlexEnd      FlexBoxAlignContent = "flex-end"
	FlexBoxAlignContent_SpaceBetween FlexBoxAlignContent = "space-between"
	FlexBoxAlignContent_SpaceAround  FlexBoxAlignContent = "space-around"
	FlexBoxAlignContent_SpaceEvenly  FlexBoxAlignContent = "space-evenly"
	FlexBoxAlignContent_Initial      FlexBoxAlignContent = "initial"
	FlexBoxAlignContent_Inherit      FlexBoxAlignContent = "inherit"
)

// This is the up/down alignment of the boxes
func (s cssStyleFlexBoxS) AlignContent(in FlexBoxAlignContent) CSSStyleFlexBoxI {
	s.Set("alignContent", string(in))
	return s
}

/*
https://www.w3schools.com/jsref/prop_style_order.asp
*/
type FlexBoxOrder string

const (
	FlexBoxOrder_Initial FlexBoxOrder = "initial"
	FlexBoxOrder_Inherit FlexBoxOrder = "inherit"
)

func (s cssStyleFlexBoxS) Order(in int) CSSStyleFlexBoxI {
	if in < 0 {
		log.Println("Order must be set with positive int")
		return s
	}
	s.Set("order", fmt.Sprint(in))
	return s
}
func (s cssStyleFlexBoxS) OrderDefault(in FlexBoxOrder) CSSStyleFlexBoxI {
	s.Set("order", string(in))
	return s
}

/*
https://www.w3schools.com/jsref/prop_style_flexgrow.asp
*/
type FlexBoxFlexGrow string

const (
	FlexBoxFlexGrow_Initial FlexBoxFlexGrow = "initial"
	FlexBoxFlexGrow_Inherit FlexBoxFlexGrow = "inherit"
)

func (s cssStyleFlexBoxS) FlexGrow(in float64) CSSStyleFlexBoxI {
	if in < 0 {
		log.Println("Order must be set with positive int")
		return s
	}
	s.Set("flexGrow", fmt.Sprintf("%.6f", in))
	return s
}
func (s cssStyleFlexBoxS) FlexGrowDefault(in FlexBoxFlexGrow) CSSStyleFlexBoxI {
	s.Set("flexGrow", string(in))
	return s
}

/*
https://www.w3schools.com/jsref/prop_style_flexshrink.asp
*/
type FlexBoxFlexShrink string

const (
	FlexBoxFlexShrink_Initial FlexBoxFlexShrink = "initial"
	FlexBoxFlexShrink_Inherit FlexBoxFlexShrink = "inherit"
)

func (s cssStyleFlexBoxS) FlexShrink(in float64) CSSStyleFlexBoxI {
	if in < 0 {
		log.Println("Order must be set with positive int")
		return s
	}
	s.Set("flexShrink", fmt.Sprintf("%.6f", in))
	return s
}
func (s cssStyleFlexBoxS) FlexShrinkDefault(in FlexBoxFlexShrink) CSSStyleFlexBoxI {
	s.Set("flexShrink", string(in))
	return s
}

/*
https://www.w3schools.com/jsref/prop_style_flexshrink.asp
*/
type FlexBoxFlexBasis string

const (
	FlexBoxFlexBasis_Auto    FlexBoxFlexBasis = "auto"
	FlexBoxFlexBasis_Initial FlexBoxFlexBasis = "initial"
	FlexBoxFlexBasis_Inherit FlexBoxFlexBasis = "inherit"
)

func (s cssStyleFlexBoxS) FlexFlexBasis(in string) CSSStyleFlexBoxI {
	s.Set("flexBasis", in)
	return s
}
func (s cssStyleFlexBoxS) FlexSFlexBasisDefault(in FlexBoxFlexBasis) CSSStyleFlexBoxI {
	s.Set("flexBasis", string(in))
	return s
}

/*
https://www.w3schools.com/jsref/prop_style_alignself.asp
*/
type FlexBoxAlignSelf string

const (
	FlexBoxAlignSelf_Stretch   FlexBoxAlignSelf = "stretch"
	FlexBoxAlignSelf_Center    FlexBoxAlignSelf = "center"
	FlexBoxAlignSelf_FlexStart FlexBoxAlignSelf = "flex-start"
	FlexBoxAlignSelf_FlexEnd   FlexBoxAlignSelf = "flex-end"
	FlexBoxAlignSelf_BaseLine  FlexBoxAlignSelf = "baseline"
	FlexBoxAlignSelf_auto      FlexBoxAlignSelf = "auto"
	FlexBoxAlignSelf_Initial   FlexBoxAlignSelf = "initial"
	FlexBoxAlignSelf_Inherit   FlexBoxAlignSelf = "inherit"
)

func (s cssStyleFlexBoxS) AlignSelf(in FlexBoxAlignSelf) CSSStyleFlexBoxI {
	s.Set("alignSelf", string(in))
	return s
}
