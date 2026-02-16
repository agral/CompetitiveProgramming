# Runtime complexity: O(n)
# Auxiliary space complexity: O(2*n) == O(n)
# Subjective level: easy
# Solved on: 2026-02-16
class Solution:
    def rotateString(self, s: str, goal: str) -> bool:
        # Answer uses the same trick as 0459/Repeated Substring Pattern.
        return len(s) == len(goal) and (s+s).find(goal) != -1
