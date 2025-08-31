#!/bin/bash

OLD_PWD=$(pwd)
SHELL_FOLDER=$(
    cd "$(dirname "$0")" || exit
    pwd
)
PROJECT_FOLDER=$SHELL_FOLDER/../..

cd "$SHELL_FOLDER" || exit >/dev/null 2>&1

# shellcheck source=/dev/null
source "$PROJECT_FOLDER/scripts/base/env.sh"

check_gh_exist

VERSION="v$(go run main-version.go)"

TITLE="Release $VERSION"

NOTES_FILE="$PROJECT_FOLDER/release_notes.md"

if [ ! -f "$NOTES_FILE" ]; then
    echo "Error: Release notes file '$NOTES_FILE' not found."
    exit 1
fi

PROJECT_NAME="windows-os-info"

RELEASE_DIR="$PROJECT_FOLDER/release"
mkdir -p "$RELEASE_DIR"

rm -rf "${RELEASE_DIR:?}"/*

echo "Compiling binaries for multiple platforms..."

cd "$PROJECT_FOLDER" || exit >/dev/null 2>&1

# Windows 32-bit
COMMAND="GOOS=windows GOARCH=386 go build -ldflags '-w -s' -trimpath -o $RELEASE_DIR/${PROJECT_NAME}.exe"
echo exec: "$COMMAND"
if eval "$COMMAND"; then
    echo "Windows 32-bit binary compiled successfully."
else
    echo "Failed to compile Windows 32-bit binary."
    exit 1
fi

echo "Compilation completed."

echo "Generated binaries:"
ls -lh "$RELEASE_DIR"

echo "Creating release on GitHub..."

# shellcheck disable=SC2034
EXISTING_RELEASE=$(gh release view "$VERSION" 2>/dev/null)
# shellcheck disable=SC2181
if [ $? -eq 0 ]; then
    echo "Release $VERSION already exists. Aborting."
    exit 1
fi

COMMAND="gh release create $VERSION --title \"$TITLE\" --notes-file \"$NOTES_FILE\" $RELEASE_DIR/*"
echo exec: "$COMMAND"
if eval "$COMMAND"; then
    echo "Release $VERSION successfully created and binaries uploaded."
else
    echo "Failed to create GitHub release."
    exit 1
fi

cd "$OLD_PWD" || exit >/dev/null 2>&1
