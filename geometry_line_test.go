package main

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLine_IsSame(t *testing.T) {
	var ln Line
	ln = Line{Point{0, 0}, Point{300, 400}}
	assert.True(t, ln.IsSame(Line{Point{0, 0}, Point{300, 400}}))
	ln = Line{Point{0, 0}, Point{300, 400}}
	assert.False(t, ln.IsSame(Line{Point{300, 400}, Point{0, 0}}))
}

func TestLine_Length(t *testing.T) {
	var ln Line
	ln = NewLine(Point{0, 0}, Point{300, 400})
	assert.EqualValues(t, 500, ln.Length())
	ln = NewLine(Point{300, 400}, Point{-300, -400})
	assert.EqualValues(t, 1000, ln.Length())
}

func TestLine_Vector(t *testing.T) {
	var ln Line
	ln = Line{Point{200, 250}, Point{300, 300}}
	assert.Equal(t, Point{100, 50}, ln.Vector())
	ln = Line{Point{300, 300}, Point{200, 250}}
	assert.Equal(t, Point{-100, -50}, ln.Vector())
}

func TestLine_Slope(t *testing.T) {
	tests := []struct {
		name string
		line Line
		want float64
	}{
		{
			name: `horizontal left to right`,
			line: Line{Point{0, 0}, Point{300, 0}},
			want: 0,
		},
		{
			name: `horizontal right to left`,
			line: Line{Point{300, 0}, Point{0, 0}},
			want: 0,
		},
		{
			name: `vertical bottom to top`,
			line: Line{Point{0, 0}, Point{0, 300}},
			want: math.Inf(1),
		},
		{
			name: `vertical top to bottom`,
			line: Line{Point{0, 300}, Point{0, 0}},
			want: math.Inf(-1),
		},
		{
			name: `diagonal ascending`,
			line: Line{Point{0, 0}, Point{300, 300}},
			want: 1,
		},
		{
			name: `diagonal descending`,
			line: Line{Point{0, 300}, Point{300, 0}},
			want: -1,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.want, tc.line.Slope())
		})
	}
}

func TestLine_CrossProduct(t *testing.T) {
	var ln Line
	ln = Line{Point{200, 200}, Point{300, 300}}
	assert.EqualValues(t, 0, ln.CrossProduct())
	ln = Line{Point{100, 200}, Point{300, 400}}
	assert.EqualValues(t, -20000, ln.CrossProduct())
}

func TestLine_IsHorizontal(t *testing.T) {
	var ln Line
	ln = Line{Point{0, 0}, Point{300, 0}}
	assert.True(t, ln.IsHorizontal())
	ln = Line{Point{0, 0}, Point{0, 300}}
	assert.False(t, ln.IsHorizontal())
}

func TestLine_IsVertical(t *testing.T) {
	var ln Line
	ln = Line{Point{0, 0}, Point{0, 300}}
	assert.True(t, ln.IsVertical())
	ln = Line{Point{0, 0}, Point{300, 0}}
	assert.False(t, ln.IsVertical())
}

func TestLine_IsMoving(t *testing.T) {
	var ln Line
	ln = Line{Point{0, 0}, Point{300, 400}}
	assert.True(t, ln.IsMoving())
	ln = Line{Point{300, 400}, Point{300, 400}}
	assert.False(t, ln.IsMoving())
}

func TestLine_Segment(t *testing.T) {
	tests := []struct {
		name   string
		line   Line
		length float64
		want   Line
	}{
		{
			name:   `horizontal longer`,
			line:   Line{Point{0, 0}, Point{300, 0}},
			length: 500,
			want:   Line{Point{0, 0}, Point{500, 0}},
		},
		{
			name:   `vertical longer`,
			line:   Line{Point{0, 0}, Point{0, 300}},
			length: 500,
			want:   Line{Point{0, 0}, Point{0, 500}},
		},
		{
			name:   `diagonal exact`,
			line:   Line{Point{0, 0}, Point{300, 400}},
			length: 500,
			want:   Line{Point{0, 0}, Point{300, 400}},
		},
		{
			name:   `diagonal shorter`,
			line:   Line{Point{0, 0}, Point{600, 800}},
			length: 500,
			want:   Line{Point{0, 0}, Point{300, 400}},
		},
		{
			name:   `diagonal longer`,
			line:   Line{Point{0, 0}, Point{150, 200}},
			length: 500,
			want:   Line{Point{0, 0}, Point{300, 400}},
		},
		{
			name:   `zero`,
			line:   Line{Point{0, 0}, Point{300, 400}},
			length: 0,
			want:   Line{Point{0, 0}, Point{0, 0}},
		},
		{
			name:   `negative`,
			line:   Line{Point{300, 400}, Point{600, 800}},
			length: -500,
			want:   Line{Point{300, 400}, Point{0, 0}},
		},
		{
			name:   `backwards`,
			line:   Line{Point{300, 400}, Point{0, 0}},
			length: 250,
			want:   Line{Point{300, 400}, Point{150, 200}},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.want, tc.line.Segment(tc.length))
		})
	}
}

func TestLine_IsPointOnLine(t *testing.T) {
	tests := []struct {
		name  string
		line  Line
		point Point
		want  bool
	}{
		{
			name:  `horizontal`,
			line:  Line{Point{0, 0}, Point{300, 0}},
			point: Point{150, 0},
			want:  true,
		},
		{
			name:  `vertical`,
			line:  Line{Point{0, 0}, Point{0, 300}},
			point: Point{0, 150},
			want:  true,
		},
		{
			name:  `diagonal ascending`,
			line:  Line{Point{0, 0}, Point{300, 300}},
			point: Point{150, 150},
			want:  true,
		},
		{
			name:  `diagonal descending`,
			line:  Line{Point{0, 300}, Point{300, 0}},
			point: Point{150, 150},
			want:  true,
		},
		{
			name:  `diagonal reverse`,
			line:  Line{Point{300, 400}, Point{0, 0}},
			point: Point{600, 800},
			want:  true,
		},
		{
			name:  `false`,
			line:  Line{Point{0, 0}, Point{300, 300}},
			point: Point{150, 100},
			want:  false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.want, tc.line.IsPointOnLine(tc.point))
		})
	}
}

func TestLine_IsPointOnSegment(t *testing.T) {
	tests := []struct {
		name  string
		line  Line
		point Point
		want  bool
	}{
		{
			name:  `horizontal out of bounds`,
			line:  Line{Point{0, 0}, Point{300, 0}},
			point: Point{450, 0},
			want:  false,
		},
		{
			name:  `vertical out of bounds`,
			line:  Line{Point{0, 0}, Point{0, 300}},
			point: Point{0, 450},
			want:  false,
		},
		{
			name:  `diagonal ascending out of bounds`,
			line:  Line{Point{0, 0}, Point{300, 300}},
			point: Point{450, 450},
			want:  false,
		},
		{
			name:  `diagonal descending out of bounds`,
			line:  Line{Point{0, 300}, Point{300, 0}},
			point: Point{450, 450},
			want:  false,
		},
		{
			name:  `false in bounds`,
			line:  Line{Point{0, 0}, Point{300, 300}},
			point: Point{150, 250},
			want:  false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.want, tc.line.IsPointOnSegment(tc.point))
		})
	}
}

func TestLine_ClosestPointToLine(t *testing.T) {
	tests := []struct {
		name  string
		line  Line
		point Point
		want  Point
	}{
		{
			name:  `horizontal`,
			line:  Line{Point{0, 0}, Point{300, 0}},
			point: Point{400, 100},
			want:  Point{400, 0},
		},
		{
			name:  `vertical`,
			line:  Line{Point{0, 0}, Point{0, 300}},
			point: Point{100, 400},
			want:  Point{0, 400},
		},
		{
			name:  `diagonal from`,
			line:  Line{Point{150, 150}, Point{250, 300}},
			point: Point{50, 50},
			want:  Point{73, 35},
		},
		{
			name:  `diagonal to`,
			line:  Line{Point{0, 0}, Point{100, 50}},
			point: Point{400, 0},
			want:  Point{320, 160},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.want, tc.line.ClosestPointToLine(tc.point))
		})
	}
}

func TestLine_ClosestPointToSegment(t *testing.T) {
	tests := []struct {
		name  string
		line  Line
		point Point
		want  Point
	}{
		{
			name:  `horizontal`,
			line:  Line{Point{0, 0}, Point{300, 0}},
			point: Point{400, 100},
			want:  Point{300, 0},
		},
		{
			name:  `vertical`,
			line:  Line{Point{0, 0}, Point{0, 300}},
			point: Point{100, 400},
			want:  Point{0, 300},
		},
		{
			name:  `diagonal from`,
			line:  Line{Point{150, 150}, Point{250, 300}},
			point: Point{50, 50},
			want:  Point{150, 150},
		},
		{
			name:  `diagonal to`,
			line:  Line{Point{0, 0}, Point{100, 50}},
			point: Point{400, 0},
			want:  Point{100, 50},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.want, tc.line.ClosestPointToSegment(tc.point))
		})
	}
}

func TestLine_LinesIntersection(t *testing.T) {
	tests := []struct {
		name string
		a, b Line
		want Point
		ok   bool
	}{
		{
			name: `parallel horizontal`,
			a:    Line{Point{0, 0}, Point{300, 0}},
			b:    Line{Point{0, 100}, Point{300, 100}},
			want: Point{},
			ok:   false,
		},
		{
			name: `parallel vertical`,
			a:    Line{Point{0, 0}, Point{0, 300}},
			b:    Line{Point{100, 0}, Point{100, 300}},
			want: Point{},
			ok:   false,
		},
		{
			name: `horizontal to vertical`,
			a:    Line{Point{0, 150}, Point{300, 150}},
			b:    Line{Point{150, 0}, Point{150, 300}},
			want: Point{150, 150},
			ok:   true,
		},
		{
			name: `vertical to horizontal`,
			a:    Line{Point{150, 0}, Point{150, 300}},
			b:    Line{Point{0, 150}, Point{300, 150}},
			want: Point{150, 150},
			ok:   true,
		},
		{
			name: `parallel diagonal`,
			a:    Line{Point{0, 0}, Point{300, 400}},
			b:    Line{Point{300, 0}, Point{600, 400}},
			want: Point{},
			ok:   false,
		},
		{
			name: `diagonal`,
			a:    Line{Point{0, 0}, Point{300, 400}},
			b:    Line{Point{0, 400}, Point{300, 0}},
			want: Point{150, 200},
			ok:   true,
		},
		{
			name: `diagonal reverse`,
			a:    Line{Point{300, 300}, Point{0, 0}},
			b:    Line{Point{300, 0}, Point{0, 300}},
			want: Point{150, 150},
			ok:   true,
		},
		{
			name: `diagonal/horizontal`,
			b:    Line{Point{0, 0}, Point{300, 300}},
			a:    Line{Point{0, 150}, Point{300, 150}},
			want: Point{150, 150},
			ok:   true,
		},
		{
			name: `diagonal/vertical`,
			a:    Line{Point{0, 0}, Point{300, 300}},
			b:    Line{Point{150, 0}, Point{150, 300}},
			want: Point{150, 150},
			ok:   true,
		},
		{
			name: `backwards`,
			a:    Line{Point{300, 300}, Point{600, 600}},
			b:    Line{Point{0, 200}, Point{200, 0}},
			want: Point{100, 100},
			ok:   true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got, ok := tc.a.LinesIntersection(tc.b)
			assert.Equal(t, tc.want, got)
			assert.Equal(t, tc.ok, ok)
		})
	}
}

func TestLine_SegmentsIntersection(t *testing.T) {
	tests := []struct {
		name string
		a, b Line
		want Point
		ok   bool
	}{
		{
			name: `inside`,
			a:    Line{Point{0, 0}, Point{300, 400}},
			b:    Line{Point{200, 200}, Point{200, 100}},
		},
		{
			name: `outside`,
			a:    Line{Point{300, 300}, Point{600, 600}},
			b:    Line{Point{0, 200}, Point{200, 0}},
			ok:   false,
		},
		{
			name: `true`,
			a:    Line{Point{0, 0}, Point{300, 400}},
			b:    Line{Point{0, 200}, Point{300, 200}},
			want: Point{150, 200},
			ok:   true,
		},
		{
			name: `false`,
			a:    Line{Point{0, 0}, Point{300, 400}},
			b:    Line{Point{300, 0}, Point{600, 400}},
			ok:   false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got, ok := tc.a.SegmentsIntersection(tc.b)
			assert.Equal(t, tc.want, got)
			assert.Equal(t, tc.ok, ok)
		})
	}
}

func TestLine_LineSegmentIntersection(t *testing.T) {
	tests := []struct {
		name string
		a, b Line
		want Point
		ok   bool
	}{
		{
			name: `inside`,
			a:    Line{Point{0, 200}, Point{100, 200}},
			b:    Line{Point{0, 0}, Point{300, 400}},
			want: Point{150, 200},
			ok:   true,
		},
		{
			name: `outside`,
			a:    Line{Point{0, 0}, Point{300, 400}},
			b:    Line{Point{0, 800}, Point{1000, 800}},
			want: Point{600, 800},
			ok:   true,
		},
		{
			name: `outside segment`,
			a:    Line{Point{0, 0}, Point{300, 400}},
			b:    Line{Point{200, 200}, Point{400, 200}},
			ok:   false,
		},
		{
			name: `true`,
			a:    Line{Point{0, 0}, Point{300, 400}},
			b:    Line{Point{0, 200}, Point{300, 200}},
			want: Point{150, 200},
			ok:   true,
		},
		{
			name: `false`,
			a:    Line{Point{0, 0}, Point{300, 400}},
			b:    Line{Point{300, 0}, Point{600, 400}},
			ok:   false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got, ok := tc.a.LineSegmentIntersection(tc.b)
			assert.Equal(t, tc.want, got)
			assert.Equal(t, tc.ok, ok)
		})
	}
}

func TestLine_Rotate(t *testing.T) {
	tests := []struct {
		name  string
		line  Line
		angle float64
		want  Line
	}{
		{
			name:  `cw 90`,
			line:  Line{Point{500, 500}, Point{1000, 500}},
			angle: 90,
			want:  Line{Point{500, 500}, Point{500, 0}},
		},
		{
			name:  `cw 45`,
			line:  Line{Point{500, 500}, Point{1000, 500}},
			angle: 45,
			want:  Line{Point{500, 500}, Point{854, 146}},
		},
		{
			name:  `ccw 90`,
			line:  Line{Point{500, 500}, Point{1000, 500}},
			angle: -90,
			want:  Line{Point{500, 500}, Point{500, 1000}},
		},
		{
			name:  `ccw 45`,
			line:  Line{Point{500, 500}, Point{1000, 500}},
			angle: -45,
			want:  Line{Point{500, 500}, Point{854, 854}},
		},
		{
			name:  `180`,
			line:  Line{Point{500, 500}, Point{1000, 500}},
			angle: 180,
			want:  Line{Point{500, 500}, Point{0, 500}},
		},
		{
			name:  `360`,
			line:  Line{Point{500, 500}, Point{1000, 500}},
			angle: 360,
			want:  Line{Point{500, 500}, Point{1000, 500}},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.want, tc.line.Rotate(tc.angle))
		})
	}
}

func TestLine_IsCollision(t *testing.T) {
	tests := []struct {
		name   string
		a, b   Line
		radius float64
		want   bool
	}{
		{
			name:   `simple`,
			a:      Line{Point{100, 100}, Point{200, 200}},
			b:      Line{Point{300, 300}, Point{200, 200}},
			radius: 50,
			want:   true,
		},
		{
			name:   `radius`,
			a:      Line{Point{100, 100}, Point{200, 200}},
			b:      Line{Point{400, 400}, Point{300, 300}},
			radius: 150,
			want:   true,
		},
		{
			name:   `same speed`,
			a:      Line{Point{0, 0}, Point{300, 400}},
			b:      Line{Point{300, 400}, Point{600, 800}},
			radius: 100,
			want:   false,
		},
		{
			name:   `different direction`,
			a:      Line{Point{300, 300}, Point{0, 0}},
			b:      Line{Point{600, 600}, Point{900, 900}},
			radius: 100,
			want:   false,
		},
		{
			name:   `towards`,
			a:      Line{Point{0, 0}, Point{300, 400}},
			b:      Line{Point{400, 400}, Point{100, 0}},
			radius: 100,
			want:   true,
		},
		{
			name:   `towards over time`,
			a:      Line{Point{0, 0}, Point{600, 800}},
			b:      Line{Point{300, 400}, Point{150, 200}},
			radius: 50,
			want:   true,
		},
		{
			name:   `collides standing`,
			a:      Line{Point{0, 0}, Point{300, 400}},
			b:      Line{Point{150, 150}, Point{150, 150}},
			radius: 50,
			want:   true,
		},
		{
			name:   `both standing`,
			a:      Line{Point{0, 0}, Point{0, 0}},
			b:      Line{Point{300, 300}, Point{300, 300}},
			radius: 50,
			want:   false,
		},
		{
			name:   `discriminant negative`,
			a:      Line{Point{0, 0}, Point{0, 400}},
			b:      Line{Point{100, 0}, Point{400, 0}},
			radius: 50,
			want:   false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.want, tc.a.IsCollision(tc.b, tc.radius))
		})
	}
}

func TestLine_Rect(t *testing.T) {
	line := NewLine(Point{300, 400}, Point{0, 0})
	assert.Equal(t, Rect{0, 300, 0, 400}, line.Rect())
}

func TestLine_String(t *testing.T) {
	line := NewLine(Point{0, 0}, Point{300, 400})
	assert.Equal(t, "([X:0,Y:0]->[X:300,Y:400])", line.String())
}

func TestNewLine(t *testing.T) {
	ln := NewLine(Point{0, 0}, Point{300, 400})
	assert.Equal(t, Line{Point{0, 0}, Point{300, 400}}, ln)
}

// Benchmarks

var (
	GlobalLine Line
)

func BenchmarkIsPointOnLine(b *testing.B) {
	r := false
	line := Line{Point{0, 0}, Point{300, 300}}
	point := Point{150, 150}
	for i := 0; i < b.N; i++ {
		r = isPointOnLine(line, point, true)
	}
	GlobalB = r
}

func BenchmarkClosestPoint(b *testing.B) {
	t := Point{100, 100}
	p := Point{}
	ln := Line{Point{0, 0}, Point{300, 400}}
	for i := 0; i < b.N; i++ {
		p = closestPoint(ln, t, true)
	}
	GlobalPoint = p
}

func BenchmarkLinesIntersection(b *testing.B) {
	p := Point{}
	r := false
	al := Line{Point{0, 0}, Point{300, 400}}
	bl := Line{Point{0, 400}, Point{300, 0}}
	for i := 0; i < b.N; i++ {
		p, r = linesIntersection(al, bl, true, true)
	}
	GlobalPoint = p
	GlobalB = r
}

func BenchmarkLine_IsCollision(b *testing.B) {
	r := false
	al := Line{Point{0, 0}, Point{600, 800}}
	bl := Line{Point{300, 400}, Point{150, 200}}
	for i := 0; i < b.N; i++ {
		r = al.IsCollision(bl, 50)
	}
	GlobalB = r
}
