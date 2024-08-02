package depl

import (
	_ "embed"
)

var (
	//go:embed scripts/dockerCatEvent.sh
	dockerCatEvent string

	//go:embed scripts/dockerCatLog.sh
	dockerCatLog string

	//go:embed scripts/dockerUp.sh
	dockerUp string

	//go:embed scripts/dockerUpDev.sh
	dockerUpDev string

	//go:embed scripts/gitPull.sh
	gitPull string

	//go:embed scripts/gitPush.sh
	gitPush string
	//go:embed scripts/gitSubmoduleUpdate.sh
	gitSubmoduleUpdate string

	//go:embed scripts/gradleBuild.sh
	gradleBuild string

	//go:embed scripts/gradleBuildLib.sh
	gradleBuildLib string

	//go:embed scripts/gradleCopyLibs.sh
	gradleCopyLibs string

	//go:embed scripts/gradleCopyLibsAll.sh
	gradleCopyLibsAll string

	//go:embed scripts/gradleDeploy.sh
	gradleDeploy string

	//go:embed scripts/gradleTest.sh
	gradleTest string

	//go:embed scripts/gradleShadowJar.sh
	gradleShadowJar string

	//go:embed scripts/remote.sh
	remote string
)
