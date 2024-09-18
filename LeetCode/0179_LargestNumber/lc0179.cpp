#include <algorithm>
#include <cassert>
#include <sstream>
#include <string>
#include <vector>

class Solution {
public:
    std::string largestNumber(std::vector<int>& nums) {
        // convert vector of nums to vector of strings:
        std::vector<std::string> strnums;
        // but also handle the case of all zeroes in input array:
        int nums_ORed = 0; // faster than setting & clearing a bool flag, OR all the nums
        strnums.reserve(nums.size());
        for (const int num: nums) {
            strnums.push_back(std::to_string(num));
            nums_ORed |= num;
        }

        if (nums_ORed == 0) {
            return "0";
        }

        // sort the numbers in reverse alphabetical order:
        std::sort(strnums.begin(), strnums.end(), [](const std::string& lhs, const std::string& rhs) {
            return lhs + rhs > rhs + lhs;
        });

        std::stringstream ss;
        for (const std::string& strnum: strnums) {
            ss << strnum;
        }
        return ss.str();
    }
};

int main() {
    Solution s;
    {
        std::vector<int> nums = {10, 2};
        assert(s.largestNumber(nums) == "210");
    }
    {
        std::vector<int> nums = {3, 30, 34, 5, 9};
        assert(s.largestNumber(nums) == "9534330");
    }
    {
        std::vector<int> nums = {0, 0, 0, 0, 0, 0};
        assert(s.largestNumber(nums) == "0");
    }
}
