package main

import (
	"math"
)

// DistanceToBound returns the distance between x and the field 0 to bound.
func DistanceToBound(x, bound float64) float64 {
	return math.Min(x, bound-x)
}
