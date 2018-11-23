#!/usr/bin/env sh

set -eo pipefail

root_path=$PWD

# Start in examples/rust/ even if run from root directory
cd "$(dirname "$0")"

docker build . --tag rust-example
docker rm rust-example || true
docker create --name rust-example rust-example
docker cp rust-example:/server.exe .

cd $root_path