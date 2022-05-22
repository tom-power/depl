package depl

import (
	"errors"
	"strings"
)

const buildLib = "sh/buildLib.sh"

const copyLibs = `
. sh/.env &&
cp -p ../$lib-lib.jar ./libs/ &&
cp -p ../$lib-lib-javadoc.jar ./libs/ &&
cp -p ../$lib-lib-sources.jar ./libs/ &&

cp -p ../$lib-testing.jar ./libs/ &&
cp -p ../$lib-testing-javadoc.jar ./libs/ &&
cp -p ../$lib-testing-sources.jar ./libs/
`
const copyLibsNested = `
. sh/.env &&
cp -p ../../$lib-lib.jar ./libs/ &&
cp -p ../../$lib-lib-javadoc.jar ./libs/ &&
cp -p ../../$lib-lib-sources.jar ./libs/ &&

cp -p ../../$lib-testing.jar ./libs/ &&
cp -p ../../$lib-testing-javadoc.jar ./libs/ &&
cp -p ../../$lib-testing-sources.jar ./libs/
`
const remote = `
. sh/.env &&
ssh -t $remoteUser@$remoteHost "cd $deployDir && sh/depl.sh $2"
`
const deployCopyLibs = `
cp -p ../flanalyst-lib.jar ./libs/ &&
cp -p ../depl ./sh/ &&
cp -p ../depl.sh ./sh/
`
const deploy = `
sh/depl.sh pull &&
sh/depl.sh deployCopyLibs &&
sh/depl.sh shadowJar &&
sh/depl.sh dockerUp
`
const catLog = "sh/catLog.sh"
const catEvent = "sh/catEvent.sh"

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
