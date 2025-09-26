package depl_test

import (
	"strings"
	"testing"

	"github.com/tom-power/depl/depl"
)

func Test_run(t *testing.T) {
	t.Run("can list commands", func(t *testing.T) {
		run, _ := depl.Run("list")
		for command := range depl.CommandToScript {
			if !strings.Contains(run, command) {
				t.Errorf("'%v' should contain '%v'", run, command)
			}
		}
		if !strings.Contains(run, "list") {
			t.Errorf("'%v' should contain '%v'", run, "list")
		}
	})

	t.Run("can run a command", func(t *testing.T) {
		for command, script := range depl.CommandToScript {
			run, _ := depl.Run(command)
			if !strings.Contains(run, script) {
				t.Errorf("'%v' should contain '%v'", run, script)
			}
			if strings.HasPrefix(run, "echo") {
				t.Errorf("'%v' should not contain '%v'", run, "echo")
			}
		}
	})
}
