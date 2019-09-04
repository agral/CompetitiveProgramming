# A Makefile for building a solution to competitive programming task in C++
# Author: Adam Graliński (adam@gralin.ski)
# License: CC0

PROGRAM_NAME="ans.out"
CXX=clang++
CXXFLAGS=-g -c -Wall -Wextra -Wpedantic -D_REENTRANT -std=c++17
LDFLAGS=
LDLIBS=

DIR_SRC=.
DIR_BIN=.
DIR_OBJ=.

SOURCES=$(shell find ${DIR_SRC} -name "*.cpp")
OBJECTS=$(patsubst ${DIR_SRC}/%.cpp,${DIR_OBJ}/%.o,${SOURCES})
TESTCASES=$(shell find ${DIR_SRC} -name "*.testcase")

dir_guard=@mkdir -p $(@D)

.PHONY: all clean run

# Makes the parameterless  invocation build the app, as long as this line is the first rule:
all: ${DIR_BIN}/${PROGRAM_NAME}

clean:
	@rm -fv ${OBJECTS}
	@rm -fv ${DIR_BIN}/${PROGRAM_NAME}

run: ${DIR_BIN}/${PROGRAM_NAME}
	@${DIR_BIN}/${PROGRAM_NAME}

test: ${DIR_BIN}/${PROGRAM_NAME}
#	@for testcase in ${TESTCASES[@]}; do echo "${testcase}"; done
#	$(shell for tc in *.testcase; do ${DIR_BIN}/${PROGRAM_NAME} < ${tc}; done)
	${DIR_BIN}/${PROGRAM_NAME} < 0.testcase
	${DIR_BIN}/${PROGRAM_NAME} < 1.testcase
	${DIR_BIN}/${PROGRAM_NAME} < 2.testcase
	${DIR_BIN}/${PROGRAM_NAME} < 3.testcase


### Auto-rules: ###
# Creates an application's executable by linking all the app object files:
${DIR_BIN}/${PROGRAM_NAME} : ${OBJECTS}
	@${dir_guard}
	${CXX} ${LDFLAGS} -o ${@} ${OBJECTS} ${LDLIBS}

# Compiles a single file of application's source code to an object file:
${DIR_OBJ}/%.o : ${DIR_SRC}/%.cpp
	@${dir_guard}
	${CXX} ${CXXFLAGS} ${<} -o ${@}
