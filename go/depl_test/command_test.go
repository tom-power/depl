package depl_test

import (
	"github.com/tom-power/depl/go/depl"
	"strings"
	"testing"
)

func Test_command(t *testing.T) {
	t.Run("can run command", func(t *testing.T) {
		for command, value := range depl.Commands {
			sh, _ := depl.GetCommand(command)
			if !strings.Contains(sh, value) {
				t.Errorf("'%v' should contain '%v'", sh, value)
			}
		}
	})
	t.Run("can list commands", func(t *testing.T) {
		sh, _ := depl.GetCommand("list")
		for command, _ := range depl.Commands {
			if !strings.Contains(sh, command) {
				t.Errorf("'%v' should contain '%v'", sh, command)
			}
		}
		if !strings.Contains(sh, "list") {
			t.Errorf("'%v' should contain '%v'", sh, "list")
		}
	})
}
