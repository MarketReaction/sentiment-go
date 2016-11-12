#!/usr/bin/env bash

docker run -v "$PWD/analyser":/usr/src/myapp -w /usr/src/myapp golang \
        go get -t -d -v ./... \
        go test -v ./...