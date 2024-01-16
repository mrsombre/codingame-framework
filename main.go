package main

// Example main: Read game state, construct business objects, apply logic, and execute commands.

import (
	"bufio"
	"math/rand"
	"os"
	"runtime"
	"time"
)

var rnd *rand.Rand

func init() {
	runtime.GOMAXPROCS(1)
	rnd = rand.New(rand.NewSource(time.Now().UnixNano()))
	debug = true
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(make([]byte, 1000000), 1000000)

	dataGame := ReadGame(scanner)
	asText(DataExport(dataGame))
	game := InputGame(dataGame)

	dataStep := ReadStep(scanner)
	asText(DataExport(dataStep))
	step := InputStep(dataStep)

	// some game logic for the first step
	u(game, step)

	for {
		dataStep = ReadStep(scanner)
		asText(DataExport(dataStep))
		step = InputStep(dataStep)

		// some game logic for the next step
		u(game, step)

		commands := Commands{
			MockCommand{1, 2},
		}
		ExecuteCommands(commands)
	}
}
