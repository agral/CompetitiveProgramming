#include <algorithm>
#include <cassert>
#include <string>
#include <unordered_set>
#include <vector>

const std::size_t NUM_CHARS = 'z' - 'a' + 1;

class Solution {
public:
  int minDeletions(std::string s) {
    int ans = 0;
    std::vector<int> freq(NUM_CHARS, 0);
    for (const auto& ch: s) {
      freq[ch - 'a']++;
    }

    // Sort the frequency array. This loses the data of which letter maps to which frequency,
    // but the only data that matters in this task is the sizes of individual groups of letters.
    std::sort(freq.begin(), freq.end(), std::greater<>());
    std::unordered_set<int> groups;

    for (int i = 0; i < NUM_CHARS; i++) {
      while (groups.find(freq[i]) != groups.end()) {
        freq[i]--;
        ans++;
      }
      if (freq[i] > 0) { // if given group still exists and has unique count, store it:
        groups.insert(freq[i]);
      }
    }

    return ans;
  }
};

int main() {
  Solution s;

  assert(s.minDeletions("aab") == 0);
  assert(s.minDeletions("aaabbbcc") == 2);
  assert(s.minDeletions("ceabaacb") == 2);

  // Test that it handles groups that get deleted to count of 0:
  assert(s.minDeletions("aaabbbcccdddeee") == 9);
}
