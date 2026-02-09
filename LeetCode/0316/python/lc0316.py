from typing import List
from collections import Counter

# Runtime complexity: O(n)
# Auxiliary space complexity: O(n)
# Subjective level: medium
# Solved on: 2026-02-09
class Solution:
    def removeDuplicateLetters(self, s: str) -> str:
        NUM_LETTERS = ord('z') - ord('a') + 1
        is_used = [False] * NUM_LETTERS
        char_count = Counter(s)
        ans = []

        for ch in s:
            char_count[ch] -= 1
            if not is_used[ord(ch) - ord('a')]:
                while ans and ans[-1] > ch and char_count[ans[-1]] > 0:
                    is_used[ord(ans[-1]) - ord('a')] = False
                    ans.pop()
                ans.append(ch)
                is_used[ord(ans[-1]) - ord('a')] = True

        return "".join(ans)
