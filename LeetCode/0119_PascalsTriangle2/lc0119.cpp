#include <cassert>
#include <vector>

class Solution {
public:
  std::vector<int> getRow(int rowIndex) {
    std::vector<std::vector<int>> ans;
    ans.resize(rowIndex + 1, {});
    ans[0] = {1};
    for (int i = 1; i <= rowIndex; i++) {
      ans[i].resize(i + 1);
      ans[i][0] = 1;
      for (int k = 1; k < i; k++) {
        ans[i][k] = ans[i-1][k-1] + ans[i-1][k];
      }
      ans[i][i] = 1;
    }
    return ans[rowIndex];
  }
};

int main() {
  Solution s;
  assert(s.getRow(3) == std::vector<int>({1, 3, 3, 1}));
  assert(s.getRow(0) == std::vector<int>({1}));
  assert(s.getRow(1) == std::vector<int>({1, 1}));
}
