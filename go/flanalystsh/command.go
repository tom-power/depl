package flanalystsh

import (
	"errors"
)

const testPushDeploy = `
sh/sh/flanalyst.sh test \
&& sh/sh/flanalyst.sh push \
&& sh/sh/flanalyst.sh remote deploy
`

const test = `
sh/sh/flanalyst.sh copyLibs && \
./gradlew test -Denv=local --parallel\
`

const remote = `
. sh/.env &&
ssh -t $remoteUser@$remoteHost "cd $deployDir && sh/sh/flanalyst.sh $2"
`
const deploy = `
sh/sh/flanalyst.sh pull && \
sh/sh/flanalyst.sh deployCopyLibs && \
sh/sh/flanalyst.sh build && \
sh/sh/flanalyst.sh dockerUp
`
const deployCopyLibs = `
cp ../flanalyst-lib.jar ./libs/ &&
cp ../flanalystsh ./sh/sh/ &&
cp ../flanalyst.sh ./sh/sh/
`

const catLog = "sh/sh/catLog.sh"
const catEvent = "sh/sh/catEvent.sh"
const buildLib = "sh/sh/buildLib.sh"

var Commands = map[string]string{
	"copyLibs":       "sh/sh/copyLibs.sh",
	"test":           test,
	"push":           "git push origin master -f",
	"testPushDeploy": testPushDeploy,
	"pull":           "git fetch --all && git reset --hard origin/master",
	"build":          "./gradlew :app:shadowJar --no-daemon",
	"dockerUp":       "docker-compose -f docker/docker-compose-app.yml up --build --detach",
	"dockerDown":     "docker-compose -f docker/docker-compose-app.yml down",
	"deploy":         deploy,
	"deployCopyLibs": deployCopyLibs,
	"remote":         remote,
	"catLog":         catLog,
	"catEvent":       catEvent,
	"buildLib":       buildLib}

var GetCommand = func(command string) (string, error) {
	sh := Commands[command]
	if sh == "" {
		return "", errors.New("unknown command " + command)
	}
	return sh, nil
}
