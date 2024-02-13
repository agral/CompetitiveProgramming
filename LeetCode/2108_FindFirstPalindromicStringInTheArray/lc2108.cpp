#include <cassert>
#include <string>
#include <vector>

class Solution {
public:
  std::string firstPalindrome(std::vector<std::string>& words) {
    for (std::string& word: words) {
      if (isPalindrome(word)) {
        return word;
      }
    }
    return "";
  }

private:
  bool isPalindrome(std::string& word) {
    const int SIZE = word.size();
    for (int i = 0; i <= SIZE / 2; ++i) {
      if (word[i] != word[SIZE - 1 - i]) {
        return false;
      }
    }
    return true;
  }
};

int main() {
  Solution s;
  {
    std::vector<std::string> words = {"abc", "car", "ada", "racecar", "cool"};
    assert(s.firstPalindrome(words) == "ada");
  }
  {
    std::vector<std::string> words = {"notapalindrome", "racecar"};
    assert(s.firstPalindrome(words) == "racecar");
  }
  {
    std::vector<std::string> words = {"def", "ghi"};
    assert(s.firstPalindrome(words) == "");
  }
  {
    std::vector<std::string> words = {"abc", "def", "ghi", "a", "racecar"};
    assert(s.firstPalindrome(words) == "a");
  }
  {
    std::vector<std::string> words = {"abc", "def", "aa"};
    assert(s.firstPalindrome(words) == "aa");
  }
}
