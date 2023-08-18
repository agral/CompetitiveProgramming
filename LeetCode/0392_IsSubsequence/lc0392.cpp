#include <cassert>
#include <string>

class Solution {
public:
  bool isSubsequence(std::string s, std::string t) {
    if (s.empty()) {
      return true;
    }

    auto it_s = s.cbegin();
    for (auto it_t = t.cbegin(); it_t != t.cend(); it_t++) {
      if ((*it_t) == (*it_s)) {
        it_s++;
        if (it_s == s.cend()) {
          return true;
        }
      }
    }
    return false;
  }
};

int main() {
  Solution s{};
  assert(s.isSubsequence("abc", "ahbgdc") == true);
  assert(s.isSubsequence("axc", "ahbgdc") == false);
}
