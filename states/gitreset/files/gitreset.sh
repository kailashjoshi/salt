#!/usr/bin/env bash

if [[ `git status --porcelain` ]]; then
  git reset --hard origin
  git clean -f
else
  echo "no change"
fi