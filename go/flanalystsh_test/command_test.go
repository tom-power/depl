package flanalystsh_test

import (
	"../flanalystsh"
	"strings"
	"testing"
)

func Test_command(t *testing.T) {
	t.Run("can run command", func(t *testing.T) {
		for command, value := range flanalystsh.Commands {
			sh, _ := flanalystsh.GetCommand(command)
			if !strings.Contains(sh, value) {
				t.Errorf("'%v' should contain '%v'", sh, value)
			}
		}
	})
	t.Run("can list commands", func(t *testing.T) {
		sh, _ := flanalystsh.GetCommand("list")
		for command, _ := range flanalystsh.Commands {
			if !strings.Contains(sh, command) {
				t.Errorf("'%v' should contain '%v'", sh, command)
			}
		}
		if !strings.Contains(sh, "list") {
			t.Errorf("'%v' should contain '%v'", sh, "list")
		}
	})
}
