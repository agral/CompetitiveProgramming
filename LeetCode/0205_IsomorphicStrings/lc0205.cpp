#include <cassert>
#include <string>

class Solution {
public:
  bool isIsomorphic(std::string s, std::string t) {
    // s and t consist of any number of valid ascii characters.
    static constexpr int MAX_CHARS = 128;
    int last_index_of_char_s[MAX_CHARS] = {0};
    int last_index_of_char_t[MAX_CHARS] = {0};

    // Note: it is guaranteed that s.size() == t.size();
    // no need to check it separately.
    for (int i = 0; i < s.size(); ++i) {
      if (last_index_of_char_s[s[i]] != last_index_of_char_t[t[i]]) {
        return false;
      }
      // i+1 instead of i: value has to be different than zero (from zero-initialized array) for first char, i.e. s[0], t[0].
      last_index_of_char_s[s[i]] = i + 1;
      last_index_of_char_t[t[i]] = i + 1;
    }
    return true;
  }
};

int main() {
  Solution s;
  assert(s.isIsomorphic("egg", "add") == true);
  assert(s.isIsomorphic("foo", "bar") == false);
  assert(s.isIsomorphic("paper", "title") == true);
}
