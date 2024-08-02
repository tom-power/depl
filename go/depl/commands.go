package depl

var Commands = map[string]string{
	"dockerUp":           dockerUp,
	"dockerUpDev":        dockerUpDev,
	"dockerCatLog":       dockerCatLog,
	"dockerCatEvent":     dockerCatEvent,
	"gitPull":            gitPull,
	"gitPush":            gitPush,
	"gitSubmoduleUpdate": gitSubmoduleUpdate,
	"gradleBuild":        gradleBuild,
	"gradleBuildLib":     gradleBuildLib,
	"gradleCopyLibs":     gradleCopyLibs,
	"gradleCopyLibsAll":  gradleCopyLibsAll,
	"gradleDeploy":       gradleDeploy,
	"gradleTest":         gradleTest,
	"gradleShadowJar":    gradleShadowJar,
	"remote":             remote,
}
