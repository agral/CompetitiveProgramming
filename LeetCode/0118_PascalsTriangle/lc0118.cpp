#include <cassert>
#include <vector>

class Solution {
public:
  std::vector<std::vector<int>> generate(int numRows) {
    std::vector<std::vector<int>> ans;
    ans.resize(numRows, {});
    ans[0] = {1};
    for (int i = 1; i < numRows; i++) {
      ans[i].resize(i+1); // n-th row has n entries when referring to 1-indexed rows.
                          // i-th 0-indexed row has i+1 entries.
      ans[i][0] = 1;
      ans[i][i] = 1;
      for (int k = 1; k < i; k++) {
        ans[i][k] = ans[i-1][k-1] + ans[i-1][k];
      }
    }
    return ans;
  }
};

int main() {
  Solution s;
  std::vector<std::vector<int>> expected_5 = {{1}, {1, 1}, {1, 2, 1}, {1, 3, 3, 1}, {1, 4, 6, 4, 1}};
  assert(s.generate(5) == expected_5);

  std::vector<std::vector<int>> expected_1 = {{1}};
  assert(s.generate(1) == expected_1);
}
