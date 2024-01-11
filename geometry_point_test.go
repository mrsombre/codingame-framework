package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPoint_IsSame(t *testing.T) {
	var p Point
	p = Point{200, 100}
	assert.True(t, p.IsSame(Point{200, 100}))
	p = Point{200, 100}
	assert.False(t, p.IsSame(Point{100, 200}))
}

func TestPoint_Index(t *testing.T) {
	tests := []struct {
		name  string
		p     Point
		width int
		want  int
		panic bool
	}{
		{
			name:  `first`,
			p:     Point{0, 0},
			width: 10,
			want:  0,
		},
		{
			name:  `last in row`,
			p:     Point{9, 0},
			width: 10,
			want:  9,
		},
		{
			name:  `next row`,
			p:     Point{0, 1},
			width: 10,
			want:  10,
		},
		{
			name:  `next column`,
			p:     Point{1, 0},
			width: 10,
			want:  1,
		},
		{
			name:  `last`,
			p:     Point{9, 9},
			width: 10,
			want:  99,
		},
		{
			name:  `out of bound`,
			p:     Point{10, 0},
			width: 10,
			panic: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if tc.panic {
				assert.Panics(t, func() { tc.p.Index(tc.width) })
				return
			}
			assert.Equal(t, tc.want, tc.p.Index(tc.width))
		})
	}
}

func TestPoint_IsInXBound(t *testing.T) {
	var p Point
	p = Point{200, 100}
	assert.True(t, p.IsInXBound(1000))
	p = Point{200, 100}
	assert.False(t, p.IsInXBound(100))
}

func TestPoint_IsInYBound(t *testing.T) {
	var p Point
	p = Point{100, 200}
	assert.True(t, p.IsInYBound(1000))
	p = Point{100, 200}
	assert.False(t, p.IsInYBound(100))
}

func TestPoint_IsInBound(t *testing.T) {
	var p Point
	p = Point{200, 100}
	assert.True(t, p.IsInBound(1000, 1000))
	p = Point{200, 100}
	assert.False(t, p.IsInBound(100, 100))
	p = Point{100, 200}
	assert.False(t, p.IsInBound(100, 100))
}

func TestPoint_SymmetricX(t *testing.T) {
	var p Point
	p = Point{10, 30}
	assert.Equal(t, Point{90, 30}, p.SymmetricX(100))
	p = Point{80, 30}
	assert.Equal(t, Point{20, 30}, p.SymmetricX(100))
}

func TestPoint_SymmetricY(t *testing.T) {
	var p Point
	p = Point{30, 10}
	assert.Equal(t, Point{30, 90}, p.SymmetricY(100))
	p = Point{30, 80}
	assert.Equal(t, Point{30, 20}, p.SymmetricY(100))
}

func TestPoint_Symmetric(t *testing.T) {
	var p Point
	p = Point{10, 10}
	assert.Equal(t, Point{90, 90}, p.Symmetric(100, 100))
	p = Point{80, 80}
	assert.Equal(t, Point{20, 20}, p.Symmetric(100, 100))
}

func TestPoint_Add(t *testing.T) {
	var p Point
	p = Point{200, 100}
	assert.Equal(t, Point{300, 200}, p.Add(Point{100, 100}))
	p = Point{200, 100}
	assert.Equal(t, Point{100, 0}, p.Add(Point{-100, -100}))
}

func TestPoint_Sub(t *testing.T) {
	var p Point
	p = Point{200, 100}
	assert.Equal(t, Point{100, 0}, p.Sub(Point{100, 100}))
	p = Point{200, 100}
	assert.Equal(t, Point{300, 200}, p.Sub(Point{-100, -100}))
}

func TestPoint_Distance(t *testing.T) {
	tests := []struct {
		name string
		a, b Point
		want float64
	}{
		{
			name: `straight`,
			a:    Point{0, 100},
			b:    Point{0, 300},
			want: 200,
		},
		{
			name: `diagonal`,
			a:    Point{0, 0},
			b:    Point{300, 400},
			want: 500,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.want, tc.a.Distance(tc.b))
		})
	}
}

func TestPoint_DistanceManhattan(t *testing.T) {
	tests := []struct {
		name string
		a, b Point
		want float64
	}{
		{
			name: `straight`,
			a:    Point{0, 100},
			b:    Point{0, 300},
			want: 200,
		},
		{
			name: `diagonal`,
			a:    Point{0, 0},
			b:    Point{300, 400},
			want: 700,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.want, tc.a.DistanceManhattan(tc.b))
		})
	}
}

func TestPoint_DistanceChebyshev(t *testing.T) {
	tests := []struct {
		name string
		a, b Point
		want float64
	}{
		{
			name: `straight`,
			a:    Point{0, 100},
			b:    Point{0, 300},
			want: 200,
		},
		{
			name: `diagonal`,
			a:    Point{0, 0},
			b:    Point{300, 400},
			want: 400,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.want, tc.a.DistanceChebyshev(tc.b))
		})
	}
}

func TestPoint_DistanceToXBound(t *testing.T) {
	var p Point
	p = Point{10, 30}
	assert.EqualValues(t, 10, p.DistanceToXBound(100))
	p = Point{80, 30}
	assert.EqualValues(t, 20, p.DistanceToXBound(100))
}

func TestPoint_DistanceToYBound(t *testing.T) {
	var p Point
	p = Point{30, 10}
	assert.EqualValues(t, 10, p.DistanceToYBound(100))
	p = Point{30, 80}
	assert.EqualValues(t, 20, p.DistanceToYBound(100))
}

func TestPoint_NeighborsCross(t *testing.T) {
	p := Point{5, 5}
	assert.Equal(t, Points{
		top:    {5, 6},
		right:  {6, 5},
		bottom: {5, 4},
		left:   {4, 5},
	}, p.NeighborsCross())
}

func TestPoint_NeighborsAround(t *testing.T) {
	p := Point{5, 5}
	assert.Equal(t, Points{
		top:         {5, 6},
		topRight:    {6, 6},
		right:       {6, 5},
		bottomRight: {6, 4},
		bottom:      {5, 4},
		bottomLeft:  {4, 4},
		left:        {4, 5},
		topLeft:     {4, 6},
	}, p.NeighborsAround())
}

func TestPoint_String(t *testing.T) {
	p := Point{200, 100}
	assert.Equal(t, "[X:200,Y:100]", p.String())
}

func TestNewPoint(t *testing.T) {
	tests := []struct {
		name string
		x, y float64
		want Point
	}{
		{
			name: `standard`,
			x:    200,
			y:    100,
			want: Point{200, 100},
		},
		{
			name: `float`,
			x:    200.5,
			y:    100.5,
			want: Point{201, 101},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.want, NewPoint(tc.x, tc.y))
		})
	}
}
