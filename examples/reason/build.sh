#!/usr/bin/env sh

set -eo pipefail

root_path=$PWD

# Start in examples/rust/ even if run from root directory
cd "$(dirname "$0")"

docker build . --tag reason-example
docker rm reason-example || true
docker create --name reason-example reason-example
docker cp reason-example:/app/main.exe .

cd $root_path