package depl

import (
	"errors"
	"strings"
)

var Run = func(command string) (string, error) {
	content, err := contentFor(command)
	if err != nil {
		return "", err
	}
	return content, nil
}

func contentFor(command string) (string, error) {
	switch command {
	case "list":
		return commands(), nil
	default:
		return contentFromScriptFor(command)
	}
}

func commands() string {
	return strings.Join(sorted(keys(CommandToScript)), " ") + " list"
}

func contentFromScriptFor(command string) (string, error) {
	content := CommandToScript[command]
	if content == "" {
		return "", errors.New("unknown command " + command)
	}
	return content, nil
}
