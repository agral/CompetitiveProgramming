#include <iostream>
#include <sstream>
#include <vector>

class UnionFind {
public:
    UnionFind(int n) {
        m_id.resize(n);
        for (int i = 0; i < n; i++) {
            m_id[i] = i;
        }
    }

    int find(int node) {
        if (m_id[node] == node) {
            return node;
        } else {
            m_id[node] = find(m_id[node]);
            return m_id[node];
        }
    }

    void join(int lhs, int rhs) {
        int l = find(lhs);
        int r = find(rhs);
        if (l < r) {
            m_id[r] = l;
        } else {
            m_id[l] = r;
        }
    }
private:
    std::vector<int> m_id;
};

// Runtime complexity: O(s1.size() + baseStr.size())
// Auxiliary space complexity: O(26) + O(baseStr.size())
class Solution {
public:
    std::string smallestEquivalentString(std::string s1, std::string s2, std::string baseStr) {
        const int NUM_LETTERS = 'z' - 'a' + 1;
        UnionFind uf(NUM_LETTERS);
        std::stringstream ans;
        for (int i = 0; i < s1.size(); i++) {
            uf.join(s1[i] - 'a', s2[i] - 'a');
        }

        for (const char c: baseStr) {
            ans << static_cast<char>(uf.find(c - 'a') + 'a');
        }
        return ans.str();
    }
};

int main() {
    struct Testcase {
        std::string s1;
        std::string s2;
        std::string baseStr;
        std::string expected;
    };
    Testcase testcases[] = {
        {"parker", "morris", "parser", "makkek"},
        {"hello", "world", "hold", "hdld"},
        {"leetcode", "programs", "sourcecode", "aauaaaaada"},
    };
    Solution s{};
    int numGood = 0, numBad = 0;
    for (Testcase& tc: testcases) {
        auto actual = s.smallestEquivalentString(tc.s1, tc.s2, tc.baseStr);
        if (actual != tc.expected) {
            std::cout << "Testcase #" << numGood+numBad+1 << " failed. Got: " << actual
                << ", want: " << tc.expected << "\n";
            ++numBad;
        } else {
            ++numGood;
        }
    }
    std::cout << (numBad == 0 ? "[OK]" : "[FAIL]") << " "
        << numGood << "/" << (numBad + numGood) << " testcases passed successfully.\n";
}
