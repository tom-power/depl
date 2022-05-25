package depl

import (
	_ "embed"
)

var (
  //go:embed sh/buildLib.sh
  buildLib string

  //go:embed sh/copyLibs.sh
  copyLibs string

  //go:embed sh/copyLibsNested.sh
  copyLibsNested string

  //go:embed sh/test.sh
  test string

  //go:embed sh/push.sh
  push string

  //go:embed sh/remote.sh
  remote string

  //go:embed sh/pull.sh
  pull string

  //go:embed sh/deployCopyLibs.sh
  deployCopyLibs string

  //go:embed sh/shadowJar.sh
  shadowJar string

  //go:embed sh/dockerUp.sh
  dockerUp string

  //go:embed sh/deploy.sh
  deploy string

  //go:embed sh/catEvent.sh
  catEvent string

  //go:embed sh/catLog.sh
  catLog string
)