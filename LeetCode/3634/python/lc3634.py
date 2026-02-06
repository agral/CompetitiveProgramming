from typing import List

# Runtime complexity: O(sort) + O(n^2) (can be O(n*log2(n) with binary search)
# Auxiliary space complexity: O(sort)
# Subjective level: medium (an easy greedy approach looks feasible, but is actually not!)
# Solved on: 2026-02-06
class Solution:
    def minRemoval(self, nums: List[int], k: int) -> int:
        # Handle a trivial corner case where no sorting is necessary:
        if len(nums) == 1:
            return 0

        nums.sort()
        maxLength = 1
        for b in range(len(nums)):
            for e in range(b+maxLength, len(nums)):
                if k * nums[b] < nums[e]:
                    break
                maxLength = max(maxLength, e-b+1)
        return len(nums) - maxLength

    # Runtime complexity: O(sort) + O(n)
    # Auxiliary space complexity: O(sort)
    # Note: greedy approach will not work on all test cases. Keeping this for posterity.
    def minRemovalGreedy(self, nums: List[int], k: int) -> int:
        nums.sort()
        b, e = 0, len(nums)-1
        cntRemoved = 0 # 
        print(nums)
        while k*nums[b] < nums[e]:
            ratioB, ratioE = nums[b+1]/nums[b], nums[e]/nums[e-1]
            print(ratioB, ratioE)
            cntRemoved += 1
            if ratioE > ratioB:
                e -= 1
            else:
                b += 1
        return cntRemoved
