#!/bin/bash

export GOPATH=$(pwd)

mkdir bin
cd bin

go build testplugin
go build client