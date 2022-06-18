package depl

import (
	"errors"
	"strings"
)

var GetCommand = func(cmd string) (string, error) {
	if cmd == "list" {
		return commandList()
	}
	return command(cmd)
}

func commandList() (string, error) {
	return strings.Join(keys(Commands), " ") + " list", nil
}

func keys[K comparable, V any](m map[K]V) []K {
	keys := make([]K, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}

func command(cmd string) (string, error) {
	command := Commands[cmd]
	if command == "" {
		return "", errors.New("unknown command " + command)
	}
	return command, nil
}
