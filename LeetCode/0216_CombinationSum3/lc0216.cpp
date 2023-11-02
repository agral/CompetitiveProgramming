#include <cassert>
#include <vector>

class Solution {
public:
  std::vector<std::vector<int>> combinationSum3(int k, int n) {
    m_ans.clear();
    std::vector<int> accumulator;
    backtrack(accumulator, n, 0, 1, k);
    return m_ans;
  }

  void backtrack(std::vector<int>& contents, int target_value, int sum, int index, int k) {
    if ((contents.size() == k) && (sum == target_value)) {
      m_ans.push_back(contents);
    }
    else if (sum < target_value && contents.size() < k) {
      for (int i = index; i <= 9; i++) {
        contents.push_back(i);
        backtrack(contents, target_value, sum + i, i + 1, k);
        contents.pop_back();
      }
    }
  }
private:
  std::vector<std::vector<int>> m_ans;
};

int main() {
  Solution s;
  std::vector<std::vector<int>> expected_ans_3_7 = {{1, 2, 4}};
  assert(s.combinationSum3(3, 7) == expected_ans_3_7);
  std::vector<std::vector<int>> expected_ans_3_9 = {{1, 2, 6}, {1, 3, 5}, {2, 3, 4}};
  assert(s.combinationSum3(3, 9) == expected_ans_3_9);
  std::vector<std::vector<int>> expected_ans_4_1 = {};
  assert(s.combinationSum3(4, 1) == expected_ans_4_1);
}
