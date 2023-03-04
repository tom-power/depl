if [[ -f sh/test.sh ]]; then
  sh/test.sh
else
  ./gradlew test -Denv=local --parallel
fi
