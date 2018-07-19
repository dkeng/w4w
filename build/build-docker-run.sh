#!/bin/bash
cd ..
docker build -t keng/w4w:latest .
docker stop w4w
docker rm w4w
docker run --name w4w -d keng/w4w:latest