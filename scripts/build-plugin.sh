#!/bin/bash

export GOPATH=$(pwd)

mkdir bin
cd bin

go build $PLUGIN_NAME