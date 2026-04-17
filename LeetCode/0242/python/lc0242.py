from typing import List

# Runtime complexity: O(n)
# Auxiliary space complexity: O(n)
# Subjective level: easy
# Solved on: 2026-04-17
class Solution:
    def isAnagram(self, s: str, t: str) -> bool:
        NUM_LETTERS = ord('z') - ord('a') + 1
        freq = [0] * NUM_LETTERS
        L = len(s)
        if L != len(t): # maybe a premature optimization, but sticking with it.
            return False

        for i in range(L):
            freq[ord(s[i]) - ord('a')] += 1 # keep adding one for each letter in s,
            freq[ord(t[i]) - ord('a')] -= 1 # and removing one for each letter in t.

        return all(f == 0 for f in freq)
        # follow up: what if the inputs contain Unicode characters?
        # then the same algo works, just use a hashmap (dict).
        # For lowercase english letters this approach works and is fast, so sticking with this.
