#!/bin/sh

docker run --rm -v "$PWD":/go/src/wonderboy -w /go/src/wonderboy golang:1.17-stretch go run main.go