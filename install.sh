#!/bin/sh

if [ "$1" = "db" ]; then

    sudo sh -c 'echo "deb http://apt.postgresql.org/pub/repos/apt $(lsb_release -cs)-pgdg main" > /etc/apt/sources.list.d/pgdg.list'
    wget --quiet -O - https://www.postgresql.org/media/keys/ACCC4CF8.asc | sudo apt-key add -
    sudo apt-get update
    sudo apt-get -y install postgresql

elif [ "$1" = "air" ]; then
    # binary will be $(go env GOPATH)/bin/air
    curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b $(go env GOPATH)/bin

    # or install it into ./bin/
    curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s

    air -v

elif [ "$1" = "docs" ]; then
    cd docs
    yarn install

elif [ "$1" = "all" ]; then
    sh install.sh db
    sh install.sh air
    sh install.sh docs

else
    echo "What do you wish to install? db | air | docs | all";

fi