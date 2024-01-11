package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRect_IsSame(t *testing.T) {
	var r Rect
	r = Rect{100, 200, 300, 400}
	assert.True(t, r.IsSame(Rect{100, 200, 300, 400}))
	r = Rect{100, 200, 300, 400}
	assert.False(t, r.IsSame(Rect{300, 400, 100, 200}))
}

func TestRect_Width(t *testing.T) {
	r := Rect{100, 200, 300, 400}
	assert.EqualValues(t, 100, r.Width())
}

func TestRect_Height(t *testing.T) {
	r := Rect{100, 200, 300, 400}
	assert.EqualValues(t, 100, r.Height())
}

func TestRect_Area(t *testing.T) {
	r := Rect{100, 200, 300, 400}
	assert.EqualValues(t, 10000, r.Area())
}

func TestRect_Center(t *testing.T) {
	r := Rect{100, 200, 300, 400}
	assert.Equal(t, Point{150, 350}, r.Center())
}

func TestRect_Symmetric(t *testing.T) {
	r := Rect{100, 200, 300, 400}
	assert.Equal(t, Rect{900, 800, 700, 600}, r.Symmetric(1000, 1000))
}

func TestRect_IsContainsPoint(t *testing.T) {
	var r Rect
	r = Rect{100, 200, 300, 400}
	assert.True(t, r.IsContainsPoint(Point{150, 350}))
	r = Rect{100, 200, 300, 400}
	assert.False(t, r.IsContainsPoint(Point{50, 250}))
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
			r:    Rect{100, 200, 300, 400},
			t:    Rect{120, 180, 320, 380},
			want: true,
		},
		{
			name: `false`,
			r:    Rect{100, 200, 300, 400},
			t:    Rect{150, 250, 350, 450},
			want: false,
		},
		{
			name: `same`,
			r:    Rect{100, 200, 300, 400},
			t:    Rect{100, 200, 300, 400},
			want: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.want, tc.r.IsContainsRectangle(tc.t))
		})
	}
}

func TestRect_IsIntersectsRect(t *testing.T) {
	tests := []struct {
		name string
		r    Rect
		t    Rect
		want bool
	}{
		{
			name: `true`,
			r:    Rect{100, 200, 300, 400},
			t:    Rect{150, 250, 350, 450},
			want: true,
		},
		{
			name: `false`,
			r:    Rect{100, 200, 300, 400},
			t:    Rect{300, 400, 500, 600},
			want: false,
		},
		{
			name: `same`,
			r:    Rect{100, 200, 300, 400},
			t:    Rect{100, 200, 300, 400},
			want: true,
		},
		{
			name: `inside`,
			r:    Rect{100, 200, 300, 400},
			t:    Rect{120, 180, 320, 380},
			want: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.want, tc.r.IsIntersectsRect(tc.t))
		})
	}
}

func TestRect_RectsIntersection(t *testing.T) {
	tests := []struct {
		name string
		r    Rect
		t    Rect
		want Rect
		ok   bool
	}{
		{
			name: `true`,
			r:    Rect{100, 200, 300, 400},
			t:    Rect{150, 250, 350, 450},
			want: Rect{150, 200, 350, 400},
			ok:   true,
		},
		{
			name: `false`,
			r:    Rect{100, 200, 300, 400},
			t:    Rect{300, 400, 500, 600},
			ok:   false,
		},
		{
			name: `same`,
			r:    Rect{100, 200, 300, 400},
			t:    Rect{100, 200, 300, 400},
			want: Rect{100, 200, 300, 400},
			ok:   true,
		},
		{
			name: `second>first`,
			r:    Rect{100, 200, 300, 400},
			t:    Rect{120, 180, 320, 380},
			want: Rect{120, 180, 320, 380},
			ok:   true,
		},
		{
			name: `first>second`,
			r:    Rect{100, 200, 300, 400},
			t:    Rect{80, 220, 280, 420},
			want: Rect{100, 200, 300, 400},
			ok:   true,
		},
		{
			name: `line`,
			r:    Rect{100, 200, 300, 400},
			t:    Rect{200, 300, 300, 400},
			ok:   false,
		},
		{
			name: `point`,
			r:    Rect{100, 200, 300, 400},
			t:    Rect{200, 300, 200, 300},
			ok:   false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got, ok := tc.r.RectsIntersection(tc.t)
			assert.Equal(t, tc.ok, ok)
			assert.Equal(t, tc.want, got)
		})
	}
}

func TestRect_Vertices(t *testing.T) {
	r := Rect{100, 200, 300, 400}
	assert.Equal(t, Points{
		topLeft0:     {100, 400},
		topRight0:    {200, 400},
		bottomRight0: {200, 300},
		bottomLeft0:  {100, 300},
	}, r.Vertices())
}

func TestRect_Edges(t *testing.T) {
	r := Rect{100, 200, 300, 400}
	assert.Equal(t, Lines{
		top:    {Point{100, 400}, Point{200, 400}},
		right:  {Point{200, 300}, Point{200, 400}},
		bottom: {Point{100, 300}, Point{200, 300}},
		left:   {Point{100, 300}, Point{100, 400}},
	}, r.Edges())
}

func TestRect_String(t *testing.T) {
	r := Rect{100, 200, 300, 400}
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
			want: Rect{100, 200, 300, 400},
		},
		{
			name: `reverse`,
			xf:   200,
			xt:   100,
			yf:   400,
			yt:   300,
			want: Rect{100, 200, 300, 400},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.want, NewRectangle(tc.xf, tc.xt, tc.yf, tc.yt))
		})
	}
}
