#include <cassert>
#include <string>
#include <vector>

class Solution {
public:
  std::string reverseVowels(std::string s) {
    std::vector<char> vowels{};
    for (const auto& ch : s) {
      if (ch == 'a' || ch == 'A' || ch == 'e' || ch == 'E' || ch == 'i' || ch == 'I'
                    || ch == 'o' || ch == 'O' || ch == 'u' || ch == 'U') {
        vowels.push_back(ch);
      }
    }
    std::vector<char> ans;
    ans.resize(s.size());
    std::size_t vowel_pos = vowels.size() - 1;

    for (std::size_t i = 0; i < s.size(); i++) {
      if (s[i] == 'a' || s[i] == 'A' || s[i] == 'e' || s[i] == 'E' || s[i] == 'i' || s[i] == 'I'
                      || s[i] == 'o' || s[i] == 'O' || s[i] == 'u' || s[i] == 'U') {
        ans[i] = vowels[vowel_pos--];
      } else {
        ans[i] = s[i];
      }
    }

    return std::string{ans.begin(), ans.end()};
  }
};

int main() {
  Solution s{};
  assert(s.reverseVowels("hello") == "holle");
  assert(s.reverseVowels("leetcode") == "leotcede");
  assert(s.reverseVowels("babebibobu") == "bubobibeba");
}
