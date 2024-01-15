package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type mockWriter struct {
	data []byte
}

func (w *mockWriter) Write(p []byte) (n int, err error) {
	w.data = p
	return len(p), nil
}

func TestMockCommand_String(t *testing.T) {
	cmd := MockCommand{1, 2}
	assert.Equal(t, "1 2", cmd.String())
}

func TestExecuteCommand(t *testing.T) {
	commandOutput = &mockWriter{}
	ExecuteCommands(Commands{MockCommand{1, 2}})

	want := "1 2\n"
	got := commandOutput.(*mockWriter).data
	assert.Equal(t, want, string(got))
}
