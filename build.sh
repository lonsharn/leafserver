#!/bin/bash

export GOPATH=`pwd`
go build -o bin/server -gcflags "-N -l" src/server/main.go
