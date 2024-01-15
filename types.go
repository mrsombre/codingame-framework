package main

// A collection of helper functions for type conversions.

import (
	"strconv"
)

// IntToStr converts an integer to a string.
func IntToStr(x int) string {
	return strconv.Itoa(x)
}

// StrToInt converts a string to an integer.
func StrToInt(s string) int {
	result, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return result
}

// BoolToInt converts a boolean value to an integer.
func BoolToInt(b bool) int {
	if b {
		return 1
	}
	return 0
}

// IntToBool converts an integer to a boolean value.
func IntToBool(x int) bool {
	if x != 0 {
		return true
	}
	return false
}
