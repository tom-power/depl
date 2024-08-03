if [[ -f sh/build.sh ]]; then
  sh/build.sh
else
  ./gradlew build -Denv=local --parallel
fi
