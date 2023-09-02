#include <cassert>
#include <string>
#include <vector>
#include <unordered_set>
#include <unordered_map>

class Solution {
public:
  int dfs(int i) {
    if (ans.count(i) > 0) { // key i exists in ans:
      return ans[i];
    }

    ans[i] = 1 + dfs(i+1);
    for (int j = i; j < s.size(); j++) {
      std::string sub = s.substr(i, j+1-i);
      if (dict.count(sub) > 0) {
        ans[i] = std::min(ans[i], dfs(j+1));
      }
    }
    return ans[i];
  }

  int minExtraChar(std::string s, std::vector<std::string>& dictionary) {
    this->s = s;
    // Convert dictionary to a hash-set:
    dict.clear();
    for (const auto& word: dictionary) {
      dict.insert(word);
    }

    // Create a DP style ans hashmap
    ans.clear();
    ans[s.size()] = 0;

    return dfs(0);
  }
private:
  std::string s;
  std::unordered_set<std::string> dict;
  std::unordered_map<int, int> ans;
};

int main() {
  Solution s;

  std::vector<std::string> dict1 = {"leet", "code", "leetcode"};
  assert(s.minExtraChar("leetscode", dict1) == 1);

  std::vector<std::string> dict2 = {"hello", "world"};
  assert(s.minExtraChar("sayhelloworld", dict2) == 3);
};
