#include <cassert>
#include <string>
#include <unordered_map>

// Note: s is guaranteed to contain only lowercase English letters;
// so no need to take care of uppercase vowels.
const std::string VOWELS{"aeiou"};

class Solution {
public:
    int findTheLongestSubstring(std::string s) {
        int ans = 0;
        int binary_prefix = 0;
        std::unordered_map<int, int> prefix_to_index{};
        prefix_to_index[0] = -1;

        for (int i = 0; i < s.size(); ++i) {
            int index = VOWELS.find(s[i]);
            if (index > -1) {
                binary_prefix ^= (1 << index);
            }
            // Store the very first appearance of given prefix in the map;
            // do not update the map on subsequent occurrences of this prefix:
            if (!prefix_to_index.contains(binary_prefix)) {
                prefix_to_index[binary_prefix] = i;
            }
            ans = std::max(ans, i - prefix_to_index[binary_prefix]);
        }
        return ans;
    }
};

int main() {
    Solution s;
    assert(s.findTheLongestSubstring("eleetminicoworoep") == 13); // "leetminicowor"
    assert(s.findTheLongestSubstring("leetcodeisgreat") == 5); // "leetc"
    assert(s.findTheLongestSubstring("bcbcbc") == 6); // "bcbcbc"
}
