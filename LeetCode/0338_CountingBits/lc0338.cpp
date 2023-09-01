#include <cassert>
#include <vector>

class Solution {
public:
  std::vector<int> countBits(int n) {
    std::vector<int> ans;
    ans.resize(n + 1, 0);
    for (int i = 1; i < n+1; i++) {
      ans[i] = ans[i & (i - 1)] + 1; 
    }
    return ans;
  };
};

int main() {
  Solution s{};

  std::vector<int> expected_output_2 = {0, 1, 1};
  assert(s.countBits(2) == expected_output_2);
  std::vector<int> expected_output_5 = {0, 1, 1, 2, 1, 2};
  assert(s.countBits(5) == expected_output_5);
};
