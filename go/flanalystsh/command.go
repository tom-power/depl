package flanalystsh

import (
	"errors"
	"strings"
)

const buildLib = "sh/sh/buildLib.sh"

const copyLibs = `
cp -p ../flanalyst-lib.jar ./libs/ &&
cp -p ../flanalyst-lib-javadoc.jar ./libs/ &&
cp -p ../flanalyst-lib-sources.jar ./libs/ &&

cp -p ../flanalyst-testing.jar ./libs/ &&
cp -p ../flanalyst-testing-javadoc.jar ./libs/ &&
cp -p ../flanalyst-testing-sources.jar ./libs/ &&

cp -p ../flanalystsh ./sh/sh/ &&
cp -p ../flanalyst.sh ./sh/sh/
`
const copyLibsNested = `
cp -p ../../flanalyst-lib.jar ./libs/ &&
cp -p ../../flanalyst-lib-javadoc.jar ./libs/ &&
cp -p ../../flanalyst-lib-sources.jar ./libs/ &&

cp -p ../../flanalyst-testing.jar ./libs/ &&
cp -p ../../flanalyst-testing-javadoc.jar ./libs/ &&
cp -p ../../flanalyst-testing-sources.jar ./libs/ &&

cp -p ../../flanalystsh ./sh/sh/ &&
cp -p ../../flanalyst.sh ./sh/sh/
`
const remote = `
. sh/.env &&
ssh -t $remoteUser@$remoteHost "cd $deployDir && sh/sh/flanalyst.sh $2"
`
const deployCopyLibs = `
cp -p ../flanalyst-lib.jar ./libs/
cp -p ../flanalystsh ./sh/sh/
cp -p ../flanalyst.sh ./sh/sh/
`
const deploy = `
sh/sh/flanalyst.sh pull &&
sh/sh/flanalyst.sh deployCopyLibs &&
sh/sh/flanalyst.sh shadowJar &&
sh/sh/flanalyst.sh dockerUp
`
const catLog = "sh/sh/catLog.sh"
const catEvent = "sh/sh/catEvent.sh"

var Commands = map[string]string{
	"buildLib":       buildLib,
	"copyLibs":       copyLibs,
	"copyLibsNested": copyLibsNested,
	"test":           "./gradlew test -Denv=local --parallel",
	"push":           "git push origin master -f",
	"remote":         remote,
	"pull":           "git fetch --all && git reset --hard origin/master",
	"deployCopyLibs": deployCopyLibs,
	"shadowJar":      "./gradlew :app:shadowJar --no-daemon",
	"dockerUp":       "docker-compose -f docker/docker-compose-app.yml up --build --detach",
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
