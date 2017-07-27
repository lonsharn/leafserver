#!/bin/bash

export GOPATH=`pwd`
go build -o bin/server -gcflags "-N -l" src/server/main.go
go build -o bin/client -gcflags "-N -l" src/test/client.go
