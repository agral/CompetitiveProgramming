#include <algorithm>
#include <cassert>
#include <string>
#include <vector>
#include <sstream>
#include <utility>

typedef std::pair<char, int> Count;

class Solution {
public:
  std::string frequencySort(std::string s) {
    std::vector<int> characters_count_in_s(128);
    for (const char c: s) {
      ++characters_count_in_s[c];
    }

    std::vector<Count> charCount;
    for (int i = 0; i < characters_count_in_s.size(); i++) {
      if (characters_count_in_s[i] > 0) {
        charCount.push_back(Count{static_cast<char>(i), characters_count_in_s[i]});
      }
    }

    std::ranges::sort(charCount, [](const Count& lhs, const Count& rhs) { return lhs.second > rhs.second; });

    std::stringstream ss_ans;
    for (const Count& pair: charCount) {
      ss_ans << std::string(pair.second, pair.first);
    }

    return ss_ans.str();
  }
};

int main() {
  Solution s;
  {
    std::string fsorted = s.frequencySort("tree");
    assert(fsorted == "eert" || fsorted == "eetr");
  }
  {
    std::string fsorted = s.frequencySort("cccaaa");
    assert(fsorted == "aaaccc" || fsorted == "cccaaa");
  }
  {
    std::string fsorted = s.frequencySort("Aabb");
    assert(fsorted == "bbAa" || fsorted == "bbaA");
  }
}
