ROOT_DIR="$(pwd)"

for goModFile in $(find . -name '*.mod'); do
  currentGoRootDirectory="$(dirname "$goModFile")"
  echo ""
  echo "### $currentGoRootDirectory ###"
  echo ""

  cd "$currentGoRootDirectory" || exit 1

  # synch packages and clear test cache
  go mod tidy
  go clean -testcache

  # run tests and create coverage report
  go test -race -coverprofile=coverage.txt -covermode=atomic

  # check if all tests were successfull, otherwise return with error
  if [ $? -ne 0 ]; then
    rm -f cover.out
    exit 1
  fi

  cd "$ROOT_DIR" || exit 1
done
