package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInputGame(t *testing.T) {
	game := InputGame(readGameTests)
	want := Game{
		Units: []Unit{
			{1, 2, 3},
			{4, 5, 6},
			{7, 8, 9},
		},
	}
	assert.Equal(t, want, game)

	data := []string{
		"1",
		"1 2 a",
	}
	assert.Panics(t, func() { InputGame(data) })
}

func TestInputStep(t *testing.T) {
	turn := InputStep(readStepTests)
	want := Turn{
		Power: 1,
		L:     "R",
		R:     "L",
	}
	assert.Equal(t, want, turn)

	data := []string{
		"a a a",
	}
	assert.Panics(t, func() { InputStep(data) })
}
