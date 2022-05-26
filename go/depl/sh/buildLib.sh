. sh/.env &&
./gradlew build &&

cp lib/lib/build/libs/lib-1.0-SNAPSHOT.jar $libPath$lib-lib.jar &&
cp lib/lib/build/libs/lib-1.0-SNAPSHOT-sources.jar $libPath$lib-lib-sources.jar &&
cp lib/lib/build/libs/lib-1.0-SNAPSHOT-javadoc.jar $libPath$lib-lib-javadoc.jar &&

cp lib/testing/build/libs/testing-1.0-SNAPSHOT.jar $libPath$lib-testing.jar &&
cp lib/testing/build/libs/testing-1.0-SNAPSHOT-sources.jar $libPath$lib-testing-sources.jar &&
cp lib/testing/build/libs/testing-1.0-SNAPSHOT-javadoc.jar $libPath$lib-testing-javadoc.jar