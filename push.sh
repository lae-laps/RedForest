#!/usr/bin/bash
cd src 
make clean
cd ..
git add --all
git status
git commit -m "automatic commit > cli > at ./push.sh"
git push --set-upstream origin master
