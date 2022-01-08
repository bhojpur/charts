#!/bin/bash

version="v0.0.6"

function pre_check() {
  info=$(diff -u <(echo -n) <(format))
  if [ -n "$info" ]; then
    echo ">> Make sure you have formatted the codebase before committing."
    echo ">> Use 'git status' command to show what have changed."
    exit 1
  fi
}

function format() {
    go fmt ./... && ls -d */ | xargs goimports -w
}

function test() {
    go test -v ./pkg/charts/...
    go test -v ./pkg/opts/...
}

function release() {
    pre_check

    echo ">> Building tag..."
    git tag -a $version -m "Release: $version"

    echo ">> Pushing tag to the remote...."
    git push origin $version
}

function help() {
    echo "$0 --cmd [format|test|release]"
}

if [ "$1" == "" ]; then
    help
elif [ "$1" == "format" ];then
    format
elif [ "$1" == "test" ];then
    test
elif [ "$1" == "release" ];then
    release
else
    help
fi
