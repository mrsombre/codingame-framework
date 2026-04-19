package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIntToStr(t *testing.T) {
	var x int
	x = 1
	assert.Equal(t, "1", IntToStr(x))
	x = -1
	assert.Equal(t, "-1", IntToStr(x))
	x = 0
	assert.Equal(t, "0", IntToStr(x))
}

func TestStrToInt(t *testing.T) {
	var s string
	s = "1"
	assert.Equal(t, 1, StrToInt(s))
	s = "-1"
	assert.Equal(t, -1, StrToInt(s))
	s = "0"
	assert.Equal(t, 0, StrToInt(s))
	s = "a"
	assert.Panics(t, func() { StrToInt(s) })
}

func TestBoolToInt(t *testing.T) {
	var b bool
	b = true
	assert.Equal(t, 1, BoolToInt(b))
	b = false
	assert.Equal(t, 0, BoolToInt(b))
}

func TestIntToBool(t *testing.T) {
	var x int
	x = 1
	assert.Equal(t, true, IntToBool(x))
	x = 0
	assert.Equal(t, false, IntToBool(x))
}
