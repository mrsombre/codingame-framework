package main

// Example of how to read input data into game business objects.

import (
	"fmt"
)

func InputGame(data []string) Game {
	var err error
	var game Game

	var size int
	size = StrToInt(data[0])
	data = data[1:]
	var unit Unit
	game.Units = make([]Unit, 0, size)
	for i := 0; i < size; i++ {
		_, err = fmt.Sscan(data[i], &unit.x, &unit.y, &unit.z)
		if err != nil {
			panic(err)
		}
		game.Units = append(game.Units, unit)
	}

	// some additional logic

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
