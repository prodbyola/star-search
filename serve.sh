#!/bin/sh

. ~/.profile

if [ "$1" = "docs" ]; then
    cd docs
    quasar dev

else
    sudo service postgresql start
    sudo service redis-server start
    cd app && air

fi