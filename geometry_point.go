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
)

type Point struct {
	X, Y float64
}

// IsEqual returns true if two Points are equal.
func (p Point) IsEqual(t Point) bool {
	return p.X == t.X && p.Y == t.Y
}

// Index returns the 0-index of the Point in a field of specified width.
func (p Point) Index(width int) int {
	if p.X < 0 || p.X >= float64(width) || p.Y < 0 {
		panic(fmt.Sprintf("point %v is out of bound %d", p, width))
	}
	return int(p.Y)*width + int(p.X)
}

// IsInXBound returns true if the Point is in the positive bound by X axis.
func (p Point) IsInXBound(width float64) bool {
	return p.X >= 0 && p.X <= width
}

// IsInYBound returns true if the Point is in the positive bound by Y axis.
func (p Point) IsInYBound(height float64) bool {
	return p.Y >= 0 && p.Y <= height
}

// IsInBound returns true if the Point is in the positive bound.
func (p Point) IsInBound(width, height float64) bool {
	return p.IsInXBound(width) && p.IsInYBound(height)
}

// Add returns the sum of two Point.
func (p Point) Add(t Point) Point {
	return Point{p.X + t.X, p.Y + t.Y}
}

// Sub returns the difference of two Point.
func (p Point) Sub(t Point) Point {
	return Point{p.X - t.X, p.Y - t.Y}
}

// Distance returns the distance between two Points using the Pythagorean theorem.
func (p Point) Distance(t Point) float64 {
	return math.Sqrt(math.Pow(p.X-t.X, 2) + math.Pow(p.Y-t.Y, 2))
}

// DistanceManhattan returns the Manhattan distance between two Points.
func (p Point) DistanceManhattan(t Point) float64 {
	return math.Abs(p.X-t.X) + math.Abs(p.Y-t.Y)
}

// DistanceToXBound returns the distance between the Point and the field bound of specified width.
func (p Point) DistanceToXBound(bound float64) float64 {
	return DistanceToBound(p.X, bound)
}

// DistanceToYBound returns the distance between the Point and the field bound of specified height.
func (p Point) DistanceToYBound(bound float64) float64 {
	return DistanceToBound(p.Y, bound)
}

func (p Point) NeighborsCross() Points {
	return Points{
		top:    {p.X, p.Y + 1},
		right:  {p.X + 1, p.Y},
		bottom: {p.X, p.Y - 1},
		left:   {p.X - 1, p.Y},
	}
}

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
	return fmt.Sprintf("(X:%.f,Y:%.f)", p.X, p.Y)
}

func NewPoint(x, y float64) Point {
	return Point{
		X: math.Round(x),
		Y: math.Round(y),
	}
}

type Points []Point
