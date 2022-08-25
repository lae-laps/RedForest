#!/usr/bin/bash
cd src 
make clean
cd ..
git add --all
git status
#git commit -m "automatic commit > cli > at ./push.sh"
git commit -m "$1"
git status
git push origin HEAD:master
