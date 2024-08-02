package depl_test

import (
	"strings"
	"testing"

	"github.com/tom-power/depl/depl"
)

func Test_command(t *testing.T) {
	t.Run("can run command", func(t *testing.T) {
		for command, value := range depl.CommandsToScripts {
			sh, _ := depl.ScriptFor(command)
			if !strings.Contains(sh, value) {
				t.Errorf("'%v' should contain '%v'", sh, value)
			}
		}
	})
	t.Run("can list commands", func(t *testing.T) {
		sh, _ := depl.ScriptFor("list")
		for command := range depl.CommandsToScripts {
			if !strings.Contains(sh, command) {
				t.Errorf("'%v' should contain '%v'", sh, command)
			}
		}
		if !strings.Contains(sh, "list") {
			t.Errorf("'%v' should contain '%v'", sh, "list")
		}
	})
}
