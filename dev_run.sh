#!/bin/bash

# If getting message zsh permission denied to run .sh:
# chmod -x ./script_name.sh

# get the location of this script
SCRIPT_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"

source ${SCRIPT_DIR}/env.sh
if [ -e ${SCRIPT_DIR}/env-local.sh ]; then
    # env-local should be ignored by git, and used to override any 'env.sh' settings for local development
  source ${SCRIPT_DIR}/env-local.sh
fi

go run ${SCRIPT_DIR}/main.go