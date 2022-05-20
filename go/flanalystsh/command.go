package flanalystsh

const testPushDeploy = `
sh/flanalyst.sh test \
&& sh/flanalyst.sh push \
&& sh/flanalyst.sh remote deploy
`

const remote = `
. sh/.env &&
ssh -t $remoteUser@$remoteHost "cd $deployDir && sh/flanalyst.sh $2"
`
const deploy = `
sh/flanalyst.sh pull \
&& sh/deployCopyLibs.sh \
&& sh/flanalyst.sh build \
&& sh/flanalyst.sh dockerUp
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
	"remote":         remote}

var GetCommand = func(command string) string {
	return Commands[command]
}
