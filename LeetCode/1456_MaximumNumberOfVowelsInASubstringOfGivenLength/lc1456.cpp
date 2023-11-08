#include <cassert>
#include <string>

class Solution {
public:
  int maxVowels(std::string s, int k) {
    int score = 0;
    int rs = 0; // Range start, inclusive
    int re = 0; // Range end, not inclusive

    // Calculate the vowel content in the initial range:
    while (re < k) {
      if (s[re] == 'a' || s[re] == 'e' || s[re] == 'i' || s[re] == 'o' || s[re] == 'u') {
        score++;
      }
      re++;
    }
    int ans = score;

    // Then shift the window by moving range-start and range-end pointers until range-end
    // grows to string's size. Increase the score when range's end contains a vowel, and decrease
    // it when range-start moves past a vowel.
    while (re < s.size()) {
      if (s[re] == 'a' || s[re] == 'e' || s[re] == 'i' || s[re] == 'o' || s[re] == 'u') {
        score++;
      }
      re++;
      if (s[rs] == 'a' || s[rs] == 'e' || s[rs] == 'i' || s[rs] == 'o' || s[rs] == 'u') {
        score--;
      }
      rs++;
      if (score > ans) { ans = score; }
    }
    return ans;
  }
};

int main() {
  Solution s;
  assert(s.maxVowels("abciiidef", 3) == 3); // "iii"
  assert(s.maxVowels("aoeiu", 2) == 2); // "ao"
  assert(s.maxVowels("leetcode", 3) == 2); // "ee"
  assert(s.maxVowels("", 0) == 0); // ""
}
