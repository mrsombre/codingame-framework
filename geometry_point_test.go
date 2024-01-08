package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPoint_IsEqual(t *testing.T) {
	tests := []struct {
		name string
		a, b Point
		want bool
	}{
		{
			name: `true`,
			a:    Point{200, 100},
			b:    Point{200, 100},
			want: true,
		},
		{
			name: `false`,
			a:    Point{200, 100},
			b:    Point{100, 200},
			want: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.want, tc.a.IsEqual(tc.b))
		})
	}
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

func TestPoint_IsInBound(t *testing.T) {
	tests := []struct {
		name string
		p    Point
		w, h float64
		want bool
	}{
		{
			name: `true`,
			p:    Point{200, 100},
			w:    1000,
			h:    1000,
			want: true,
		},
		{
			name: `false width`,
			p:    Point{200, 50},
			w:    100,
			h:    100,
			want: false,
		},
		{
			name: `false height`,
			p:    Point{50, 200},
			w:    100,
			h:    100,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.want, tc.p.IsInBound(tc.w, tc.h))
		})
	}
}

func TestPoint_Add(t *testing.T) {
	tests := []struct {
		name string
		a, b Point
		want Point
	}{
		{
			name: `positive`,
			a:    Point{200, 100},
			b:    Point{300, 200},
			want: Point{500, 300},
		},
		{
			name: `negative`,
			a:    Point{200, 100},
			b:    Point{-300, -200},
			want: Point{-100, -100},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.want, tc.a.Add(tc.b))
		})
	}
}

func TestPoint_Sub(t *testing.T) {
	tests := []struct {
		name string
		a, b Point
		want Point
	}{
		{
			name: `positive`,
			a:    Point{200, 100},
			b:    Point{300, 200},
			want: Point{-100, -100},
		},
		{
			name: `negative`,
			a:    Point{200, 100},
			b:    Point{-300, -200},
			want: Point{500, 300},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.want, tc.a.Sub(tc.b))
		})
	}
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

func TestNewPoint(t *testing.T) {
	p := NewPoint(200.5, 100.5)
	assert.Equal(t, Point{201, 101}, p)
}
