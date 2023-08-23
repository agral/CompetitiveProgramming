#include <cassert>
#include <map>
#include <queue>
#include <sstream>
#include <string>
#include <utility>

const std::pair<int, char> NULL_PAIR = {0, 0};

class Solution {
public:
  std::string reorganizeString(std::string s) {
    // Create a frequency map:
    std::map<char, int> frequency;
    for (const auto& ch: s) {
      frequency[ch] += 1;
    }

    // Create a priority queue with characters and their remaining counts:
    std::priority_queue<std::pair<int, char>> queue;
    for (const auto& [character, count] : frequency) {
      queue.push({count, character});
    }

    // Create a reorganized string char by char using a stringstream:
    std::stringstream ans;
    std::pair<int, char> last_char = NULL_PAIR;
    while (!queue.empty()) {
      // take out the largest char set from the queue:
      auto [current_char_count, current_char] = queue.top();
      queue.pop();

      // add this largest char to the answer and decrease its total count by one:
      ans << current_char;
      current_char_count -= 1;

      // if available, put last_char back into the queue:
      if (last_char != NULL_PAIR) {
        queue.push(last_char);
        last_char = NULL_PAIR;
      }

      // unless all the letters are used up, store current pair as last_char.
      if (current_char_count > 0) {
        last_char = std::make_pair(current_char_count, current_char);
      }
    }
    if (last_char != NULL_PAIR) {
      return {""};
    }

    return ans.str();
  }
};

int main() {
  Solution s{};

  // Note: this implementation resolves ties by largest ascii code
  // (so in case of "bbcc" the answer will be "cbcb" even though "bcbc" is valid as well.
  // It would take extra time to handle resolving all valid cases, so just be aware of it.
  // It's good enough for competitive programming :-)
  assert(s.reorganizeString("aab") == "aba");
  assert(s.reorganizeString("aaab") == "");
  assert(s.reorganizeString("aaaaaabbbcc") == "abacabacaba");
  assert(s.reorganizeString("bfrbs") == "bsrfb");
}
