#!/bin/bash

auto() {
    [[ -z "$1" ]] && echo "Please enter a commit message:";
    typeset msg="$( [[ -n "$1" ]] && echo "$*" || echo $(head -1) )";
    date;
    git pull;
    git add .;
    git commit -m "$msg";
    git push;
    date
}

auto