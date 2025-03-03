#!/usr/bin/env bash
# get some test data
mkdir -p testdata
pushd testdata
  wget https://raw.githubusercontent.com/jupyter/notebook/refs/heads/main/package.json
  mksquashfs package.json test.squashfs
popd

# build and run the example
go run ./main.go

