package depl

var Commands = map[string]string{
	"dockerCatEvent":     dockerCatEvent,
	"dockerCatLog":       dockerCatLog,
	"dockerUp":           dockerUp,
	"dockerUpDev":        dockerUpDev,
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
