from typing import List
from collections import Counter

class Solution:
    def maximumSubarraySum(self, nums: List[int], k: int) -> int:
        c = Counter()
        sum = 0
        uniqueNums = 0
        ans = 0
        for n, num in enumerate(nums):
            # remove the first number from the subarray, add last number to subarray
            if n >= k:
                removed = nums[n - k]
                c[removed] -= 1
                if c[removed] == 0:
                    uniqueNums -= 1
                sum -= removed
            sum += num
            c[num] += 1
            if c[num] == 1:
                uniqueNums += 1
            # if all numbers in subarray are distinct and subarray is full (n >= k - 1 in 0-base),
            # consider this subarray for an answer:
            if uniqueNums == k and (n >= k - 1):
                ans = max(ans, sum)
        return ans

def main():
    s = Solution()
    assert(s.maximumSubarraySum([1, 5, 4, 2, 9, 9, 9], 3) == 15)
    assert(s.maximumSubarraySum([4, 4, 4], 3) == 0)

if __name__ == "__main__":
    main()
