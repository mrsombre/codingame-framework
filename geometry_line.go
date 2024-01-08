package main

import (
	"fmt"
	"math"
)

// Line represents a line in a 2D plane.
type Line struct {
	From, To Point
}

// IsSame tests if two Lines segments are the same.
func (ln Line) IsSame(t Line) bool {
	return ln.From.IsSame(t.From) && ln.To.IsSame(t.To)
}

// Length returns the length of the Line.
func (ln Line) Length() float64 {
	return ln.From.Distance(ln.To)
}

// Segment returns the Line segment with the given length.
func (ln Line) Segment(length float64) Line {
	if length == 0 {
		return Line{ln.From, ln.From}
	}

	v := ln.Direction()
	vl := ln.Length()
	ux := v.X / vl
	uy := v.Y / vl

	return Line{
		ln.From,
		NewPoint(
			ln.From.X+ux*length,
			ln.From.Y+uy*length,
		),
	}
}

// IsMoving tests if the Line represents a moving object.
func (ln Line) IsMoving() bool {
	return !ln.From.IsSame(ln.To)
}

// Direction returns the direction of the Line.
func (ln Line) Direction() Point {
	return ln.To.Sub(ln.From)
}

// IsHorizontal tests if the Line is horizontal.
func (ln Line) IsHorizontal() bool {
	return ln.From.Y == ln.To.Y
}

// IsVertical tests if the Line is vertical.
func (ln Line) IsVertical() bool {
	return ln.From.X == ln.To.X
}

// Slope returns the slope value of the Line.
func (ln Line) Slope() float64 {
	return (ln.To.Y - ln.From.Y) / (ln.To.X - ln.From.X)
}

// IsPointOnLine tests if the Point is on the Line.
func (ln Line) IsPointOnLine(p Point) bool {
	if ln.IsHorizontal() {
		return ln.From.Y == p.Y
	}
	if ln.IsVertical() {
		return ln.From.X == p.X
	}

	// diagonal
	slope := ln.Slope()
	yIntercept := ln.From.Y - slope*ln.From.X
	expectedY := slope*p.X + yIntercept

	return math.Round(p.Y) == math.Round(expectedY)
}

// IsPointOnSegment tests if the Point is on the Line segment.
func (ln Line) IsPointOnSegment(p Point) bool {
	if !ln.Rect().IsContainsPoint(p) {
		return false
	}
	return ln.IsPointOnLine(p)
}

// LinesIntersection returns the crossing Point of two Lines.
func (ln Line) LinesIntersection(tl Line) (Point, bool) {
	// parallel
	if (ln.IsHorizontal() && tl.IsHorizontal()) || (ln.IsVertical() && tl.IsVertical()) {
		return Point{}, false
	}

	// vertical / horizontal
	if ln.IsVertical() && tl.IsHorizontal() {
		return tl.LinesIntersection(ln)
	}
	if ln.IsHorizontal() && tl.IsVertical() {
		return NewPoint(tl.From.X, ln.From.Y), true
	}

	// diagonal and vertical
	if tl.IsVertical() {
		return tl.LinesIntersection(ln)
	}
	if ln.IsVertical() {
		x := ln.From.X
		y := tl.Slope()*x + tl.From.Y
		return NewPoint(x, y), true
	}

	// diagonal
	slope1 := ln.Slope()
	slope2 := tl.Slope()
	// parallel
	if slope1 == slope2 {
		return Point{}, false
	}

	yIntercept1 := ln.From.Y - slope1*ln.From.X
	yIntercept2 := tl.From.Y - slope2*tl.From.X

	x := (yIntercept2 - yIntercept1) / (slope1 - slope2)
	y := slope1*x + yIntercept1

	return NewPoint(x, y), true
}

// SegmentsIntersection returns the crossing Point of two Line segments.
func (ln Line) SegmentsIntersection(tl Line) (Point, bool) {
	p, ok := ln.LinesIntersection(tl)
	if !ok {
		return Point{}, false
	}
	if !ln.Rect().IsContainsPoint(p) || !tl.Rect().IsContainsPoint(p) {
		return Point{}, false
	}
	return p, true
}

// LineSegmentIntersection returns the crossing Point of the Line and the Line segment.
func (ln Line) LineSegmentIntersection(tl Line) (Point, bool) {
	p, ok := ln.LinesIntersection(tl)
	if !ok {
		return Point{}, false
	}
	if !tl.Rect().IsContainsPoint(p) {
		return Point{}, false
	}
	return p, true
}

// Rotate returns the Line rotated by the given angle.
func (ln Line) Rotate(angle float64) Line {
	radians := angle * math.Pi / 180

	v := ln.Direction()
	rx := v.X*math.Cos(radians) + v.Y*math.Sin(radians)
	ry := v.X*-math.Sin(radians) + v.Y*math.Cos(radians)

	return Line{From: ln.From, To: NewPoint(ln.From.X+rx, ln.From.Y+ry)}
}

// IsCollision tests whether a moving object collides with another moving object within a given radius.
func (ln Line) IsCollision(tl Line, radius float64) bool {
	tv := tl.Direction()
	lv := ln.Direction()
	dx := tl.From.Sub(ln.From)
	vx2 := tv.X - lv.X
	vy2 := tv.Y - lv.Y

	a := vx2*vx2 + vy2*vy2
	if a <= 0 {
		return false
	}

	b := 2 * (dx.X*vx2 + dx.Y*vy2)
	c := dx.X*dx.X + dx.Y*dx.Y - radius*radius
	d := b*b - 4*a*c
	if d < 0 {
		return false
	}

	t := (-b - math.Sqrt(d)) / (2 * a)
	if t <= 0 || t > 1 {
		return false
	}

	return true
}

// Rect returns the rectangle of the Line.
func (ln Line) Rect() Rect {
	return NewRectangle(ln.From.X, ln.To.X, ln.From.Y, ln.To.Y)
}

func (ln Line) String() string {
	return fmt.Sprintf("(%s->%s)", ln.From, ln.To)
}

func NewLine(from, to Point) Line {
	return Line{from, to}
}

type Lines []Line
