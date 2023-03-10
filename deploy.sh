#!/bin/sh

sh build.sh app

cd deploy
git add .
git commit -m "One"
git push origin main