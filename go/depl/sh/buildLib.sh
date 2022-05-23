. sh/.env &&
./gradlew build &&

cp lib/lib/build/libs/lib-1.0-SNAPSHOT.jar ../$lib-lib.jar &&
cp lib/lib/build/libs/lib-1.0-SNAPSHOT-sources.jar ../$lib-lib-sources.jar &&
cp lib/lib/build/libs/lib-1.0-SNAPSHOT-javadoc.jar ../$lib-lib-javadoc.jar &&

cp lib/testing/build/libs/testing-1.0-SNAPSHOT.jar ../$lib-testing.jar &&
cp lib/testing/build/libs/testing-1.0-SNAPSHOT-sources.jar ../$lib-testing-sources.jar &&
cp lib/testing/build/libs/testing-1.0-SNAPSHOT-javadoc.jar ../$lib-testing-javadoc.jar