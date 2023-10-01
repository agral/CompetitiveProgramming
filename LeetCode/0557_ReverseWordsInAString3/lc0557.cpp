#include <algorithm>
#include <string>
#include <cassert>
#include <sstream>
#include <iostream>

class Solution {
public:
  std::string reverseWords(std::string s) {
    std::stringstream ans;
    std::istringstream iss(s);
    bool is_first_word = true;
    while (iss) {
      std::string word;
      iss >> word;
      if (!word.empty()) {
        std::reverse(word.begin(), word.end());
        if (is_first_word) {
          is_first_word = false;
        } else {
          ans << ' ';
        }
        ans << word;
      }
    }
    return ans.str();
  }
};

int main() {
  Solution s;
  std::string example1 = "Let's take LeetCode contest";
  assert(s.reverseWords(example1) == "s'teL ekat edoCteeL tsetnoc");

  std::string example2 = "God Ding";
  assert(s.reverseWords(example2) == "doG gniD");
}
