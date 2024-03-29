package main

import (
	"fmt"
	"math"
)

const (
	// sides
	top    = 0
	right  = 1
	bottom = 2
	left   = 3

	// corners
	topLeft     = 4
	topRight    = 5
	bottomRight = 6
	bottomLeft  = 7

	// 0-corners
	topLeft0     = 0
	topRight0    = 1
	bottomRight0 = 2
	bottomLeft0  = 3
)

// Point represents a point in a 2D plane.
type Point struct {
	X, Y float64
}

// IsSame tests if two Points are the same.
func (p Point) IsSame(t Point) bool {
	return p.X == t.X && p.Y == t.Y
}

// Index returns the 0-index of the Point in a field of specified width.
func (p Point) Index(width float64) int {
	if p.X < 0 || p.X >= width || p.Y < 0 {
		panic(fmt.Sprintf("point %s is out of bound %.f", p, width))
	}
	return int(p.Y*width + p.X)
}

// IsInXRange tests if the Point is in the range by X axis.
func (p Point) IsInXRange(from, to float64) bool {
	return p.X >= from && p.X <= to
}

// IsInXBound tests if the Point is in the field defined by width.
func (p Point) IsInXBound(width float64) bool {
	return p.X >= 0 && p.X < width
}

// IsInYRange tests if the Point is in the range by Y axis.
func (p Point) IsInYRange(from, to float64) bool {
	return p.Y >= from && p.Y <= to
}

// IsInYBound tests if the Point is in the field defined by height.
func (p Point) IsInYBound(height float64) bool {
	return p.Y >= 0 && p.Y < height
}

// IsInRange tests if the Point is in the range by X and Y axis.
func (p Point) IsInRange(from, to float64) bool {
	return p.X >= from && p.X <= to && p.Y >= from && p.Y <= to
}

// IsInBound tests if the Point is in the positive bound by X and Y axis.
func (p Point) IsInBound(width, height float64) bool {
	return p.X >= 0 && p.X < width && p.Y >= 0 && p.Y < height
}

// SymmetricX returns the symmetric Point by X axis.
func (p Point) SymmetricX(width float64) Point {
	return Point{width - p.X, p.Y}
}

// SymmetricY returns the symmetric Point by Y axis.
func (p Point) SymmetricY(height float64) Point {
	return Point{p.X, height - p.Y}
}

// Symmetric returns the symmetric Point by X and Y axis.
func (p Point) Symmetric(width, height float64) Point {
	return Point{width - p.X, height - p.Y}
}

// Add returns the sum of two Points.
func (p Point) Add(t Point) Point {
	return Point{p.X + t.X, p.Y + t.Y}
}

// Sub returns the difference of two Points.
func (p Point) Sub(t Point) Point {
	return Point{p.X - t.X, p.Y - t.Y}
}

// Distance returns the distance between two Points using the Pythagorean theorem.
func (p Point) Distance(t Point) float64 {
	x := p.X - t.X
	y := p.Y - t.Y
	return math.Sqrt(x*x + y*y)
}

// DistanceManhattan returns the Manhattan distance between two Points.
// https://en.wikipedia.org/wiki/Taxicab_geometry
func (p Point) DistanceManhattan(t Point) float64 {
	return math.Abs(p.X-t.X) + math.Abs(p.Y-t.Y)
}

// DistanceChebyshev returns the Chebyshev distance between two Points.
// https://en.wikipedia.org/wiki/Chebyshev_distance
func (p Point) DistanceChebyshev(t Point) float64 {
	return math.Max(math.Abs(p.X-t.X), math.Abs(p.Y-t.Y))
}

// SquareLength returns the square length of the Point vector.
func (p Point) SquareLength() float64 {
	return p.X*p.X + p.Y*p.Y
}

// DotProduct returns the dot product of two Points vectors.
func (p Point) DotProduct(t Point) float64 {
	return p.X*t.X + p.Y*t.Y
}

// CrossProduct returns the cross product of two Points vectors.
func (p Point) CrossProduct(t Point) float64 {
	return p.X*t.Y - p.Y*t.X
}

// Normalize returns the normalized Point vector of specified length.
func (p Point) Normalize(length float64) Point {
	return Point{p.X / length, p.Y / length}
}

// distanceToBound returns the distance between x and the expected bound.
func distanceToBound(x, bound float64) float64 {
	return math.Min(x, bound-x)
}

// DistanceToXBound returns the distance between the Point and the field bound of specified width.
func (p Point) DistanceToXBound(bound float64) float64 {
	return distanceToBound(p.X, bound)
}

// DistanceToYBound returns the distance between the Point and the field bound of specified height.
func (p Point) DistanceToYBound(bound float64) float64 {
	return distanceToBound(p.Y, bound)
}

// NeighborsCross returns the neighbors of the Point in the cross shape.
func (p Point) NeighborsCross() Points {
	return Points{
		top:    {p.X, p.Y + 1},
		right:  {p.X + 1, p.Y},
		bottom: {p.X, p.Y - 1},
		left:   {p.X - 1, p.Y},
	}
}

// NeighborsAround returns the neighbors of the Point in the around shape.
func (p Point) NeighborsAround() Points {
	return Points{
		top:         {p.X, p.Y + 1},
		topRight:    {p.X + 1, p.Y + 1},
		right:       {p.X + 1, p.Y},
		bottomRight: {p.X + 1, p.Y - 1},
		bottom:      {p.X, p.Y - 1},
		bottomLeft:  {p.X - 1, p.Y - 1},
		left:        {p.X - 1, p.Y},
		topLeft:     {p.X - 1, p.Y + 1},
	}
}

func (p Point) String() string {
	return fmt.Sprintf("[X:%.f,Y:%.f]", p.X, p.Y)
}

func NewPoint(x, y float64) Point {
	return Point{
		X: math.Round(x),
		Y: math.Round(y),
	}
}

type Points []Point
