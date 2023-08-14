#include <cassert>
#include <string>
#include <vector>

class Solution {
public:
  size_t gcd(size_t a, size_t b) {
    while (b > 0) {
      size_t mod = a % b;
      a = b;
      b = mod;
    }
    return a;
  }

  std::string gcdOfStrings(std::string str1, std::string str2) {
    size_t limit = gcd(str1.size(), str2.size());
    for (size_t i = gcd(str1.size(), str2.size()); i > 0; i--) {
      if ((str1.size() % i == 0) && (str2.size() % i == 0)) {
        std::string divisor_candidate = str1.substr(0, i);
        bool isDivisible = true; // so far, but this can change any moment! :-)
        for (int i = 0; isDivisible && i < str1.size(); i += divisor_candidate.size()) {
          if (str1.substr(i, divisor_candidate.size()) != divisor_candidate) {
            isDivisible = false;
          }
        }
        for (int j = 0; isDivisible && j < str2.size(); j += divisor_candidate.size()) {
          if (str2.substr(j, divisor_candidate.size()) != divisor_candidate) {
            isDivisible = false;
          }
        }
        if (isDivisible) {
          return divisor_candidate;
        }
      }
    }
    return {};
  }
};

int main() {
  Solution s{};
  
  assert(s.gcdOfStrings("ABCABC", "ABC") == "ABC");
  assert(s.gcdOfStrings("ABABAB", "ABAB") == "AB");
  assert(s.gcdOfStrings("LEET", "CODE") == "");
  assert(s.gcdOfStrings("A", "A") == "A");
}
