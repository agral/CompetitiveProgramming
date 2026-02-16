# Runtime complexity: O(n^2)
# Auxiliary space complexity: O(n)
# Subjective level: easy+ (but the problem requires brute force, and is not that interesting).
# Solved on: 2026-02-16
class Solution:
    def maxRepeating(self, sequence: str, word: str) -> int:
        ans = 0
        while (word*(ans+1)) in sequence:
            ans += 1
        return ans
