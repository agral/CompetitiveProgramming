#pragma once

#include <variant>
#include <vector>

class NestedInteger {
public:
    NestedInteger(int num): m_contents(num), m_isPlainInteger(true) {}
    NestedInteger(std::vector<NestedInteger> list): m_contents(list), m_isPlainInteger(false) {}

    // Return true if this NestedInteger holds a single integer instance,
    // or false if it holds a NestedInteger instance.
    bool isInteger() const;

    // Return a single integer that this NestedInteger holds, if it holds a single integer.
    // Throw when called on instance that holds a nested list.
    int getInteger() const;

    // Return a nested list held by this NestedInteger instance.
    // Throw when called on an instance that holds a single integer.
    const std::vector<NestedInteger>& getList() const;
private:
    std::variant<int, std::vector<NestedInteger>> m_contents;
    bool m_isPlainInteger;
};
