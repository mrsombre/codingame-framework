package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDistanceToBound(t *testing.T) {
	tests := []struct {
		name  string
		x     float64
		bound float64
		want  float64
	}{
		{
			name:  `left`,
			x:     400,
			bound: 1000,
			want:  400,
		},
		{
			name:  `right`,
			x:     600,
			bound: 1000,
			want:  400,
		},
		{
			name:  `middle`,
			x:     500,
			bound: 1000,
			want:  500,
		},
		{
			name:  `outside`,
			x:     1400,
			bound: 1000,
			want:  -400,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.want, DistanceToBound(tc.x, tc.bound))
		})
	}
}
