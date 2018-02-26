#!/usr/bin/env bash

if [[ `git status --porcelain` ]]; then
  git reset --hard origin
else
  echo "no change"
fi