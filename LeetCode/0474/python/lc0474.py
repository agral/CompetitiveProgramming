from typing import List

# Runtime complexity: don't know, but lower bound is O(numZeroes*numOnes*len(strs)).
# Auxiliary space complexity: O(numZeroes*numOnes)
# Subjective level: hard.
# Solved on: 2025-11-11
class Solution:
    def findMaxForm(self, strs: List[str], numZeroes: int, numOnes: int) -> int:
        # dp[z][o] holds the maximum size of the subset with at most `z` zeroes and `o` ones.
        dp = [[0] * (numOnes+1) for _ in range(numZeroes+1)]
        for s in strs:
            zeroes = s.count("0")
            ones = len(s) - zeroes

            for z in range(numZeroes, zeroes-1, -1):
                for o in range(numOnes, ones-1, -1):
                    dp[z][o] = max(dp[z][o], dp[z-zeroes][o-ones]+1)
        return dp[numZeroes][numOnes]
