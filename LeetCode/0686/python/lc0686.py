import math

# Runtime complexity: O(len(a) + len(b))
# Auxiliary space complexity: O(len(a) + len(b))
# Subjective level: medium
# Solved on: 2026-02-16
class Solution:
    def repeatedStringMatch(self, a: str, b: str) -> int:
        minReps = math.ceil(len(b) / len(a))
        repeated = a * minReps
        if repeated.find(b) != -1:
            return minReps
        repeated += a
        if repeated.find(b) != -1:
            return minReps + 1
        return -1
