#!/bin/bash

# If HEAD is an annotated tag and the working tree is not dirty, return the tag.
if git describe --exact-match HEAD > /dev/null 2>&1 && git diff --quiet; then
    echo $(git describe --abbrev=0 --tags)
    exit 0
fi

# If the working tree is clean, print the short commit, otherwise print 'dirty'.
if git diff --quiet; then
    echo $(git rev-parse --short HEAD)
    exit 0
else
    echo 'dirty'
    exit 0
fi
