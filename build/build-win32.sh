#!/bin/bash
rm -rf ../bin/*
mkdir ../bin/config
cp -a ../config/ ../bin/config
cp -a ../src/templates ../bin/
CGO_ENABLED=0 GOOS=windows GOARCH=386 go build -o ../bin/w4w32.exe ../src