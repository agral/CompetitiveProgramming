from typing import List

# Runtime complexity: O(n)
# Auxiliary space complexity: O(n) (for holding both the answer array and the stack)
# Subjective level: easy+
# Solved on: 2026-02-09
class Solution:
    def dailyTemperatures(self, temperatures: List[int]) -> List[int]:
        stack = [] # Holds indices of temperatures that currently don't have a higher temp to the right (so far).
        ans = [0] * len(temperatures) # important to be 0, as this solves the case of last entry/entries effortlessly.

        for i in range(len(temperatures)):
            while stack and temperatures[i] > temperatures[stack[-1]]:
                solvedIdx = stack.pop()
                ans[solvedIdx] = i - solvedIdx
            stack.append(i)
        return ans
