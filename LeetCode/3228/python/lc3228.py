# Runtime complexity: O(n)
# Auxiliary space complexity: O(1)
# Subjective level: easy
# Solved on: 2025-11-13
class Solution:
    def maxOperations(self, s: str) -> int:
        numOnes = 0
        n = len(s)
        ans = 0
        for i in range(n):
            if s[i] == '1':
                numOnes += 1
            elif (i == n - 1) or s[i + 1] == '1':
                ans += numOnes
        return ans
