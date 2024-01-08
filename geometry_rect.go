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

// IsContainsPoint tests if the Rect contains the Point.
func (r Rect) IsContainsPoint(c Point) bool {
	return c.X >= r.Xf && c.X <= r.Xt && c.Y >= r.Yf && c.Y <= r.Yt
}

// IsContainsRectangle tests if the Rect contains the other Rect.
func (r Rect) IsContainsRectangle(t Rect) bool {
	return r.Xf <= t.Xf && r.Xt >= t.Xt && r.Yf <= t.Yf && r.Yt >= t.Yt
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
