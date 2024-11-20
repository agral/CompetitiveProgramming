from collections import Counter

class Solution:
    def takeCharacters(self, s: str, k: int) -> int:
        cnt = Counter(s)
        SZ = len(s)
        if (any(cnt[char] < k for char in "abc")):
            return -1
        ans = SZ

        left = 0
        for right, char in enumerate(s):
            cnt[char] -= 1
            while cnt[char] < k:
                cnt[s[left]] += 1
                left += 1
            ans = min(ans, SZ - 1 - (right - left))

        return ans

def main():
    s = Solution()
    assert(s.takeCharacters("aabaaaacaabc", 2) == 8)
    assert(s.takeCharacters("a", 1) == -1)
