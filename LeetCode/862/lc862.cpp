#include <cassert>
#include <deque>
#include <vector>

class Solution {
public:
    int shortestSubarray(std::vector<int>& nums, int k) {
        const int SZ = nums.size();
        std::deque<int> q;
        std::vector<long> prefix{0};
        int ans = SZ + 1;

        for (int i = 0; i < SZ; ++i) {
            prefix.push_back(prefix.back() + nums[i]);
        }

        for (int i = 0; i < SZ + 1; ++i) {
            while (!q.empty() && prefix[i] - prefix[q.front()] >= k) {
                ans = std::min(ans, i - q.front());
                q.pop_front();
            }
            while (!q.empty() && prefix[i] <= prefix[q.back()]) {
                q.pop_back();
            }
            q.push_back(i);
        }

        return ans <= SZ ? ans : -1;
    }
};

int main() {
    Solution s;
    {
        std::vector<int> nums = {1};
        assert(s.shortestSubarray(nums, 1) == 1);
    }
    {
        std::vector<int> nums = {1, 2};
        assert(s.shortestSubarray(nums, 4) == -1);
    }
    {
        std::vector<int> nums = {2, -1, 2};
        assert(s.shortestSubarray(nums, 3) == 3);
    }
}
