package depl

import (
	_ "embed"
	"errors"
	"strings"
)

//go:embed sh/buildLib.sh
var buildLib string

//go:embed sh/copyLibs.sh
var copyLibs string

//go:embed sh/copyLibsNested.sh
var copyLibsNested string

//go:embed sh/test.sh
var test string

//go:embed sh/push.sh
var push string

//go:embed sh/remote.sh
var remote string

//go:embed sh/pull.sh
var pull string

//go:embed sh/deployCopyLibs.sh
var deployCopyLibs string

//go:embed sh/shadowJar.sh
var shadowJar string

//go:embed sh/dockerUp.sh
var dockerUp string

//go:embed sh/deploy.sh
var deploy string

//go:embed sh/catEvent.sh
var catEvent string

//go:embed sh/catLog.sh
var catLog string

var Commands = map[string]string{
	"buildLib":       buildLib,
	"copyLibs":       copyLibs,
	"copyLibsNested": copyLibsNested,
	"test":           test,
	"push":           push,
	"remote":         remote,
	"pull":           pull,
	"deployCopyLibs": deployCopyLibs,
	"shadowJar":      shadowJar,
	"dockerUp":       dockerUp,
	"deploy":         deploy,
	"catLog":         catLog,
	"catEvent":       catEvent}

var GetCommand = func(cmd string) (string, error) {
	if cmd == "list" {
		return commandList()
	}
  return command(cmd)
}

func commandList() (string, error) {
  return strings.Join(keys(Commands), " ") + " list", nil
}

func command(cmd string) (string, error) {
  command := Commands[cmd]
  if command == "" {
    return "", errors.New("unknown command " + command)
  }
  return command, nil
}

func keys[K comparable, V any](m map[K]V) []K {
    keys := make([]K, 0, len(m))
    for k := range m {
        keys = append(keys, k)
    }
    return keys
}
