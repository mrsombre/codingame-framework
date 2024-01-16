package main

import (
	"bufio"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

var readGameTests = []string{
	"3",
	"1 2 3",
	"4 5 6",
	"7 8 9",
}

func TestReadGame(t *testing.T) {
	s := strings.Join(readGameTests, "\n")
	r := strings.NewReader(s)
	b := bufio.NewScanner(r)

	data := ReadGame(b)
	assert.Equal(t, readGameTests, data)
}

var readStepTests = []string{
	"1 R L",
}

func TestReadStep(t *testing.T) {
	s := strings.Join(readStepTests, "\n")
	r := strings.NewReader(s)
	b := bufio.NewScanner(r)

	data := ReadStep(b)
	assert.Equal(t, readStepTests, data)
}
