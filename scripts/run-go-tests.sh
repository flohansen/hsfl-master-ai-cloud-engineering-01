ROOT_DIR="$(pwd)"

for goModFile in $(find ./src -name '*.mod'); do
  currentGoRootDirectory="$(dirname "$goModFile")"
  echo ""
  echo "### $currentGoRootDirectory ###"
  echo ""

  cd "$currentGoRootDirectory" || exit 1

  go mod tidy
  go test ./...

  if [ $? -ne 0 ]; then
    exit 1
  fi

  cd "$ROOT_DIR" || exit 1
done
