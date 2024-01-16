package main

import (
	"fmt"
	"math"
)

// Rect represents a rectangle in a 2D plane.
type Rect struct {
	Xf, Xt, Yf, Yt float64
}

// IsSame tests if two Rects are the same.
func (r Rect) IsSame(t Rect) bool {
	return r.Xf == t.Xf && r.Xt == t.Xt && r.Yf == t.Yf && r.Yt == t.Yt
}

// Width returns the width of the Rect.
func (r Rect) Width() float64 {
	return r.Xt - r.Xf
}

// Height returns the height of the Rect.
func (r Rect) Height() float64 {
	return r.Yt - r.Yf
}

// Area returns the area of the Rect.
func (r Rect) Area() float64 {
	return r.Width() * r.Height()
}

// Center returns the center Point of the Rect.
func (r Rect) Center() Point {
	return NewPoint(
		math.Round(r.Xf+r.Width()/2),
		math.Round(r.Yf+r.Height()/2),
	)
}

// Symmetric returns the symmetric Rect according to the given width and height.
func (r Rect) Symmetric(width, height float64) Rect {
	return Rect{
		Xf: width - r.Xf,
		Xt: width - r.Xt,
		Yf: height - r.Yf,
		Yt: height - r.Yt,
	}
}

// IsContainsPoint tests if the Rect contains the Point.
func (r Rect) IsContainsPoint(t Point) bool {
	return t.X >= r.Xf && t.X <= r.Xt && t.Y >= r.Yf && t.Y <= r.Yt
}

// IsContainsRectangle tests if the Rect contains the other Rect.
func (r Rect) IsContainsRectangle(t Rect) bool {
	return r.Xf <= t.Xf && r.Xt >= t.Xt && r.Yf <= t.Yf && r.Yt >= t.Yt
}

// IsIntersectsRect tests if the Rect intersects the other Rect.
func (r Rect) IsIntersectsRect(t Rect) bool {
	return !(r.Xt < t.Xf || r.Xf > t.Xt || r.Yt < t.Yf || r.Yf > t.Yt)
}

// RectsIntersection returns the intersection Rect of two Rects.
func (r Rect) RectsIntersection(t Rect) (Rect, bool) {
	if !r.IsIntersectsRect(t) {
		return Rect{}, false
	}

	ir := NewRectangle(
		math.Max(r.Xf, t.Xf),
		math.Min(r.Xt, t.Xt),
		math.Max(r.Yf, t.Yf),
		math.Min(r.Yt, t.Yt),
	)
	if ir.Width() == 0 || ir.Height() == 0 {
		return Rect{}, false
	}

	return ir, true
}

// Vertices returns the Points vertices of the Rect.
func (r Rect) Vertices() Points {
	return Points{
		topLeft0:     {r.Xf, r.Yt},
		topRight0:    {r.Xt, r.Yt},
		bottomRight0: {r.Xt, r.Yf},
		bottomLeft0:  {r.Xf, r.Yf},
	}
}

// Edges returns the Lines edges of the Rect.
func (r Rect) Edges() Lines {
	return Lines{
		top:    {Point{r.Xf, r.Yt}, Point{r.Xt, r.Yt}},
		right:  {Point{r.Xt, r.Yf}, Point{r.Xt, r.Yt}},
		bottom: {Point{r.Xf, r.Yf}, Point{r.Xt, r.Yf}},
		left:   {Point{r.Xf, r.Yf}, Point{r.Xf, r.Yt}},
	}
}

func (r Rect) String() string {
	return fmt.Sprintf("[X:%.f>%.f,Y:%.f>%.f]", r.Xf, r.Xt, r.Yf, r.Yt)
}

func NewRectangle(xf, xt, yf, yt float64) Rect {
	if xf > xt {
		xf, xt = xt, xf
	}
	if yf > yt {
		yf, yt = yt, yf
	}

	return Rect{
		Xf: xf,
		Xt: xt,
		Yf: yf,
		Yt: yt,
	}
}

type Rects []Rect
