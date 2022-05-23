sh/depl.sh pull &&
sh/depl.sh deployCopyLibs &&
sh/depl.sh shadowJar &&
sh/depl.sh dockerUp