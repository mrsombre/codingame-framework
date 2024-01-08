package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRect_IsSame(t *testing.T) {
	tests := []struct {
		name string
		a, b Rect
		want bool
	}{
		{
			name: `true`,
			a:    Rect{100, 200, 300, 400},
			b:    Rect{100, 200, 300, 400},
			want: true,
		},
		{
			name: `false`,
			a:    Rect{100, 200, 300, 400},
			b:    Rect{300, 400, 100, 200},
			want: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.want, tc.a.IsSame(tc.b))
		})
	}
}

func TestRect_Width(t *testing.T) {
	r := Rect{Xf: 100, Xt: 200, Yf: 300, Yt: 400}
	assert.EqualValues(t, 100, r.Width())
}

func TestRect_Height(t *testing.T) {
	r := Rect{Xf: 100, Xt: 200, Yf: 300, Yt: 400}
	assert.EqualValues(t, 100, r.Height())
}

func TestRect_Area(t *testing.T) {
	r := Rect{Xf: 100, Xt: 200, Yf: 300, Yt: 400}
	assert.EqualValues(t, 10000, r.Area())
}

func TestRect_Center(t *testing.T) {
	r := Rect{Xf: 100, Xt: 200, Yf: 300, Yt: 400}
	assert.Equal(t, Point{150, 350}, r.Center())
}

func TestRect_IsContainsPoint(t *testing.T) {
	tests := []struct {
		name string
		r    Rect
		p    Point
		want bool
	}{
		{
			name: `true`,
			r:    Rect{Xf: 100, Xt: 200, Yf: 300, Yt: 400},
			p:    Point{150, 350},
			want: true,
		},
		{
			name: `false`,
			r:    Rect{Xf: 100, Xt: 200, Yf: 300, Yt: 400},
			p:    Point{50, 250},
			want: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.want, tc.r.IsContainsPoint(tc.p))
		})
	}
}

func TestRect_IsContainsRectangle(t *testing.T) {
	tests := []struct {
		name string
		r    Rect
		t    Rect
		want bool
	}{
		{
			name: `true`,
			r:    Rect{Xf: 100, Xt: 300, Yf: 300, Yt: 500},
			t:    Rect{Xf: 150, Xt: 250, Yf: 350, Yt: 450},
			want: true,
		},
		{
			name: `false`,
			r:    Rect{Xf: 100, Xt: 200, Yf: 300, Yt: 400},
			t:    Rect{Xf: 150, Xt: 250, Yf: 350, Yt: 450},
			want: false,
		},
		{
			name: `same`,
			r:    Rect{Xf: 100, Xt: 200, Yf: 300, Yt: 400},
			t:    Rect{Xf: 100, Xt: 200, Yf: 300, Yt: 400},
			want: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.want, tc.r.IsContainsRectangle(tc.t))
		})
	}
}

func TestRectangle_String(t *testing.T) {
	r := Rect{Xf: 100, Xt: 200, Yf: 300, Yt: 400}
	assert.Equal(t, `[X:100>200,Y:300>400]`, r.String())
}

func TestNewRectangle(t *testing.T) {
	tests := []struct {
		name   string
		xf, xt float64
		yf, yt float64
		want   Rect
	}{
		{
			name: `standard`,
			xf:   100,
			xt:   200,
			yf:   300,
			yt:   400,
			want: Rect{Xf: 100, Xt: 200, Yf: 300, Yt: 400},
		},
		{
			name: `reverse`,
			xf:   200,
			xt:   100,
			yf:   400,
			yt:   300,
			want: Rect{Xf: 100, Xt: 200, Yf: 300, Yt: 400},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.want, NewRectangle(tc.xf, tc.xt, tc.yf, tc.yt))
		})
	}
}
