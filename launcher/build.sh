#!/usr/bin/env sh

set -eo pipefail

root_path=$PWD

# Start in launcher/ even if run from root directory
cd "$(dirname "$0")"

docker build . --tag now-static-bin-launcher
docker rm now-static-bin-launcher || true
docker create --name now-static-bin-launcher now-static-bin-launcher
docker cp now-static-bin-launcher:/root/go/app/launcher ../dist

cd $root_path