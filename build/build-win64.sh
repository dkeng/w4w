#!/bin/bash
cp -a ../config ../bin/
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o ../bin/w4w64.exe ../src