package depl

import (
	_ "embed"
)

var (
	//go:embed sh/dockerCatEvent.sh
	dockerCatEvent string

	//go:embed sh/dockerCatLog.sh
	dockerCatLog string

	//go:embed sh/dockerUp.sh
	dockerUp string

	//go:embed sh/dockerUpDev.sh
	dockerUpDev string

	//go:embed sh/gitPull.sh
	gitPull string

	//go:embed sh/gitPush.sh
	gitPush string
	//go:embed sh/gitSubmoduleUpdate.sh
	gitSubmoduleUpdate string

	//go:embed sh/gradleBuild.sh
	gradleBuild string

	//go:embed sh/gradleBuildLib.sh
	gradleBuildLib string

	//go:embed sh/gradleCopyLibs.sh
	gradleCopyLibs string

	//go:embed sh/gradleCopyLibsAll.sh
	gradleCopyLibsAll string

	//go:embed sh/gradleDeploy.sh
	gradleDeploy string

	//go:embed sh/gradleTest.sh
	gradleTest string

	//go:embed sh/gradleShadowJar.sh
	gradleShadowJar string

	//go:embed sh/remote.sh
	remote string
)
