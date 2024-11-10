from typing import List
import collections

class Solution:
    def minimumSubarrayLength(self, nums: List[int], k: int) -> int:
        ans = len(nums) + 1
        window_or = 0 # holds the result of OR of all the numbers in current window
        count = collections.Counter()
        left = 0;
        for right, num in enumerate(nums):
            window_or = self.calc_or(window_or, num, count)
            while window_or >= k and left <= right:
                ans = min(ans, right - left + 1)
                window_or = self.undo_or(window_or, nums[left], count)
                left += 1
        return -1 if ans == len(nums) + 1 else ans

    def calc_or(self, window_or: int, num: int, count: dict[int, int]) -> int:
        for i in range(32):
            if num >> i & 1:
                count[i] += 1
                if count[i] == 1:
                    window_or += 1 << i
        return window_or

    def undo_or(self, window_or: int, num: int, count: dict[int, int]) -> int:
        for i in range(32):
            if num >> i & 1:
                count[i] -= 1
                if count[i] == 0:
                    window_or -= 1 << i
        return window_or

def main():
    s = Solution()
    assert(s.minimumSubarrayLength([1, 2, 3], 2) == 1)
    assert(s.minimumSubarrayLength([2, 1, 8], 11) == 3)
    assert(s.minimumSubarrayLength([1, 2], 0) == 1)
    assert(s.minimumSubarrayLength([1, 2, 32, 19], 51) == 2)

if __name__ == "__main__":
    main()

