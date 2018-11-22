#!/usr/bin/env sh

docker build . --tag rust-example
docker create --name rust-example rust-example
docker cp rust:/server.exe .