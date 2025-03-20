#include <iostream>
#include <vector>
#include <sstream>

class UnionFind {
public:
    // the minimum number that has all set bits in its binary representation,
    // greater than 10^5 is 2^17-1 (equal to 131071).
    static constexpr int MAX_WEIGHT = (1 << 17) - 1;

    UnionFind(int n) : m_id(n), m_weight(n, MAX_WEIGHT), m_rank(n) {
        for (int i = 0; i < n; i++) {
            m_id[i] = i;
        }
    }

    int minCost(int u, int v) {
        if (u == v) {
            return 0;
        }
        int start = find(u);
        int end = find(v);
        if (start == end) {
            return m_weight[start];
        }
        return -1;
    }

    void doUnionByRank(int u, int v, int w) {
        int start = find(u);
        int end = find(v);
        int combinedWeight = m_weight[start] & m_weight[end] & w;
        m_weight[start] = combinedWeight;
        m_weight[end] = combinedWeight;

        if (start == end) {
            return;
        }

        if (m_rank[start] < m_rank[end]) {
            m_id[start] = end;
        }
        else if (m_rank[end] < m_rank[start]) {
            m_id[end] = start;
        }
        else {
            m_id[start] = end;
            m_rank[end]++;
        }
    }

private:
    /*
    int find(int vertex_id) {
        if (m_id[vertex_id] != vertex_id) {
            m_id[vertex_id] = find(m_id[vertex_id]);
        }
        return m_id[vertex_id];
    }
    */
    int find(int vid) {
        return m_id[vid] == vid ? vid : m_id[vid] = find(m_id[vid]);
    }

    std::vector<int> m_id;
    std::vector<int> m_weight;
    std::vector<int> m_rank;
};

// Runtime complexity: O(n*logn)
// Auxiliary space complexity: O(n)
class Solution {
public:
    std::vector<int> minimumCost(int n, std::vector<std::vector<int>>& edges, std::vector<std::vector<int>>& query) {
        UnionFind uf{n};
        for (const std::vector<int>& edge: edges) {
            uf.doUnionByRank(edge[0], edge[1], edge[2]);
        }

        std::vector<int> ans = {};
        for (const std::vector<int>& q: query) {
            ans.push_back(uf.minCost(q[0], q[1]));
        }
        return ans;
    }
};

template<typename T>
std::string vectorToString(std::vector<T> vec) {
    std::string separator = "";
    std::stringstream ss;
    ss << "[";
    for (const auto& elem: vec) {
        ss << separator << elem;
        separator = ", ";
    }
    ss << "]";
    return ss.str();
}

int main() {
    struct Testcase {
        int n;
        std::vector<std::vector<int>> edges;
        std::vector<std::vector<int>> query;
        std::vector<int> expected;
    };
    Testcase testcases[] = {
        {5, {{0, 1, 7}, {1, 3, 7}, {1, 2, 1}}, {{0, 3}, {3, 4}}, {1, -1}},
        {3, {{0, 2, 7}, {0, 1, 15}, {1, 2, 6}, {1, 2, 1}}, {{1, 2}}, {0}}
    };
    Solution s{};
    int numGood = 0, numBad = 0;
    for (Testcase& tc: testcases) {
        auto actual = s.minimumCost(tc.n, tc.edges, tc.query);
        if (actual != tc.expected) {
            std::cout << "Testcase failed. Got: " << vectorToString(actual)
                << ", want: " << vectorToString(tc.expected) << "\n";
            ++numBad;
        } else {
            ++numGood;
        }
    }
    std::cout << (numBad == 0 ? "[OK]" : "[FAIL]") << " "
        << numGood << "/" << (numBad + numGood) << " testcases passed successfully.\n";
}
