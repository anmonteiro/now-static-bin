#!/usr/bin/env sh

set -eo pipefail

root_path=$PWD

# Start in examples/go-graphql/ even if run from root directory
cd "$(dirname "$0")"

docker build . --tag go-example
docker rm go-example || true
docker create --name go-example go-example
docker cp go-example:/root/go/app/main.exe .

cd $root_path