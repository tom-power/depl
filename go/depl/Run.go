package depl

import (
	"errors"
	"strings"
)

var Run = func(command string, explain bool) (string, error) {
	content, err := contentFor(command)
	if err != nil {
		return "", err
	}
	return echoFor(command, explain) + content, nil
}

func echoFor(command string, explain bool) string {
	if explain || command == "list" {
		return "echo "
	}
	return ""
}

func contentFor(command string) (string, error) {
	switch command {
	case "list":
		return commands(), nil
	default:
		return scriptFor(command)
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
