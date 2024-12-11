from typing import List

class Solution:
    def maximumBeauty(self, nums: List[int], k: int) -> int:
        nums.sort()
        ans = 0
        left = 0
        double_k = 2 * k
        # try with every possible right index. This is O(~2n) = O(n).
        for right in range(len(nums)):
            # increase left index until it's in valid range:
            # note: with this approach left can be >= right; we don't care as long as it produces valid answers.
            while nums[right] - nums[left] > double_k:
                left += 1
            ans = max(ans, right - left + 1)
        return ans

def main():
    s = Solution()
    assert(s.maximumBeauty([4, 6, 1, 2], 2) == 3)
    assert(s.maximumBeauty([1, 1, 1, 1], 10) == 4)
    print("All tests have passed.")

if __name__ == "__main__":
    main()
