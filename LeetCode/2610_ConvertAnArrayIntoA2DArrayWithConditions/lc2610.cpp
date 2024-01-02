#include <cassert>
#include <vector>
#include <unordered_set>

class Solution {
public:
  std::vector<std::vector<int>> findMatrix(std::vector<int>& nums) {
    std::vector<std::unordered_set<int>> partition = {{}};
    for (const int num: nums) {
      bool isAdded = false;
      for (int row = 0; !isAdded && row < partition.size(); ++row) {
        auto it = partition[row].find(num);
        if (it == partition[row].end()) {
          partition[row].insert(num);
          isAdded = true;
        }
      }
      if (!isAdded) {
        partition.push_back({num});
      }
    }

    // Now convert vector-of-sets to a vector-of-vectors and return it:
    std::vector<std::vector<int>> ans(partition.size(), std::vector<int>{});
    for (int row = 0; row < partition.size(); ++row) {
      std::copy(partition[row].begin(), partition[row].end(), std::back_inserter(ans[row]));
    }
    return ans;
  }
};

// Note: for easy testing, we need to sort the answers' rows.
// Then we don't have to take into account the order unordered_set gets converted
// inta a vector. This requires std::sort from <algorithm>.
#include <algorithm>
int main() {
  Solution s;
  {
    std::vector<int> nums = {1, 3, 4, 1, 2, 3, 1};
    std::vector<std::vector<int>> expected_answer = {{1, 2, 3, 4}, {1, 3}, {1}};
    std::vector<std::vector<int>> actual_answer = s.findMatrix(nums);
    for (int i = 0; i < actual_answer.size(); ++i) {
      std::sort(actual_answer[i].begin(), actual_answer[i].end());
    }
    assert(actual_answer == expected_answer);
  }
  {
    std::vector<int> nums = {1, 2, 3, 4};
    std::vector<std::vector<int>> expected_answer = {{1, 2, 3, 4}};
    std::vector<std::vector<int>> actual_answer = s.findMatrix(nums);
    for (int i = 0; i < actual_answer.size(); ++i) {
      std::sort(actual_answer[i].begin(), actual_answer[i].end());
    }
    assert(actual_answer == expected_answer);
  }
}
