#include <cassert>
#include <vector>

class Solution {
public:
  std::vector<int> findArray(std::vector<int>& pref) {
    std::vector<int> ans(pref.size(), 0);
    int cumulative_xor = 0;
    for (int i = 0; i < pref.size(); i++) {
      ans[i] = cumulative_xor ^ pref[i];
      cumulative_xor ^= ans[i];
    }
    return ans;
  };
};

int main() {
  Solution s;

  std::vector<int> example1 = {5, 2, 0, 3, 1};
  std::vector<int> expected_answer1 = {5, 7, 2, 3, 2};
  assert(s.findArray(example1) == expected_answer1);

  std::vector<int> example2 = {13};
  std::vector<int> expected_answer2 = {13};
  assert(s.findArray(example2) == expected_answer2);
}
