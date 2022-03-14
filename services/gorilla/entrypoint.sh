#!/bin/sh
set -xeuo pipefail

# echo "sleeping for 10 sec for postgres to go up"
# sleep 10

./app $@
