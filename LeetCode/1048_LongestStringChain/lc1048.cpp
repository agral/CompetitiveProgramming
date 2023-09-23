#include <cassert>
#include <string>
#include <vector>
#include <unordered_map>

class Solution {
public:
  int getChainLength(std::string& word) {
    // Return 0 for words not in the list.
    if (mwords.find(word) == mwords.end()) {
      return 0;
    }

    if (word.size() == 1) {
      return 1;
    }

    // If chainlength of this word is already known, return it:
    if (auto it_clen = clen.find(word); it_clen != clen.end()) {
      return it_clen->second;
    }
    // otherwise calculate it as 1 + max(all sub-chains), and store the result:
    int max_subl = 0;
    for (int i = 0; i < word.size(); i++) {
      std::string subword = word;
      subword.erase(i, 1);
      int subl = getChainLength(subword);
      max_subl = std::max(subl, max_subl);
    }
    clen.insert({word, 1 + max_subl});
    return 1 + max_subl;
  }

  int longestStrChain(std::vector<std::string>& words) {
    mwords.clear();
    for (int i = 0; i < words.size(); i++) {
      mwords.insert({words[i], i});
    }
    clen.clear();

    int ans = 0;
    for (int i = 0; i < words.size(); i++) {
      int l = getChainLength(words[i]);
      ans = std::max(ans, l);
    }
    // Note: each entry in words is guaranteed to be between 1 and 16 characters long.
    //       There will be at most 1000 entries.
    return ans;
  }

private:
  std::unordered_map<std::string, int> mwords; // maps a word to its index in original words vector.

  // memoization: for each word in words, store its chain length.
  std::unordered_map<std::string, int> clen;
};

int main() {
  Solution s;
  std::vector<std::string> example1 = {"a", "b", "ba", "bca", "bda", "bdca"};
  assert(s.longestStrChain(example1) == 4);

  std::vector<std::string> example2 = {"xbc", "pcxbcf", "xb", "cxbc", "pcxbc"};
  assert(s.longestStrChain(example2) == 5);

  std::vector<std::string> example3 = {"abcd", "dbqca"};
  assert(s.longestStrChain(example3) == 1);
}
