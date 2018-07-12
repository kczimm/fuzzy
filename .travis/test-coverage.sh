#!/bin/bash

go test -v -covermode=count -coverprofile=coverage.out
$HOME/gopath/bin/goveralls -coverprofile=coverage.out -service=travis-ci -repotoken $COVERALLS_TOKEN