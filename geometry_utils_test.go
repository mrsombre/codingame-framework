package main

import (
	"testing"
)

// Benchmarks

func BenchmarkLine_LinesIntersection(b *testing.B) {
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

func BenchmarkLine_IsPointOnLine(b *testing.B) {
	r := false
	line := Line{Point{0, 0}, Point{300, 300}}
	point := Point{150, 150}
	for i := 0; i < b.N; i++ {
		r = isPointOnLine(line, point, true)
	}
	GlobalB = r
}
