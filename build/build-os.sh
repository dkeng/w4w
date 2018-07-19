#!/bin/bash
rm -rf ../bin/*
mkdir ../bin/config
cp -a ../config/ ../bin/config
cp -a ../src/templates ../bin/
CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o ../bin/w4w.a ../src