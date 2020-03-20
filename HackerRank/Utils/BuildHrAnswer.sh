#!/usr/bin/env bash

# Name:          BuildHrAnswer
# Description:   Builds answers to HackerRank's challenge for a number of languages, if present.
# Options:       None, the script needs to be run from the challenge's root directory.
#                (e.g. /path/to/hacker/rank/dir/Practice/Algorithms/Search/CountLuck)
# Created on:    18.03.2020
# Last modified: 20.03.2020
# Author:        Adam Graliński (adam@gralin.ski)
# License:       MIT

DIR_ORIGIN="${PWD}"
CHALLENGE_NAME="${PWD##*/}"

abort() {
  >&2 echo "Aborting."
  exit 1
}


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
fi

SOLUTION_JAVA="${DIR_ORIGIN}/java/${CHALLENGE_NAME}.java"
if [ -f "${SOLUTION_JAVA}" ]; then
  echo "Building JAVA solution..."
  JAVA_COMPILER=javac
  if ${JAVA_COMPILER} -classpath "${DIR_ORIGIN}/java/" "${SOLUTION_JAVA}"; then
    echo "OK"
  else
    abort
  fi
fi
