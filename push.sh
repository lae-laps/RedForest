#!/usr/bin/bash

# git script to automatically commit the repo with default config
# made by laelaps -> github.com/lae-laps/RedForest

if [ $EUID != 0 ]
then
    echo "Starting GIT automatic CLI commit"
    printf "=================================\n\n"
else
    echo "Avoid running GIT as root"
    printf "=========================\n\n"
fi

alias push="./push"

cd src 

make clean

cd ..

git add --all

git status

#git commit -m "automatic commit > cli > at ./push.sh" # <- auto commit

git commit -m "$1" # use 1rst argument as commit message

git status

git push origin HEAD:master

