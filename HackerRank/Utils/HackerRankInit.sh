#!/usr/bin/env bash

# Name:          HackerRankInit
# Description:   Calls the HackerRankInit python script that downloads the challenge description,
#                sample testcases (input/output) and prepares a skeleton answer file.
# Options:       (positional, mandatory): a link to the challenge's page
# Created on:    16.03.2020
# Last modified: 20.03.2020
# Author:        Adam Graliński (adam@gralin.ski)
# License:       MIT


# This script will be called from a symlink. The following resolves the actual script location,
# even when called via a symlink:
SRC="${BASH_SOURCE[0]}"
while [ -h "${SRC}" ]; do
  DIR="$(cd -P "$(dirname "${SRC}")" >/dev/null 2>&1 && pwd)"
  SRC="$(readlink "${SRC}")"
  [[ ${SRC} != /* ]] && SRC="${DIR}/${SRC}"
done
RESOLVED_SCRIPT_DIR="$(cd -P "$(dirname "${SRC}")" >/dev/null 2>&1 && pwd)"

echo "Resolved script dir: ${RESOLVED_SCRIPT_DIR}"

PYTHON_SCRIPT_LOCATION="${RESOLVED_SCRIPT_DIR}/HackerRankInit/hr_init.py"

if [ -f "${PYTHON_SCRIPT_LOCATION}" ]; then
  # Note: ${*} needs to be passed without surrounding quotes,
  # otherwise multiple parameters would be passed as one string.
  python3 "${PYTHON_SCRIPT_LOCATION}" ${*}
else
  >&2 printf "Error: file %s not found.\nAborting.\n" "${PYTHON_SCRIPT_LOCATION}"
  exit 1
fi
