#include "nested_integer.hpp"

bool NestedInteger::isInteger() const {
    return m_isPlainInteger;
}

int NestedInteger::getInteger() const {
    if (!m_isPlainInteger) {
        throw "Plain integer was requested, but I am a list";
    }
    return std::get<int>(m_contents);
}

const std::vector<NestedInteger>& NestedInteger::getList() const {
    if (m_isPlainInteger) {
        throw "List was requested, but I am a plain integer";
    }
    return std::get<std::vector<NestedInteger>>(m_contents);
}
