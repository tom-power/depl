package flanalystsh_test

import (
	"../flanalystsh"
	"strings"
	"testing"
)

func Test_command(t *testing.T) {
	t.Run("can run command", func(t *testing.T) {
		for command, expected := range flanalystsh.Commands {
			sh, _ := flanalystsh.GetCommand(command)
			if !strings.Contains(sh, expected) {
				t.Errorf("'%v' should contain '%v'", sh, expected)
			}
		}
	})
}
