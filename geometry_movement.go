package main

import (
	"math"
)

const (
	// angles
	angleForward = 0
	angleLeft    = 90
	angleRight   = -90
	angleBack    = 180
)

// MovingDistance calculates the distance traveled by an object.
// formula: s = ut + 1/2at^2
func MovingDistance(speed, acceleration, time float64) float64 {
	return (speed * time) + (0.5 * acceleration * time * time)
}

// MovingVector calculates the vector of a moving object with static angle coordinate system.
// see https://www.codingame.com/training/easy/mars-lander-episode-1
func MovingVector(angle, power float64) Point {
	rad := angle * (math.Pi / 180)

	return Point{
		X: -power * math.Sin(rad),
		Y: power * math.Cos(rad),
	}
}
