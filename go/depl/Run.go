package depl

import (
	"errors"
	"fmt"
	"strings"
)

var Run = func(command string, explain bool) (string, error) {
	switch command {
	case "list":
		return fmt.Sprintf("echo \"%s\"", commands()), nil
	default:
		script, err := scriptFor(command)
		if err == nil {
			if explain {
				return fmt.Sprintf("echo \"%s\"", script), nil
			}
			return script, nil
		}
		return "", err
	}
}

func commands() string {
	return strings.Join(sorted(keys(CommandsToScripts)), " ") + " list"
}

func scriptFor(command string) (string, error) {
	script := CommandsToScripts[command]
	if script == "" {
		return "", errors.New("unknown command " + command)
	}
	return script, nil
}
