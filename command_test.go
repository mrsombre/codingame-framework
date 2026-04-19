package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type commandMockWriter struct {
	data []byte
}

func (w *commandMockWriter) Write(p []byte) (n int, err error) {
	w.data = p
	return len(p), nil
}

func TestMockCommand_String(t *testing.T) {
	cmd := MockCommand{1, 2}
	assert.Equal(t, "1 2", cmd.String())
}

func TestExecuteCommand(t *testing.T) {
	commandOutput = &commandMockWriter{}
	ExecuteCommands(Commands{MockCommand{1, 2}})

	want := "1 2\n"
	got := commandOutput.(*commandMockWriter).data
	assert.Equal(t, want, string(got))
}
