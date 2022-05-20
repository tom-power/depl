package flanalystsh

import (
	"errors"
)

const testPushDeploy = `
sh/sh/flanalyst.sh test \
&& sh/sh/flanalyst.sh push \
&& sh/sh/flanalyst.sh remote deploy
`

const remote = `
. sh/.env &&
ssh -t $remoteUser@$remoteHost "cd $deployDir && sh/sh/flanalyst.sh $2"
`
const deploy = `
sh/sh/flanalyst.sh pull \
&& sh/deployCopyLibs.sh \
&& sh/sh/flanalyst.sh build \
&& sh/sh/flanalyst.sh dockerUp
`
const catLog = `
#!/bin/sh
while read line; do
  if jq -e . >/dev/null 2>&1 <<<"$line"; then
    echo $line | jq
  else
    echo $line
  fi
done < <(tail docker/data/log)
`
const catEvent = `
#!/bin/sh
`

var Commands = map[string]string{
	"test":           "sh/copyLibs.sh && ./gradlew test -Denv=local --parallel",
	"push":           "git push origin master -f",
	"testPushDeploy": testPushDeploy,
	"pull":           "git fetch --all && git reset --hard origin/master",
	"build":          "./gradlew :app:shadowJar --no-daemon",
	"dockerUp":       "docker-compose -f docker/docker-compose-app.yml up --build --detach",
	"dockerDown":     "docker-compose -f docker/docker-compose-app.yml down",
	"deploy":         deploy,
	"remote":         remote,
	"catLog":         catLog,
	"catEvent":       catLog}

var GetCommand = func(command string) (string, error) {
	sh := Commands[command]
	if sh == "" {
		return "", errors.New("unknown command " + command)
	}
	return sh, nil
}
