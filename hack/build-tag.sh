#!/bin/sh

set -e

DIR=$(cd $(dirname $0) && pwd)
cd $DIR/..

VERSION=$1
PYTHON_VERSION=$2
PYTHON_STANDALONE_VERSION=$3

if [ "$VERSION" = "" ]; then
  echo "missing version (e.g. v1.0.0)"
  exit 1
fi

if [ "$PYTHON_VERSION" = "" ]; then
  echo "missing python version"
  exit 1
fi

if [ "$PYTHON_STANDALONE_VERSION" = "" ]; then
  echo "missing python-standalone version"
  exit 1
fi

if [ ! -z "$(git status --porcelain)" ]; then
  echo "working directory is dirty!"
  exit 1
fi

go run ./python/generate --python-standalone-version=$PYTHON_STANDALONE_VERSION --python-version $PYTHON_VERSION
go run -tags ratatoskr_embed ./pip/generate

echo "checking out temporary branch"
git checkout --detach
git add -f python/internal/data
git add -f pip/internal/data
git commit -m "release $VERSION (Python $PYTHON_VERSION, python-build-standalone $PYTHON_STANDALONE_VERSION)"
git tag -f $VERSION
git checkout -
