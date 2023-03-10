#!/bin/sh

. ~/.profile

if [ "$1" = "app" ]; then
    
    cd app
    go build -o starsearch main.go

    # mkdir ../deploy
    mv starsearch ../deploy/starsearch

elif [ "$1" = "docs" ]; then
    cd docs
    quasar build
    cd dist
    explorer.exe .
fi
