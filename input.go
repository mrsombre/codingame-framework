package main

import (
	"fmt"
)

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

func InputGame(data []string) Game {
	var err error
	var game Game

	var size int
	size = StrToInt(data[0])
	data = data[1:]
	var unit Unit
	units := make([]Unit, 0, size)
	for i := 0; i < size; i++ {
		_, err = fmt.Sscan(data[i], &unit.x, &unit.y, &unit.z)
		if err != nil {
			panic(err)
		}
		units = append(units, unit)
	}

	// some additional logic
	game.Units = units

	return game
}

func InputStep(data []string) Turn {
	var err error

	var turn Turn
	_, err = fmt.Sscan(
		data[0],
		&turn.Power,
		&turn.L,
		&turn.R,
	)
	if err != nil {
		panic(err)
	}

	// some additional logic

	return turn
}
