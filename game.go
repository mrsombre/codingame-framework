package main

// Example of game business logic.

type Unit struct {
	x, y, z float64
}

type Turn struct {
	Power float64
	L, R  string
}

type Game struct {
	Units []Unit
}
