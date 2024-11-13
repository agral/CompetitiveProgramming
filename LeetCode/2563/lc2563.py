from typing import List

class Solution:
    def countFairPairs(self, nums: List[int], lower: int, upper: int) -> int:
        def countPairsLessThan(target_sum: int) -> int:
            ans = 0
            left = 0
            right = len(nums) - 1
            while (left < right):
                while left < right and nums[left] + nums[right] > target_sum:
                    right -= 1
                ans += right - left
                left += 1
            return ans

        nums.sort()
        return countPairsLessThan(upper) - countPairsLessThan(lower - 1)

def main():
    s = Solution()
    assert(s.countFairPairs([0, 1, 7, 4, 4, 5], 3, 6) == 6)
    # fair pairs: (0, 3), (0, 4), (0, 5), (1, 3), (1, 4), (1,5); 6 in total.

    assert(s.countFairPairs([1, 7, 9, 2, 5], 11, 11) == 1)
    # fair pairs: (2, 3); 1 in total.

if __name__ == "__main__":
    main()
