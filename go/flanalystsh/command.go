package flanalystsh

const remote = `
#!/bin/bash
dir="$( cd -- "$( dirname -- "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )"
source $dir/.env
ssh -t $remoteUser@$remoteHost "\
cd $deployDir && \
sh/flanalyst.sh $2"
`

var Commands = map[string]string{
	"test":           "sh/copyLibs.sh && ./gradlew test -Denv=local --parallel",
	"push":           "git push origin master -f",
	"testPushDeploy": "sh/flanalyst.sh test && sh/flanalyst.sh push && sh/flanalyst.sh remote deploy",
	"pull":           "git fetch --all && git reset --hard origin/master",
	"build":          "./gradlew :app:shadowJar --no-daemon",
	"dockerUp":       "docker-compose -f docker/docker-compose-app.yml up --build --detach",
	"dockerDown":     "docker-compose -f docker/docker-compose-app.yml down",
	"deploy":         "sh/flanalyst.sh pull && sh/deployCopyLibs.sh && sh/flanalyst.sh build && sh/flanalyst.sh dockerUp",
	"remote":         remote}

var GetCommand = func(command string) string {
	return Commands[command]
}
