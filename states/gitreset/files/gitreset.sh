#!/usr/bin/env bash

if [[ `git status --porcelain` ]]; then
salt -L "master.local" beacons.disable
  git reset --hard origin
  git clean -f
sleep 3
salt -L "master.local" beacons.enable
else
  echo "no change"
fi
