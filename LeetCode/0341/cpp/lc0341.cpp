#include <cassert>
#include <stack>
#include <vector>

#include "nested_integer.hpp"

// Runtime complexity: O(n)
// Aux space complexity: O(n)
class NestedIterator {
public:
    NestedIterator(std::vector<NestedInteger>& nestedList) {
        addVectorNi(nestedList);
    };

    int next() {
        int ans = m_stack.top().getInteger();
        m_stack.pop();
        return ans;
    };

    bool hasNext() {
        // If stack is not empty and contains a list (so not isInteger()), process the list:
        while (!m_stack.empty() && !m_stack.top().isInteger()) {
            std::vector<NestedInteger> list = m_stack.top().getList();
            m_stack.pop();
            addVectorNi(list);
        }
        return !m_stack.empty();
    }

private:
    std::stack<NestedInteger> m_stack;
    void addVectorNi(const std::vector<NestedInteger>& vni) {
        for (int i = vni.size()-1; i >= 0; i--) {
            m_stack.push(vni[i]);
        }
    }
};

int main() {
}
