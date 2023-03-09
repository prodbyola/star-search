#!/bin/sh

. ~/.profile

cd app
go build -o starsearch main.go

mkdir ../deploy
mv starsearch ../deploy/starsearch
