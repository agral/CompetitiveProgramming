#include <iostream>
#include <queue>
#include <vector>
#include <sstream>
#include <algorithm>

struct GPV { // Grid Poisition & Value
    int h;
    int w;
    int value; // taken at grid[h][w]
};

struct QueryIndex {
    int query;
    int index;
};

class Solution {
public:
    std::vector<int> maxPoints(std::vector<std::vector<int>>& grid, std::vector<int>& queries) {
        const int H = grid.size();
        const int W = grid[0].size();
        const int DIRS[4][2] = {{1, 0}, {0, -1}, {-1, 0}, {0, 1}};
        std::vector<int> ans(queries.size());

        std::vector<std::vector<bool>> seen(H, std::vector<bool>(W));
        auto cmp = [](const GPV& lhs, const GPV& rhs) { return lhs.value > rhs.value; };
        std::priority_queue<GPV, std::vector<GPV>, decltype(cmp)> minHeap(cmp);
        int acc = 0;
        minHeap.emplace(0, 0, grid[0][0]);
        seen[0][0] = true;

        for (const auto& [query, index]: getQueryIndexes(queries)) {
            while (!minHeap.empty()) {
                const auto [h, w, value] = minHeap.top();
                std::cout << "h=" << h << ", w=" << w << ", value=" << value << "\n";
                minHeap.pop();
                if (value >= query) {
                    // Return (h, w, value) back to minHeap, as the smallest neighbor is still larger than query.
                    minHeap.emplace(h, w, value);
                    break; // directly go-to ans setting, by breaking the while-loop. Important.
                }
                acc += 1;
                for (const auto& [dh, dw]: DIRS) {
                    int hh = h + dh;
                    int ww = w + dw;
                    if (hh < 0 || hh == H || ww < 0 || ww == W || seen[hh][ww]) {
                        continue;
                    }
                    minHeap.emplace(hh, ww, grid[hh][ww]);
                    seen[hh][ww] = true;
                }
            }
            ans[index] = acc;
        }

        return ans;
    }
private:
    std::vector<QueryIndex> getQueryIndexes(std::vector<int>& queries) {
        std::vector<QueryIndex> ans;
        for (int i = 0; i < queries.size(); ++i) {
            ans.push_back({queries[i], i});
        }
        std::ranges::sort(ans, std::ranges::less{}, [](const QueryIndex& qi) { return qi.query; });
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
        std::vector<std::vector<int>> grid;
        std::vector<int> queries;
        std::vector<int> expected;
    };
    Testcase testcases[] = {
        {/*grid=*/{{1, 2, 3}, {2, 5, 7}, {3, 5, 1}}, /*queries=*/{5, 6, 2}, /*expected=*/{5, 8, 1}},
        {/*grid=*/{{5, 2, 1}, {1, 1, 2}}, /*queries=*/{3}, /*expected=*/{0}},
    };
    Solution s{};
    int numGood = 0, numBad = 0;
    for (Testcase& tc: testcases) {
        auto actual = s.maxPoints(tc.grid, tc.queries);
        if (actual != tc.expected) {
            std::cout << "Testcase " << vectorToString(tc.queries)
                << " failed. Got: " << vectorToString(actual)
                << ", want: " << vectorToString(tc.expected) << "\n";
            ++numBad;
        } else {
            ++numGood;
        }
    }
    std::cout << (numBad == 0 ? "[OK]" : "[FAIL]") << " "
        << numGood << "/" << (numBad + numGood) << " testcases passed successfully.\n";
}
