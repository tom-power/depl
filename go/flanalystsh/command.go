package flanalystsh

const remote = `
#!/bin/bash
dir="$( cd -- "$( dirname -- "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )"
source $dir/.env
ssh -t $remoteUser@$remoteHost "\
cd $deployDir && \
sh/$1.sh $2"
`

var Commands = map[string]string{
	"test":           "sh/copyLibs.sh && ./gradlew test -Denv=local --parallel",
	"push":           "git push origin master -f",
	"testPushDeploy": "sh/test.sh && sh/push.sh && sh/remote/run.sh deploy",
	"pull":           "git fetch --all && git reset --hard origin/master",
	"dockerUp":       "docker-compose -f docker/docker-compose-app.yml up --build --detach",
	"dockerDown":     "docker-compose -f docker/docker-compose-app.yml down",
	"deploy":         "sh/pull.sh && sh/deployCopyLibs.sh && sh/build.sh && sh/dockerUp.sh",
	"build":          "./gradlew :app:shadowJar --no-daemon",
	"remote":         remote}

var GetCommand = func(command string) string {
	return Commands[command]
}
