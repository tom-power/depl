depl.sh pull &&
depl.sh deployCopyLibs &&
depl.sh shadowJar &&
depl.sh dockerUp