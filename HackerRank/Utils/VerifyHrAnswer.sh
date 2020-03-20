#!/usr/bin/env bash

# Name:          VerifyHrAnswer
# Description:   Executes an answer to a HackerRank's challenge against all the testcases
#                associated with it.
# Options:       None, the script needs to be run from the challenge's root directory.
#                (e.g. /path/to/hacker/rank/dir/Practice/Algorithms/Search/CountLuck)
# Created on:    20.10.2019
# Last modified: 20.03.2020
# Author:        Adam Graliński (adam@gralin.ski)
# License:       MIT

DIR_ORIGIN="${PWD}"
CHALLENGE_NAME="${PWD##*/}"
DIR_TESTCASES_IN="testcases/input"
DIR_TESTCASES_OUT="testcases/output"

log() {
  echo -e "${@}"
}

loge() {
  >&2 echo -e "${@}"
}

logv() {
  if [ -n "${BE_VERBOSE}" ]; then
    echo -e "${@}"
  fi
}

abort() {
  >&2 echo "Aborting."
  exit 1
}

# run_and_compare `executable` `testcase_input_file` `expected_output_file`
run_and_compare() {
  if [ "${#}" -ne 3 ]; then
    loge "Critical: run_and_compare() called with ${#} arguments (3 required)."
    abort
  fi
  # Note: ${1} should go without surrounding quotes.
  diff "${3}" <(${1} <"${2}")
  if [ "${?}" -ne 0 ]; then
    loge "Testcase ${2} FAILED."
    abort
  fi
}

# run_test_suite `executable` `(lang)`
run_test_suite() {
  if [ "${#}" -lt 2 ]; then
    loge "Critical: run_test_suite() called with ${#} arguments (2 required)."
    abort
  fi

  log "Running testcases for ${2} solution..."
  for tc in "${TESTCASES[@]}"; do
    logv "TC ${tc}"
    expected_output_file="${tc//input/output}"
    if [ ! -f "${expected_output_file}" ]; then
      loge "Error: could not find expected answer file for TC ${tc}"
      loge "(file: ${expected_output_file} does not exist)."
      abort
    fi
    run_and_compare "${1}" "${tc}" "${expected_output_file}"
    logv "${tc}: OK\n"
  done
  log "Testcases PASSED for ${2}."
}

### Handles the parameters: ###
while [ "${#}" -gt 0 ]; do
  case "${1}" in
    -v|--verbose)
      BE_VERBOSE=1
      logv "Verbose output is set."
    ;;
  esac
  shift
done

# Verifies that the testcases directory exists, stores all testcase files in `TESTCASES` array:
if [ ! -d "${DIR_TESTCASES_IN}" ] || [ ! -d "${DIR_TESTCASES_OUT}" ]; then
  loge "Error: testcases have not been found at ${DIR_TESTCASES_IN}, ${DIR_TESTCASES_OUT}."
  abort
fi
TESTCASES=()
while IFS=  read -r -d $'\0'; do
    TESTCASES+=("$REPLY")
done < <(find "${DIR_TESTCASES_IN}" -type f -name "*.txt" -print0)


unset ANY_TESTS_RAN
EXECUTABLE_CPP="${DIR_ORIGIN}/cpp/${CHALLENGE_NAME}.exe"
if [ -f "${EXECUTABLE_CPP}" ]; then
  run_test_suite "${EXECUTABLE_CPP}" "cpp"
  ANY_TESTS_RAN=1
else
  logv "Skipping tests for cpp - solution file not found."
fi

if [ -f "${DIR_ORIGIN}/java/${CHALLENGE_NAME}.class" ]; then
  run_test_suite "java -classpath ${DIR_ORIGIN}/java ${CHALLENGE_NAME}" "java"
  ANY_TESTS_RAN=1
else
  logv "Skipping tests for java - solution file not found."
fi

if [ -z "${ANY_TESTS_RAN}" ]; then
  loge "No tests have been executed."
  abort
fi
