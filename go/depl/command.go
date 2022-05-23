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

const catLog = "sh/sh/catLog.sh"

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

var GetCommand = func(command string) (string, error) {
	if command == "list" {
		return keysOfMap(Commands) + " list", nil
	}

	sh := Commands[command]
	if sh == "" {
		return "", errors.New("unknown command " + command)
	}
	return sh, nil
}

func keysOfMap(theMap map[string]string) string {
	keys := make([]string, 0, len(theMap))
	for key := range theMap {
		keys = append(keys, key)
	}
	return strings.Join(keys, " ")
}
