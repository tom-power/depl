package depl

import (
	"errors"
	"sort"
	"strings"
)

var ScriptFor = func(command string) (string, error) {
	switch command {
	case "list":
		return commandList(), nil
	default:
		return scriptFor(command)

	}
}

func commandList() string {
	return strings.Join(sorted(keys(CommandsToScripts)), " ") + " list"
}

func sorted(input []string) []string {
	sorted := make([]string, len(input))
	copy(sorted, input)
	sort.Strings(sorted)
	return sorted
}

func keys[K comparable, V any](m map[K]V) []K {
	keys := make([]K, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}

func scriptFor(command string) (string, error) {
	script := CommandsToScripts[command]
	if script == "" {
		return "", errors.New("unknown command " + command)
	}
	return script, nil
}
