package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMovingDistance(t *testing.T) {
	tests := []struct {
		name  string
		speed float64
		acc   float64
		time  float64
		want  float64
	}{
		{
			name:  `speed`,
			speed: 10,
			time:  5,
			want:  50,
		},
		{
			name: `acc`,
			acc:  10,
			time: 5,
			want: 125,
		},
		{
			name:  `speed and acc`,
			speed: 10,
			acc:   10,
			time:  5,
			want:  175,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.want, MovingDistance(tc.speed, tc.acc, tc.time))
		})
	}
}

func TestMovingVector(t *testing.T) {
	tests := []struct {
		name  string
		angle float64
		power float64
		want  Point
	}{
		{
			name:  `top`,
			angle: angleForward,
			power: 4,
			want:  Point{0, 4},
		},
		{
			name:  `right`,
			angle: angleRight,
			power: 4,
			want:  Point{4, 0},
		},
		{
			name:  `left`,
			angle: angleLeft,
			power: 4,
			want:  Point{-4, 0},
		},
		{
			name:  `back`,
			angle: angleBack,
			power: 4,
			want:  Point{0, -4},
		},
		{
			name:  `top left`,
			angle: 45,
			power: 4,
			want:  Point{-2.82, 2.828},
		},
		{
			name:  `top right`,
			angle: -45,
			power: 4,
			want:  Point{2.82, 2.828},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			mv := MovingVector(tc.angle, tc.power)
			assert.InDelta(t, tc.want.X, mv.X, 0.1)
			assert.InDelta(t, tc.want.Y, mv.Y, 0.1)
		})
	}
}
