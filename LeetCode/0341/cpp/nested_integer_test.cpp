#include <iostream>

#include "nested_integer.hpp"

int main() {
    int numGood = 0, numBad = 0;

    // Testcase 1: plain integers.
    struct IntegerTestcase {
        NestedInteger ni;
        int expected;
    };
    IntegerTestcase testcases[] = {
        NestedInteger{1}, 1,
        NestedInteger{9}, 9,
        NestedInteger{42}, 42,
    };

    for (const IntegerTestcase& tc: testcases) {
        int actual = tc.ni.getInteger();
        if (actual != tc.expected) {
            std::cout << "Testcase " << tc.expected << " failed. Got: " << actual
                << ", want: " << tc.expected << "\n";
            ++numBad;
        } else {
            ++numGood;
        }
    }

    // Testcase 2: vectors
    struct VectorTestcase {
        NestedInteger ni;
        std::vector<int> expected;
    };
    NestedInteger ni1{1}, ni2{2}, ni3{3}, ni4{4};
    VectorTestcase vtestcases[] = {
        std::vector<NestedInteger>{ni1, ni2, ni3, ni4}, std::vector<int>{1, 2, 3, 4},
        std::vector<NestedInteger>{ni4, ni2}, std::vector<int>{4, 2},
    };
    for (const VectorTestcase& tc: vtestcases) {
        std::vector<NestedInteger> actual = tc.ni.getList();
        if (actual.size() != tc.expected.size()) {
            std::cout << "Testcase failed; list size mismatch.\n";
            ++numBad;
        }
        else {
            for (int i = 0; i < tc.expected.size(); i++) {
                if (actual[i].getInteger() != tc.expected[i]) {
                    std::cout << "Testcase failed. Got: " << actual[i].getInteger()
                            << ", want: " << tc.expected[i] << "\n";
                    ++numBad;
                }
            }
            ++numGood;
        }
    }
    std::cout << (numBad == 0 ? "[OK]" : "[FAIL]") << " "
        << numGood << "/" << (numBad + numGood) << " testcases passed successfully.\n";
}
