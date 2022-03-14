#!/usr/bin/env bash

# wating for db's no longer required
# echo "waiting 15 seconds for dependencies just in case"
# sleep 15

faust -A main worker -l info
