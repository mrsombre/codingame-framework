package main

import (
	"math"
)

// distanceToBound returns the distance between x and the expected bound.
func distanceToBound(x, bound float64) float64 {
	return math.Min(x, bound-x)
}

// linesIntersection returns the intersection point of two Lines.
// Arguments s1 and s2 are true if the Line threaten as segment.
// https://en.wikipedia.org/wiki/Line%E2%80%93line_intersection
func linesIntersection(a, b Line, s1, s2 bool) (Point, bool) {
	av := a.Vector()
	bv := b.Vector()

	vcp := av.X*bv.Y - av.Y*bv.X
	if vcp == 0 {
		return Point{}, false
	}

	sv := b.From.Sub(a.From)
	cpa := sv.X*av.Y - sv.Y*av.X
	cpb := sv.X*bv.Y - sv.Y*bv.X
	t := cpb / vcp
	u := cpa / vcp

	if s1 && (t < 0 || t > 1) {
		return Point{}, false
	}
	if s2 && (u < 0 || u > 1) {
		return Point{}, false
	}

	return Point{
		X: a.From.X + t*av.X,
		Y: a.From.Y + t*av.Y,
	}, true
}

func isPointOnLine(ln Line, point Point, s bool) bool {
	lv := ln.Vector()
	pv := point.Sub(ln.From)

	pcp := lv.X*pv.Y - lv.Y*pv.X
	if pcp != 0 {
		return false
	}

	if s {
		dp := pv.X*lv.X + pv.Y*lv.Y
		l := lv.X*lv.X + lv.Y*lv.Y

		return dp >= 0 && dp <= l
	}

	return true
}
