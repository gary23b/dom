package dom

type CanvasS struct {
	ElementI
	Value ValueI
	Ctx   ValueI
}

func NewCanvas(width, height int) *CanvasS {
	e := Doc.CreateElement("canvas")
	v := e.Underlying()

	ret := &CanvasS{
		ElementI: e,
		Value:    v,
		Ctx:      v.Get("Context2d"),
	}

	v.Set("width", width)
	v.Set("height", height)

	return ret
}

func (s *CanvasS) Width() int  { return s.Value.Get("width").Int() }
func (s *CanvasS) Height() int { return s.Value.Get("height").Int() }

func (s *CanvasS) SetWidth(v int)  { s.Value.Set("width", v) }
func (s *CanvasS) SetHeight(v int) { s.Value.Set("height", v) }

func (s *CanvasS) FillStyle() string   { return s.Ctx.Get("fillStyle").String() }
func (s *CanvasS) StrokeStyle() string { return s.Ctx.Get("strokeStyle").String() }
func (s *CanvasS) ShadowColor() string { return s.Ctx.Get("shadowColor").String() }
func (s *CanvasS) ShadowBlur() int     { return s.Ctx.Get("shadowBlur").Int() }
func (s *CanvasS) ShadowOffsetX() int  { return s.Ctx.Get("shadowOffsetX").Int() }
func (s *CanvasS) ShadowOffsetY() int  { return s.Ctx.Get("shadowOffsetY").Int() }

func (s *CanvasS) SetFillStyle(v string)   { s.Ctx.Set("fillStyle", v) }
func (s *CanvasS) SetStrokeStyle(v string) { s.Ctx.Set("strokeStyle", v) }
func (s *CanvasS) SetShadowColor(v string) { s.Ctx.Set("shadowColor", v) }
func (s *CanvasS) SetShadowBlur(v int)     { s.Ctx.Set("shadowBlur", v) }
func (s *CanvasS) SetShadowOffsetX(v int)  { s.Ctx.Set("shadowOffsetX", v) }
func (s *CanvasS) SetShadowOffsetY(v int)  { s.Ctx.Set("shadowOffsetY", v) }

// Line Styles

func (s *CanvasS) LineCap() string  { return s.Ctx.Get("lineCap").String() }
func (s *CanvasS) LineJoin() string { return s.Ctx.Get("lineJoin").String() }
func (s *CanvasS) LineWidth() int   { return s.Ctx.Get("lineWidth").Int() }
func (s *CanvasS) MiterLimit() int  { return s.Ctx.Get("miterLimit").Int() }

func (s *CanvasS) SetLineCap(v string)  { s.Ctx.Set("lineCap", v) }
func (s *CanvasS) SetLineJoin(v string) { s.Ctx.Set("lineJoin", v) }
func (s *CanvasS) SetLineWidth(v int)   { s.Ctx.Set("lineWidth", v) }
func (s *CanvasS) SetMiterLimit(v int)  { s.Ctx.Set("miterLimit", v) }

// Text

func (s *CanvasS) Font() string         { return s.Ctx.Get("font").String() }
func (s *CanvasS) TextAlign() string    { return s.Ctx.Get("textAlign").String() }
func (s *CanvasS) TextBaseline() string { return s.Ctx.Get("textBaseline").String() }

func (s *CanvasS) SetFont(v string)         { s.Ctx.Set("font", v) }
func (s *CanvasS) SetTextAlign(v string)    { s.Ctx.Set("textAlign", v) }
func (s *CanvasS) SetTextBaseline(v string) { s.Ctx.Set("textBaseline", v) }

// Compositing

func (s *CanvasS) GlobalAlpha() float64 { return s.Ctx.Get("globalAlpha").Float() }
func (s *CanvasS) GlobalCompositeOperation() string {
	return s.Ctx.Get("globalCompositeOperation").String()
}

func (s *CanvasS) SetGlobalAlpha(v float64) { s.Ctx.Set("globalAlpha", v) }
func (s *CanvasS) SetGlobalCompositeOperation(v string) {
	s.Ctx.Set("globalCompositeOperation", v)
}

// Drawing Rectangles

func (s *CanvasS) ClearRect(x, y, width, height float64) {
	s.Ctx.Call("clearRect", x, y, width, height)
}

func (s *CanvasS) FillRect(x, y, width, height float64) {
	s.Ctx.Call("fillRect", x, y, width, height)
}

func (s *CanvasS) StrokeRect(x, y, width, height float64) {
	s.Ctx.Call("strokeRect", x, y, width, height)
}

// Drawing Text

// FillText fills a given text at the given (x, y) position.
// If the optional maxWidth parameter is not -1,
// the text will be scaled to fit that width.
func (s *CanvasS) FillText(text string, x, y, maxWidth float64) {
	if maxWidth == -1 {
		s.Ctx.Call("fillText", text, x, y)
		return
	}

	s.Ctx.Call("fillText", text, x, y, maxWidth)
}

// StrokeText strokes a given text at the given (x, y) position.
// If the optional maxWidth parameter is not -1,
// the text will be scaled to fit that width.
func (s *CanvasS) StrokeText(text string, x, y, maxWidth float64) {
	if maxWidth == -1 {
		s.Ctx.Call("strokeText", text, x, y)
		return
	}

	s.Ctx.Call("strokeText", text, x, y, maxWidth)
}

// func (s *CanvasS) MeasureText(text string) *TextMetrics {
// 	textMetrics := s.Ctx.Call("measureText", text)
// 	return &TextMetrics{Value: textMetrics}
// }

// Line styles

func (s *CanvasS) GetLineDash() []float64 {
	var dashes []float64
	lineDash := s.Ctx.Call("getLineDash")
	for i := 0; i < lineDash.Length(); i++ {
		dashes = append(dashes, lineDash.Index(i).Float())
	}
	return dashes
}

func (s *CanvasS) SetLineDash(dashes []float64) {
	s.Ctx.Call("setLineDash", dashes)
}

// Gradients and patterns

// CreateLinearGradient creates a linear gradient along the line given
// by the coordinates represented by the parameters.
//
// Reference: https://developer.mozilla.org/en-US/docs/Web/API/CanvasRenderingContext2D/createLinearGradient.
// func (s *CanvasS) CreateLinearGradient(x0, y0, x1, y1 float64) *CanvasGradient {
// 	return &CanvasGradient{Value: s.Ctx.Call("createLinearGradient", x0, y0, x1, y1)}
// }

// // CreateRadialGradient creates a radial gradient given by the coordinates of the two circles
// // represented by the parameters.
// //
// // Reference: https://developer.mozilla.org/en-US/docs/Web/API/CanvasRenderingContext2D/createRadialGradient.
// func (s *CanvasS) CreateRadialGradient(x0, y0, r0, x1, y1, r1 float64) *CanvasGradient {
// 	return &CanvasGradient{Value: s.Ctx.Call("createRadialGradient", x0, y0, r0, x1, y1, r1)}
// }

// // CreatePattern creates a pattern using the specified image (a CanvasImageSource).
// // It repeats the source in the directions specified by the repetition argument.
// //
// // Reference: https://developer.mozilla.org/en-US/docs/Web/API/CanvasRenderingContext2D/createPattern.
// func (s *CanvasS) CreatePattern(image ElementI, repetition string) *CanvasPattern {
// 	return &CanvasPattern{Value: s.Ctx.Call("createPattern", image.Underlying(), repetition)}
// }

// Paths

func (s *CanvasS) BeginPath() {
	s.Ctx.Call("beginPath")
}

func (s *CanvasS) ClosePath() {
	s.Ctx.Call("closePath")
}

func (s *CanvasS) MoveTo(x, y float64) {
	s.Ctx.Call("moveTo", x, y)
}

func (s *CanvasS) LineTo(x, y float64) {
	s.Ctx.Call("lineTo", x, y)
}

func (s *CanvasS) BezierCurveTo(cp1x, cp1y, cp2x, cp2y, x, y float64) {
	s.Ctx.Call("bezierCurveTo", cp1x, cp1y, cp2x, cp2y, x, y)
}

func (s *CanvasS) QuadraticCurveTo(cpx, cpy, x, y float64) {
	s.Ctx.Call("quadraticCurveTo", cpx, cpy, x, y)
}

func (s *CanvasS) Arc(x, y, r, sAngle, eAngle float64, counterclockwise bool) {
	s.Ctx.Call("arc", x, y, r, sAngle, eAngle, counterclockwise)
}

func (s *CanvasS) ArcTo(x1, y1, x2, y2, r float64) {
	s.Ctx.Call("arcTo", x1, y1, x2, y2, r)
}

func (s *CanvasS) Ellipse(x, y, radiusX, radiusY, rotation, startAngle, endAngle float64, anticlockwise bool) {
	s.Ctx.Call("ellipse", x, y, radiusX, radiusY, rotation, startAngle, endAngle, anticlockwise)
}

func (s *CanvasS) Rect(x, y, width, height float64) {
	s.Ctx.Call("rect", x, y, width, height)
}

// Drawing paths

func (s *CanvasS) Fill() {
	s.Ctx.Call("fill")
}

func (s *CanvasS) Stroke() {
	s.Ctx.Call("stroke")
}

func (s *CanvasS) DrawFocusIfNeeded(element ElementI, path ValueI) {
	s.Ctx.Call("drawFocusIfNeeded", element, path)
}

func (s *CanvasS) ScrollPathIntoView(path ValueI) {
	s.Ctx.Call("scrollPathIntoView", path)
}

func (s *CanvasS) Clip() {
	s.Ctx.Call("clip")
}

func (s *CanvasS) IsPointInPath(x, y float64) bool {
	return s.Ctx.Call("isPointInPath", x, y).Bool()
}

func (s *CanvasS) IsPointInStroke(path ValueI, x, y float64) bool {
	return s.Ctx.Call("isPointInStroke", path, x, y).Bool()
}

// Transformations

func (s *CanvasS) Rotate(angle float64) {
	s.Ctx.Call("rotate", angle)
}

func (s *CanvasS) Scale(scaleWidth, scaleHeight float64) {
	s.Ctx.Call("scale", scaleWidth, scaleHeight)
}

func (s *CanvasS) Translate(x, y float64) {
	s.Ctx.Call("translate", x, y)
}

func (s *CanvasS) Transform(a, b, c, d, e, f float64) {
	s.Ctx.Call("transform", a, b, c, d, e, f)
}

func (s *CanvasS) SetTransform(a, b, c, d, e, f float64) {
	s.Ctx.Call("setTransform", a, b, c, d, e, f)
}

func (s *CanvasS) ResetTransform() {
	s.Ctx.Call("resetTransform")
}

// Drawing images

func (s *CanvasS) DrawImage(image ElementI, dx, dy float64) {
	s.Ctx.Call("drawImage", image.Underlying(), dx, dy)
}

func (s *CanvasS) DrawImageWithDst(image ElementI, dx, dy, dWidth, dHeight float64) {
	s.Ctx.Call("drawImage", image.Underlying(), dx, dy, dWidth, dHeight)
}

func (s *CanvasS) DrawImageWithSrcAndDst(image ElementI, sx, sy, sWidth, sHeight, dx, dy, dWidth, dHeight float64) {
	s.Ctx.Call("drawImage", image.Underlying(), sx, sy, sWidth, sHeight, dx, dy, dWidth, dHeight)
}

// Pixel manipulation

// func (s *CanvasS) CreateImageData(width, height int) *ImageData {
// 	return &ImageData{Value: s.Ctx.Call("createImageData", width, height)}
// }

// func (s *CanvasS) GetImageData(sx, sy, sw, sh int) *ImageData {
// 	return &ImageData{Value: s.Ctx.Call("getImageData", sx, sy, sw, sh)}
// }

// func (s *CanvasS) PutImageData(imageData *ImageData, dx, dy float64) {
// 	s.Ctx.Call("putImageData", imageData.Value, dx, dy)
// }

// func (s *CanvasS) PutImageDataDirty(imageData *ImageData, dx, dy float64, dirtyX, dirtyY, dirtyWidth, dirtyHeight int) {
// 	s.Ctx.Call("putImageData", imageData.Value, dx, dy, dirtyX, dirtyY, dirtyWidth, dirtyHeight)
// }

// State

func (s *CanvasS) Save() {
	s.Ctx.Call("save")
}

func (s *CanvasS) Restore() {
	s.Ctx.Call("restore")
}
