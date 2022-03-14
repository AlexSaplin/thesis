#!/bin/sh

# startup backoff no longer required
# echo "sleeping for 10 sec for dependencies to go up"
# sleep 10

python -u {{ server_name }}
