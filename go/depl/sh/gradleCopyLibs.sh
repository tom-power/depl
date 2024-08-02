. sh/.env &&

if [[ ! -d ./libs/ ]]; then
  mkdir ./libs/
fi

cp -p $libPath/lib/lib/build/libs/lib-1.0-SNAPSHOT.jar ./libs/lib.jar
