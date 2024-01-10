package main

import (
	"testing"
)

func BenchmarkLine_LinesIntersection(b *testing.B) {
	al := Line{Point{0, 0}, Point{300, 400}}
	bl := Line{Point{0, 400}, Point{300, 0}}
	for i := 0; i < b.N; i++ {
		linesIntersection(al, bl, true, true)
	}
}
