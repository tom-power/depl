depl.sh pull &&
depl.sh gradleCopyLibs &&
depl.sh shadowJar &&
depl.sh dockerUp
