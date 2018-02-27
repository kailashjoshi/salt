#!/usr/bin/env bash

if [[ `git status --porcelain` ]]; then
  salt-call beacons.disable
  git reset --hard origin
  git clean -f
  sleep 3
  salt-call beacons.enable
else
  echo "no change"
fi
