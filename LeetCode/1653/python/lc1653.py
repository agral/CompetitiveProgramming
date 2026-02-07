# Runtime complexity: O(n)
# Auxiliary space complexity: O(1)
# Subjective level: medium-
# Solved on: 2026-02-07
class Solution:
    def minimumDeletions(self, s: str) -> int:
        # holds the total count of characters to be deleted so that s[:i] is balanced.
        numDeletions = 0
        numBChars = 0
        for ch in s:
            if (ch == 'b'):
                numBChars += 1
            else:
                # Either delete this 'a' or delete all the 'b's seen so far.
                numDeletions = min(numDeletions+1, numBChars)
        return numDeletions
