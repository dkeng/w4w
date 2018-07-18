#!/bin/bash
cp -a ../config ../bin/
CGO_ENABLED=0 GOOS=windows GOARCH=386 go build -o ../bin/w4w32.exe ../src