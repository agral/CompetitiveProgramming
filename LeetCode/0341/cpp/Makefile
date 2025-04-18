.PHONY: all run answer clean unittest

UNITTEST_BINARY_NI:=run_unittest_nested_integer
ANSWER_BINARY:=test

all: answer unittest

answer:
	g++ -std=c++20 -o ${ANSWER_BINARY} nested_integer.cpp lc0341.cpp

clean:
	@rm -f ${UNITTEST_BINARY_NI}
	@rm -f ${ANSWER_BINARY}

unittest:
	g++ -std=c++20 -o ${UNITTEST_BINARY_NI} nested_integer.cpp nested_integer_test.cpp
