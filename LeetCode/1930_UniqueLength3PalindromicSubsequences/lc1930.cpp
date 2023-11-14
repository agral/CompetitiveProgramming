#include <cassert>
#include <string>
#include <vector>
#include <unordered_set>

class Solution {
public:
  const int SIZE = 'z' - 'a' + 1;
  int countPalindromicSubsequence(std::string s) {
    int ans = 0;
    std::vector<int> first_occurrence(SIZE, s.size());
    std::vector<int> last_occurrence(SIZE, 0);
    for (int i = 0; i < s.size(); i++) {
      const int offset = s[i] - 'a';
      if (i < first_occurrence[offset]) {
        first_occurrence[offset] = i;
      }
      last_occurrence[offset] = i;
    }

    for (int i = 0; i < SIZE; i++) {
      if (first_occurrence[i] < last_occurrence[i]) {
        std::unordered_set<int> uniques{s.begin() + 1 + first_occurrence[i], s.begin() + last_occurrence[i]};
        ans += uniques.size();
      }
    }
    return ans;
  }
};

int main() {
  Solution s;
  assert(s.countPalindromicSubsequence("aabca") == 3); // "aaa", "aba", "aca"
  assert(s.countPalindromicSubsequence("abc") == 0);
  assert(s.countPalindromicSubsequence("bbcbaba") == 4); // "bbb", "bcb", "bab", "aba"
}
