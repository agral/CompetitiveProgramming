class Solution:
    def numberOfSubstrings(self, s: str) -> int:
        lastSeen = {"a": -1, "b": -1, "c": -1}
        ans = 0
        for i, ch in enumerate(s):
            lastSeen[ch] = i
            ans += (min(lastSeen.values()) + 1) #+1 since the strings are zero-indexed.
        return ans
