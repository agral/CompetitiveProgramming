#include <algorithm>
#include <cassert>
#include <vector>

class Solution {
public:
    int longestSubarray(std::vector<int>& nums) {
        // We're looking for a maximum bitwise AND of any subarray of nums.
        // This means that we're looking for a max element of nums first:

        int max_value = *std::max_element(nums.begin(), nums.end());
        int count = 0, max_count = 0;
        for (int i = 0; i < nums.size(); i++) {
            if (nums[i] == max_value) {
                count += 1;
                max_count = std::max(max_count, count);
            } else {
                count = 0;
            }
        }
        return max_count;
    }
};

int main() {
    Solution s;
    {
        std::vector<int> nums {1, 2, 3, 3, 2, 2};
        assert(s.longestSubarray(nums) == 2);
    }
    {
        std::vector<int> nums {1, 2, 3, 4};
        assert(s.longestSubarray(nums) == 1);
    }
    {
        std::vector<int> nums {1, 1, 1, 2, 2, 1};
        assert(s.longestSubarray(nums) == 2);
    }
    {
        std::vector<int> nums {2, 2, 2, 1, 1, 2};
        assert(s.longestSubarray(nums) == 3);
    }
    {
        std::vector<int> nums {2, 2, 1, 1, 2, 2, 2};
        assert(s.longestSubarray(nums) == 3);
    }
}
