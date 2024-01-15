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

// Vector returns the vector of the Line.
// https://en.wikipedia.org/wiki/Euclidean_vector
func (ln Line) Vector() Point {
	return ln.To.Sub(ln.From)
}

// Slope returns the slope value of the Line.
// https://en.wikipedia.org/wiki/Slope
func (ln Line) Slope() float64 {
	return (ln.To.Y - ln.From.Y) / (ln.To.X - ln.From.X)
}

// CrossProduct returns the cross product of the Line.
// https://en.wikipedia.org/wiki/Cross_product
func (ln Line) CrossProduct() float64 {
	return ln.From.X*ln.To.Y - ln.From.Y*ln.To.X
}

// IsHorizontal tests if the Line is horizontal.
func (ln Line) IsHorizontal() bool {
	return ln.From.Y == ln.To.Y
}

// IsVertical tests if the Line is vertical.
func (ln Line) IsVertical() bool {
	return ln.From.X == ln.To.X
}

// IsMoving tests if the Line has magnitude.
func (ln Line) IsMoving() bool {
	return !ln.From.IsSame(ln.To)
}

// Segment returns the Line segment with the given length.
func (ln Line) Segment(length float64) Line {
	if length == 0 {
		return Line{ln.From, ln.From}
	}
	nv := ln.Vector().Normalize(ln.Length())

	return Line{
		ln.From,
		NewPoint(
			ln.From.X+nv.X*length,
			ln.From.Y+nv.Y*length,
		),
	}
}

// isPointOnLine tests if the Point is on the Line or Line segment.
func isPointOnLine(line Line, point Point, isSegment bool) bool {
	lv := line.Vector()
	pv := point.Sub(line.From)

	pcp := lv.CrossProduct(pv)
	if pcp != 0 {
		return false
	}

	if isSegment {
		dp := pv.DotProduct(lv)
		return dp >= 0 && dp <= lv.SquareLength()
	}

	return true
}

// IsPointOnLine tests if the Point is on the Line.
func (ln Line) IsPointOnLine(t Point) bool {
	return isPointOnLine(ln, t, false)
}

// IsPointOnSegment tests if the Point is on the Line segment.
func (ln Line) IsPointOnSegment(t Point) bool {
	return isPointOnLine(ln, t, true)
}

// closestPoint returns the closest Point on the Line or Line segment to the given Point.
func closestPoint(line Line, point Point, isSegment bool) Point {
	nv := line.Vector().Normalize(line.Length())
	dp := point.Sub(line.From).DotProduct(nv)

	if isSegment {
		if dp <= 0 {
			return line.From
		}
		if dp >= line.Length() {
			return line.To
		}
	}

	return NewPoint(
		line.From.X+nv.X*dp,
		line.From.Y+nv.Y*dp,
	)
}

// ClosestPointToLine returns the closest Point on the Line.
func (ln Line) ClosestPointToLine(t Point) Point {
	return closestPoint(ln, t, false)
}

// ClosestPointToSegment returns the closest Point on the Line segment.
func (ln Line) ClosestPointToSegment(t Point) Point {
	return closestPoint(ln, t, true)
}

// linesIntersection returns the intersection point of two Lines or Line segments.
func linesIntersection(lineA, lineB Line, isSegmentA, isSegmentB bool) (Point, bool) {
	av := lineA.Vector()
	bv := lineB.Vector()

	vcp := av.CrossProduct(bv)
	if vcp == 0 {
		return Point{}, false
	}

	sv := lineB.From.Sub(lineA.From)
	acp := sv.CrossProduct(av)
	bcp := sv.CrossProduct(bv)
	t := bcp / vcp
	u := acp / vcp

	if isSegmentA && (t < 0 || t > 1) {
		return Point{}, false
	}
	if isSegmentB && (u < 0 || u > 1) {
		return Point{}, false
	}

	return Point{
		X: lineA.From.X + t*av.X,
		Y: lineA.From.Y + t*av.Y,
	}, true
}

// LinesIntersection returns the crossing Point of two Lines.
func (ln Line) LinesIntersection(t Line) (Point, bool) {
	return linesIntersection(ln, t, false, false)
}

// SegmentsIntersection returns the crossing Point of two Line segments.
func (ln Line) SegmentsIntersection(t Line) (Point, bool) {
	return linesIntersection(ln, t, true, true)
}

// LineSegmentIntersection returns the crossing Point of the Line and the Line segment.
func (ln Line) LineSegmentIntersection(t Line) (Point, bool) {
	return linesIntersection(ln, t, false, true)
}

// Rotate returns the Line rotated by the given angle.
func (ln Line) Rotate(angle float64) Line {
	radians := angle * math.Pi / 180

	v := ln.Vector()
	rx := v.X*math.Cos(radians) + v.Y*math.Sin(radians)
	ry := v.X*-math.Sin(radians) + v.Y*math.Cos(radians)

	return Line{From: ln.From, To: NewPoint(ln.From.X+rx, ln.From.Y+ry)}
}

// IsCollision tests whether a moving object collides with another moving object within a given radius.
func (ln Line) IsCollision(t Line, radius float64) bool {
	vx := t.Vector().Sub(ln.Vector())
	dx := t.From.Sub(ln.From)

	a := vx.SquareLength()
	if a <= 0 {
		return false
	}

	b := 2 * dx.DotProduct(vx)
	c := dx.SquareLength() - radius*radius
	d := b*b - 4*a*c
	if d < 0 {
		return false
	}

	tc := (-b - math.Sqrt(d)) / (2 * a)
	if tc <= 0 || tc > 1 {
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
