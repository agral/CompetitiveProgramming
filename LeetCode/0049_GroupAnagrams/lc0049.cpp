#include <algorithm>
#include <cassert>
#include <string>
#include <unordered_map>
#include <vector>

class Solution {
public:
  std::vector<std::vector<std::string>> groupAnagrams(std::vector<std::string>& strs) {
    std::unordered_map<std::string, std::vector<std::string>> sortedToWords{};
    for (const std::string& s: strs) {
      std::string sorted{s};
      std::sort(sorted.begin(), sorted.end());
      sortedToWords[sorted].push_back(s);
    }

    std::vector<std::vector<std::string>> ans;
    for (const auto& [_, anagramGroup]: sortedToWords) {
      ans.push_back(anagramGroup);
    }
    return ans;
  }
};

int main() {
  Solution s;
  {
    std::vector<std::string> input = {"eat", "tea", "tan", "ate", "nat", "bat"};
    std::vector<std::vector<std::string>> ans = s.groupAnagrams(input);
    assert(ans.size() == 3); // sloppy testing, but string arrays are hard to test anyway.
  }
  {
    std::vector<std::string> input = {""};
    std::vector<std::vector<std::string>> ans = s.groupAnagrams(input);
    assert(ans == std::vector<std::vector<std::string>>{{""}});
  }
  {
    std::vector<std::string> input = {"a"};
    std::vector<std::vector<std::string>> ans = s.groupAnagrams(input);
    assert(ans == std::vector<std::vector<std::string>>{{"a"}});
  }
}
