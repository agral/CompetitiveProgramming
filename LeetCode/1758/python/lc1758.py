from typing import List

# Runtime complexity: O(n)
# Auxiliary space complexity: O(1)
# Subjective level: easy
# Solved on: 2026-03-05
class Solution:
    def minOperations(self, s: str) -> int:
        numFlipsZero = 0
        numFlipsOne = 0

        for i in range(len(s)):
            if i&1 == 0:
                if s[i] == '1':
                    numFlipsZero += 1
                else:
                    numFlipsOne += 1
            else:
                if s[i] == '0':
                    numFlipsZero += 1
                else:
                    numFlipsOne += 1
        return min(numFlipsZero, numFlipsOne)
