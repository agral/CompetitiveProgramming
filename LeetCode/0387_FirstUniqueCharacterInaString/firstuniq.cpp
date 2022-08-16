// Solution to LeetCode problem #0387, "First Unique Character in a String"
// https://leetcode.com/problems/first-unique-character-in-a-string/
// by https://gralin.ski/

#include <iostream>
#include <string>
#include <map>

class Histogram {
 public:
  explicit Histogram(const std::string& str) {
    for (char c = 'a'; c < 'z'; ++c) {
      m_freq[c] = 0;
    }
    for (const auto& ch: str) {
      m_freq[ch] += 1;
    }
  }

 std::map<char, int> m_freq;
};

class Solution {
 public:
  int firstUniqChar(std::string s) {
    auto histogram = Histogram{s};
    for (std::size_t k = 0; k < s.size(); ++k) {
      if (histogram.m_freq[s[k]] == 1) {
        return k;
      }
    }
    return -1;
  }
};

int main() {
  Solution s;

  {
    std::cout << "Example 1: leetcode\n";
    if (0 == s.firstUniqChar("leetcode")) {
      std::cout << "  Passed\n";
    } else {
      std::cout << "Failed\n";
    }
  }
  {
    std::cout << "Example 2: loveleetcode\n";
    if (2 == s.firstUniqChar("loveleetcode")) {
      std::cout << "  Passed\n";
    } else {
      std::cout << "Failed\n";
    }
  }
  {
    std::cout << "Example 3: aabb\n";
    if (-1 == s.firstUniqChar("aabb")) {
      std::cout << "  Passed\n";
    } else {
      std::cout << "Failed\n";
    }
  }
}
