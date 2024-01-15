package main

import (
	"fmt"
	"io"
	"os"
)

// Command is an interface for game commands.

var commandOutput io.Writer = os.Stdout

type Command interface {
	String() string
}

type Commands []Command

type MockCommand struct {
	Param1 float64
	Param2 float64
}

func (c MockCommand) String() string {
	return fmt.Sprintf("%.f %.f", c.Param1, c.Param2)
}

func ExecuteCommands(commands Commands) {
	for _, command := range commands {
		fmt.Fprintln(commandOutput, command)
	}
}
