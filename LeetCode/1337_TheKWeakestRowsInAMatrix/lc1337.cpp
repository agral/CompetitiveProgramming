#include <algorithm>
#include <numeric>
#include <cassert>
#include <vector>
#include <utility>

#include <iostream>

class Solution {
public:
  std::vector<int> kWeakestRows(std::vector<std::vector<int>>& mat, int k) {
    std::vector<std::pair<int, int>> num_and_soldier_count(mat.size());
    for (int i = 0; i < mat.size(); i++) {
      int num_soldiers = std::accumulate(mat[i].begin(), mat[i].end(), 0);
      num_and_soldier_count[i] = std::make_pair(i, num_soldiers);
    }

    std::sort(num_and_soldier_count.begin(), num_and_soldier_count.end(),
              [](const std::pair<int, int>& lhs, const std::pair<int, int>& rhs) {
                  if (lhs.second == rhs.second) {
                    return lhs.first < rhs.first;
                  };
                  return lhs.second < rhs.second;
    });

    std::vector<int> ans(k);
    for (int i = 0; i < k; i++) {
      ans[i] = num_and_soldier_count[i].first;
    }
    return ans;
  }
};

int main() {
  Solution s;

  std::vector<std::vector<int>> example1 = {
    {1, 1, 0, 0, 0},
    {1, 1, 1, 1, 0},
    {1, 0, 0, 0, 0},
    {1, 1, 0, 0, 0},
    {1, 1, 1, 1, 1},
  };
  std::vector<int> expected_ans1 = {2, 0, 3};
  assert(s.kWeakestRows(example1, 3) == expected_ans1);

  std::vector<std::vector<int>> example2 = {
    {1, 0, 0, 0},
    {1, 1, 1, 1},
    {1, 0, 0, 0},
    {1, 0, 0, 0},
  };
  std::vector<int> expected_ans2 = {0, 2};
  assert(s.kWeakestRows(example2, 2) == expected_ans2);

  std::vector<std::vector<int>> own_example3 = {
    {1, 1, 1, 1, 1},
    {1, 0, 0, 0, 0},
    {1, 1, 0, 0, 0},
    {1, 1, 1, 1, 0},
    {1, 1, 1, 1, 1},
  };
  std::vector<int> expected_ans3 = {1, 2, 3};
  assert(s.kWeakestRows(own_example3, 3) == expected_ans3);

}
