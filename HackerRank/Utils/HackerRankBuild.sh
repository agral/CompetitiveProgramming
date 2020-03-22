#!/usr/bin/env bash

# Name:          HackerRankBuild
# Description:   Builds answers to HackerRank's challenge for a number of languages, if present.
# Options:       --cpp|--java
# Note:          The script needs to be run from the challenge's root directory.
#                (e.g. /path/to/hacker/rank/dir/Practice/Algorithms/Search/CountLuck)
# Created on:    18.03.2020
# Last modified: 22.03.2020
# Author:        Adam Graliński (adam@gralin.ski)
# License:       MIT

DIR_ORIGIN="${PWD}"
CHALLENGE_NAME="${PWD##*/}"

abort() {
  >&2 echo "Aborting."
  exit 1
}

print_help() {
  echo "HackerRankBuild - build solutions to a HackerRank challenge for selected languages."
  echo "    Options: [-h|--help] [--cpp|java]"
  echo "        --help - print this help message and exit. Alias: -h."
  echo "        --cpp - add C++ solution to the build list"
  echo "        --java - add Java solution to the build list"
  echo "Returns 0 if all the builds were successful, 1 otherwise."
  echo "Note: the script needs to be run from the challenge's root directory."
  echo "Author: Adam Graliński (adam@gralin.ski)"
  echo "License: MIT"
}

unset ANY_SOLUTION_SELECTED
while [ "${#}" -gt 0 ]; do
  case "${1}" in
    -h|--help)
      print_help
      exit 0
    ;;
    --cpp)
      ANY_SOLUTION_SELECTED=1
      SOLUTION_CPP="${DIR_ORIGIN}/cpp/${CHALLENGE_NAME}.cpp"
      if [ -f "${SOLUTION_CPP}" ]; then
        echo "Building CPP solution..."
        CXX=clang++
        CXXFLAGS=(-Wall -Wextra -Wshadow -Wnon-virtual-dtor -pedantic -D_REENTRANT "-std=c++17" -g -O2)
        if ${CXX} "${CXXFLAGS[@]}" "${SOLUTION_CPP}" -o "${DIR_ORIGIN}/cpp/${CHALLENGE_NAME}.exe"; then
          echo "OK"
        else
          abort
        fi
      else
        >&2 echo "Error: CPP solution file not found."
        abort
      fi
    ;;
    --java)
      ANY_SOLUTION_SELECTED=1
      SOLUTION_JAVA="${DIR_ORIGIN}/java/${CHALLENGE_NAME}.java"
      if [ -f "${SOLUTION_JAVA}" ]; then
        echo "Building JAVA solution..."
        JAVA_COMPILER=javac
        if ${JAVA_COMPILER} -classpath "${DIR_ORIGIN}/java/" "${SOLUTION_JAVA}"; then
          echo "OK"
        else
          abort
        fi
      else
        >&2 echo "Error: JAVA solution file not found."
        abort
      fi
    ;;
  esac
  shift
done

if [ -z "${ANY_SOLUTION_SELECTED}" ]; then
  >&2 echo "Nothing to be built selected."
  abort
fi
