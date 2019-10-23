.PHONY: all clean run verify

CXX=clang++
CXXFLAGS=-g -Wall -Wextra -Wshadow -Wnon-virtual-dtor -pedantic -D_REENTRANT -std=c++14 -O0
DIR_BIN=bin
DIR_TESTCASES_IN=testcases/input
DIR_TESTCASES_OUT=testcases/output
EXE_APP=exe

all: $(DIR_BIN)/$(EXE_APP)

clean:
	rm -rf $(DIR_BIN)

run: $(DIR_BIN)/$(EXE_APP)
	@$(DIR_BIN)/$(EXE_APP)

verify: $(DIR_BIN)/$(EXE_APP)
	@VerifyHrAnswer.sh

$(DIR_BIN)/$(EXE_APP): $(DIR_BIN)/ans.o
	@mkdir -p $(@D)
	$(CXX) $(CXXFLAGS) $< -o $@

$(DIR_BIN)/ans.o: ans.cpp
	@mkdir -p $(@D)
	$(CXX) -c $(CXXFLAGS) $< -o $@
