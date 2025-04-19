#include <algorithm>
#include <iostream>
#include <set>
#include <unordered_map>
#include <vector>
#include <sstream>

// Runtime complexity:O(XlogX), where X is the size (length) of the constructed graph
// Auxiliary space complexity: O(X)
class Solution {
public:
    std::vector<std::string> findItinerary(std::vector<std::vector<std::string>>& tickets) {
        std::vector<std::string> ans{};
        std::unordered_map<std::string, std::multiset<std::string>> graph{};

        for (const std::vector<std::string>& ticket: tickets) {
            graph[ticket[0]].insert(ticket[1]);
        }
        // Start from the JFK and construct the answer recursively checking all the available connections.
        dfs(graph, "JFK", ans);
        std::reverse(ans.begin(), ans.end());
        return ans;
    }
private:
    void dfs(std::unordered_map<std::string, std::multiset<std::string>>& graph,
             const std::string& node, std::vector<std::string>& ans) {
        while (graph.contains(node) && !graph[node].empty()) {
            const std::string new_node = *graph[node].begin();
            graph[node].erase(graph[node].begin());
            dfs(graph, new_node, ans);
        }
        ans.push_back(node);
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
        std::vector<std::vector<std::string>> tickets;
        std::vector<std::string> expected;
    };
    Testcase testcases[] = {
        {
            /*tickets=*/{{"MUC", "LHR"}, {"JFK", "MUC"}, {"SFO", "SJC"}, {"LHR", "SFO"}},
            /*expected=*/{"JFK", "MUC", "LHR", "SFO", "SJC"},
        },
        {
            /*tickets=*/{{"JFK", "SFO"}, {"JFK", "ATL"}, {"SFO", "ATL"}, {"ATL", "JFK"}, {"ATL", "SFO"}},
            /*expected=*/{"JFK", "ATL", "JFK", "SFO", "ATL", "SFO"},
        },
    };
    Solution s{};
    int numGood = 0, numBad = 0;
    for (Testcase& tc: testcases) {
        auto actual = s.findItinerary(tc.tickets);
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
