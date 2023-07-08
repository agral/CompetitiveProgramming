#include <cassert>
#include <string>
#include <sstream>

class Solution {
public:
  std::string mergeAlternately(std::string word1, std::string word2) {
    auto limit = std::min(word1.size(), word2.size());
    std::stringstream out{};
    for (std::size_t k = 0; k < limit; k++) {
      out << word1[k] << word2[k];
    }
    std::string* ptr = word1.size() > word2.size() ? &word1 : &word2;
    for (std::size_t k = limit; k < ptr->size(); k++) {
      out << (*ptr)[k];
    }
    return out.str();
  }
};

int main() {
  auto s = Solution{};

  assert(s.mergeAlternately("abc", "pqr") == "apbqcr");
  assert(s.mergeAlternately("ab", "pqrs") == "apbqrs");
  assert(s.mergeAlternately("abcd", "pq") == "apbqcd");
  assert(s.mergeAlternately("abcdefgh", "") == "abcdefgh");
  assert(s.mergeAlternately("", "abcdefgh") == "abcdefgh");
}
