#include <iostream>
#include <unordered_set>
#include <vector>

class UnionFind {
public:
    UnionFind(int n) : m_id(n), m_numNodes(n, 1), m_numEdges(n), m_rank(n) {
        for (int i = 0; i < n; i++) {
            m_id[i] = i;
        }
    }

    void doUnionByRank(int u, int v) {
        int start = find(u);
        int end = find(v);
        ++m_numEdges[start];

        if (start == end) {
            return;
        }

        if (m_rank[start] <= m_rank[end]) {
            m_id[start] = end;
            m_numNodes[end] += m_numNodes[start];
            m_numEdges[end] += m_numEdges[start];
            if (m_rank[start] == m_rank[end]) {
                m_rank[end] += 1;
            }
        }
        else /*if (m_rank[end] < m_rank[start])*/ {
            m_id[end] = start;
            m_numNodes[start] += m_numNodes[end];
            m_numEdges[start] += m_numEdges[end];
        }
    }

    int find(int vid) {
        return m_id[vid] == vid ? vid : m_id[vid] = find(m_id[vid]);
    }

    bool isFormingCompleteGraph(int vid) {
        return m_numEdges[vid] == m_numNodes[vid] * (m_numNodes[vid] - 1) / 2;
    }

private:
    std::vector<int> m_id;
    std::vector<int> m_numNodes;
    std::vector<int> m_numEdges;
    std::vector<int> m_rank;
};

// Runtime complexity: O(|V| + |E|)
// Auxiliary space complexity: O(|V| + |E|)
class Solution {
public:
    int countCompleteComponents(int n, std::vector<std::vector<int>>& edges) {
        UnionFind uf{n};
        for (const std::vector<int>& edge: edges) {
            uf.doUnionByRank(edge[0], edge[1]);
        }

        std::unordered_set<int> parents;
        int ans = 0;
        for (int i = 0; i < n; i++) {
            int parent = uf.find(i);
            if (parents.insert(parent).second) {
                if (uf.isFormingCompleteGraph(parent)) {
                    ans += 1;
                }
            }
        }
        return ans;
    }
};

int main() {
    struct Testcase {
        int n;
        std::vector<std::vector<int>> edges;
        int expected;
    };
    Testcase testcases[] = {
        {6, {{0, 1}, {0, 2}, {1, 2}, {3, 4}}, 3},
        {6, {{0, 1}, {0, 2}, {1, 2}, {3, 4}, {3, 5}}, 1},
    };
    Solution s{};
    int numGood = 0, numBad = 0;
    for (Testcase& tc: testcases) {
        auto actual = s.countCompleteComponents(tc.n, tc.edges);
        if (actual != tc.expected) {
            std::cout << "Testcase failed. Got: " << actual
                << ", want: " << tc.expected << "\n";
            ++numBad;
        } else {
            ++numGood;
        }
    }
    std::cout << (numBad == 0 ? "[OK]" : "[FAIL]") << " "
        << numGood << "/" << (numBad + numGood) << " testcases passed successfully.\n";
}
