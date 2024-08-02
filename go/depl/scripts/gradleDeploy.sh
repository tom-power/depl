depl.sh gitPull &&
depl.sh gradleCopyLibs &&
depl.sh gradleShadowJar &&
depl.sh dockerUp
