#!/bin/bash
set -e

server=cargo-server
goos="" # darwin,linux
goarch="" # amd64,arm64

git_hash=$(git rev-list -1 HEAD)
git_branch=$(git branch --show-current)
build_time=$(date "+%Y-%m-%d %H:%M:%S %Z")

set -x
env CGO_ENABLED=0 GOOS="$OS" GOARCH="$ARCH" go build -o "$server" -ldflags "-s -w -X main.gitHash=$git_hash -X main.gitBranch=$git_branch -X 'main.buildTime=$build_time'"
