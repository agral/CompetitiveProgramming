from typing import List

# Runtime complexity: O(n)
# Auxiliary space complexity: O(n) (for the stack)
# Subjective level: easy
# Solved on: 2025-11-10
class Solution:
    def minOperations(self, nums: List[int]) -> int:
        stack = [0] # a standard LIFO stack.
        ans = 0
        for num in nums:
            while stack and stack[-1] > num:
                stack.pop()
            if not stack or stack[-1] < num:
                stack.append(num)
                ans += 1
        return ans

