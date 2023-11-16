#include <cassert>
#include <sstream>
#include <string>
#include <vector>

class Solution {
public:
  std::string findDifferentBinaryString(std::vector<std::string>& nums) {
    // Another approach found after more careful reading of problem description:
    // Input array is guaranteed to be of size N, each entry a string of length N.
    // (the part that there are at most N such strings is important).
    // A simpler solution: flip ith bit of ith answer; this is guaranteed
    // to generate a number of length N that is not yet in nums.
    std::stringstream ans;
    for (int idx = 0; idx < nums.size(); idx++) {
      ans << (nums[idx][idx] == '0' ? '1': '0');
    }

    return ans.str();
  }
};

int main() {
  Solution s;
  {
    std::vector<std::string> testcase = {"01", "10"};
    std::string ans = s.findDifferentBinaryString(testcase);
    assert(ans == "00" || ans == "11");
  }
  {
    std::vector<std::string> testcase = {"00", "01"};
    std::string ans = s.findDifferentBinaryString(testcase);
    assert(ans == "10" || ans == "11");
  }
  {
    std::vector<std::string> testcase = {"111", "011", "001"};
    std::string ans = s.findDifferentBinaryString(testcase);
    assert(ans == "000" || ans == "010" || ans == "100" || ans == "101" || ans == "110");
  }
}
