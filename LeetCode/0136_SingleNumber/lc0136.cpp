#include <cassert>
#include <vector>

class Solution {
public:
  // nums is an array where every element appears twice, except for one element
  // which appears just once. The question asks to find that element.
  // The trick is to xor the numbers together. Since N xor 0 is N, and N xor N is 0,
  // the result: A xor A xor B xor B xor C xor D xor D xor ... xor Z xor Z is C.
  // Note: xor is commutative as addition/multiplication, so no need to sort the array etc.
  int singleNumber(std::vector<int>& nums) {
    int ans = 0;
    for (const auto& num: nums) {
      ans ^= num;
    }
    return ans;
  }
};

int main() {
  Solution s;
  std::vector<int> testcase1 = {2, 2, 1};
  assert(s.singleNumber(testcase1) == 1);

  std::vector<int> testcase2 = {4, 1, 2, 1, 2};
  assert(s.singleNumber(testcase2) == 4);

  std::vector<int> testcase3 = {1};
  assert(s.singleNumber(testcase3) == 1);
}
