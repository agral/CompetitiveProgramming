#include <cassert>
#include <algorithm>
#include <string>
#include <vector>

class Solution {
public:
  std::string sortVowels(std::string s) {
    std::vector<char> ans{s.begin(), s.end()};
    std::vector<char> vowels;
    for (int i = 0; i < s.size(); i++) {
      if (s[i] == 'A' || s[i] == 'a' || s[i] == 'E' || s[i] == 'e' || s[i] == 'I' || s[i] == 'i'
          || s[i] == 'O' || s[i] == 'o' || s[i] == 'U' || s[i] == 'u') {
        vowels.push_back(s[i]);
      }
    }
    std::sort(vowels.begin(), vowels.end());
    int sorted_vowel_idx = 0;
    for (int i = 0; i < s.size(); i++) {
      if (s[i] == 'A' || s[i] == 'a' || s[i] == 'E' || s[i] == 'e' || s[i] == 'I' || s[i] == 'i'
          || s[i] == 'O' || s[i] == 'o' || s[i] == 'U' || s[i] == 'u') {
        ans[i] = vowels[sorted_vowel_idx++];
      }
    }

    return {ans.begin(), ans.end()};
  }
};

int main() {
  Solution s;
  assert(s.sortVowels("lEetcOde") == "lEOtcede");
  assert(s.sortVowels("lYmpH") == "lYmpH");
}
