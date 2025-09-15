from typing import List

# Runtime complexity: O(len(words) * len(brokenLetters))
# Auxiliary space complexity: O(len(words))
# Subjective level: easy peasy.
class Solution:
    def canBeTypedWords(self, text: str, brokenLetters: str) -> int:
        ans = 0
        words = text.split(" ")
        for word in words:
            if not any(ch in word for ch in brokenLetters):
                ans += 1 # this word can be typed on the broken keyboard.

        return ans
