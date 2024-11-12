#include <algorithm>
#include <cassert>
#include <vector>

class Solution {
public:
    std::vector<int> maximumBeauty(std::vector<std::vector<int>>& items, std::vector<int>& queries) {
        std::vector<int> ans;
        ans.reserve(queries.size());
        std::vector<int> prices;
        std::vector<int> beauty(items.size() + 1, 0);

        std::ranges::sort(items);

        // Store prices in separate vector. This is necessary to run upper_bound() later.
        for (const std::vector<int>& item: items) {
            prices.push_back(item[0]);
        }

        for (int i = 0; i < items.size(); i++) {
            beauty[i+1] = std::max(beauty[i], items[i][1]);
        }

        for (const int query: queries) {
            const int idx = std::ranges::upper_bound(prices, query) - prices.begin();
            ans.push_back(beauty[idx]);
        }
        return ans;
    }
};

int main() {
    Solution s;
    {
        std::vector<std::vector<int>> items {{1,2}, {3,2}, {2,4}, {5,6}, {3,5}};
        std::vector<int> queries {1, 2, 3, 4, 5, 6};
        std::vector<int> expected {2, 4, 5, 5, 6, 6};
        assert(s.maximumBeauty(items, queries) == expected);
    }
    {
        std::vector<std::vector<int>> items {{1,2}, {1,2}, {1,3}, {1,4}};
        std::vector<int> queries {1};
        std::vector<int> expected {4};
        assert(s.maximumBeauty(items, queries) == expected);
    }
    {
        std::vector<std::vector<int>> items {{10, 1000}};
        std::vector<int> queries {5};
        std::vector<int> expected {0};
        assert(s.maximumBeauty(items, queries) == expected);
    }
}
