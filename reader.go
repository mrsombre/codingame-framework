package main

// Example of reading game state from the standard input stream.

import (
	"bufio"
)

// ReadGame reads the game state from the standard input stream.
func ReadGame(s *bufio.Scanner) []string {
	data := make([]string, 0, 32)

	s.Scan()
	size := s.Text()
	data = append(data, size)
	for i := 0; i < StrToInt(size); i++ {
		s.Scan()
		data = append(data, s.Text())
	}

	return data
}

// ReadStep reads the game turn state from the standard input stream.
func ReadStep(s *bufio.Scanner) []string {
	data := make([]string, 0, 1)

	s.Scan()
	data = append(data, s.Text())

	return data
}
