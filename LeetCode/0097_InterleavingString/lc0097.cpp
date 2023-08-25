#include <cassert>
#include <string>
#include <map>
#include <utility>

class Solution {
public:
  bool isInterleave(std::string s1, std::string s2, std::string s3) {
    // Check that component string lengths sum up to target string's length:
    if (s1.size() + s2.size() != s3.size()) {
      return false;
    }

    this->s1 = s1;
    this->s2 = s2;
    this->s3 = s3;
    memo.clear();
    return dfs(0, 0);
  }

  bool dfs(int f, int s) {
    if (f == s1.size() && s == s2.size()) {
      return true;
    }
    if (memo.count({f, s})) {
      return memo[{f, s}];
    }
    if ((f < s1.size()) && (s1[f] == s3[f + s]) && (dfs(f + 1, s))) {
      return true;
    }
    if ((s < s2.size()) && (s2[s] == s3[f + s]) && (dfs(f, s + 1))) {
      return true;
    }
    // else:
    memo[{f, s}] = false;
    return false;
  }

  std::map<std::pair<int, int>, bool> memo;
  std::string s1;
  std::string s2;
  std::string s3;
};

int main() {
  Solution s{};

  //std::string tc1_s1 = "aabcc", tc1_s2 = "dbbca", tc1_s3 = "aadbbcbcac";
  //assert(s.isInterleave(tc1_s1, tc1_s2, tc1_s3) == true);

  assert(s.isInterleave("aabcc", "dbbca", "aadbbcbcac") == true);
  assert(s.isInterleave("aabcc", "dbbca", "aadbbbaccc") == false);
  assert(s.isInterleave("", "", "") == true);
  assert(s.isInterleave("XXX", "XXX", "XXXXXX") == true);
}
