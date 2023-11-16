#include <bitset>
#include <cassert>
#include <string>
#include <vector>
#include <unordered_set>

class Solution {
public:
  std::string findDifferentBinaryString(std::vector<std::string>& nums) {
    const std::size_t string_size = nums[0].size();
    const int MAX = 1 << string_size; // max such generated number, not inclusively
    //std::vector<bool> is_present(MAX, false);
    // Note: 1 <= n <= 16, std::vector would require 65536 pre-allocated bools, each taking
    // up to a machine word of space, even for very sparse inputs - that's definitely too much.
    std::unordered_set<int> have_nums;
    for (const auto& numstr: nums) {
      //const std::size_t num{std::stoul(numstr, nullptr, 2)};
      have_nums.insert(std::stoi(numstr, nullptr, 2));
    }

    for (int k = 0; k < MAX; k++) {
      if (have_nums.find(k) == have_nums.end()) {
        std::bitset<16> bans(k);
        return bans.to_string().substr(16 - string_size);
      }
    }

    throw -1;
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
