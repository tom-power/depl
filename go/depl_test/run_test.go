package depl_test

import (
	"strings"
	"testing"

	"github.com/tom-power/depl/depl"
)

func Test_command(t *testing.T) {
	t.Run("can list commands", func(t *testing.T) {
		sh, _ := depl.Run("list", false)
		for command := range depl.CommandsToScripts {
			if !strings.Contains(sh, command) {
				t.Errorf("'%v' should contain '%v'", sh, command)
			}
		}
		if !strings.Contains(sh, "list") {
			t.Errorf("'%v' should contain '%v'", sh, "list")
		}
	})

	t.Run("can explain a command", func(t *testing.T) {
		for command, script := range depl.CommandsToScripts {
			run, _ := depl.Run(command, true)
			if !strings.Contains(run, script) {
				t.Errorf("'%v' should contain '%v'", run, script)
			}
			if !strings.HasPrefix(run, "echo") {
				t.Errorf("'%v' should contain '%v'", run, "echo")
			}
		}
	})

	t.Run("can run a command", func(t *testing.T) {
		for command, script := range depl.CommandsToScripts {
			run, _ := depl.Run(command, false)
			if !strings.Contains(run, script) {
				t.Errorf("'%v' should contain '%v'", run, script)
			}
			if strings.HasPrefix(run, "echo") {
				t.Errorf("'%v' should not contain '%v'", run, "echo")
			}
		}
	})
}
