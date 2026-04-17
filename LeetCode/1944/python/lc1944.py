from typing import List

# Runtime complexity: O(n)
# Auxiliary space complexity: O(n)
# Subjective level: medium+
# Solved on: 2026-04-17
class Solution:
    def canSeePersonsCount(self, heights: List[int]) -> List[int]:
        L = len(heights)
        ans = [0] * L # initially treat every person as seeing no other.
        stack = [] # keep a stack of so-far seen heights.

        for idx, height in enumerate(heights):
            while stack and heights[stack[-1]] <= height:
                h = stack.pop()
                ans[h] += 1
            if stack: # a.k.a. if stack is non-empty at this point.
                # The person on the top of the stack can see this one.
                ans[stack[-1]] += 1
            stack.append(idx) # finally, add the current person to the stack.

        return ans
