from typing import List

class Solution:
    def maxCount(self, banned: List[int], n: int, maxSum: int) -> int:
        ans = 0
        asum = 0
        sbanned = set(banned)

        for i in range(1, n+1):
            if (not i in sbanned) and (asum + i <= maxSum):
                ans += 1
                asum += i
        return ans

def main():
    s = Solution()
    assert(s.maxCount([1, 5, 6], 5, 6) == 2)
    assert(s.maxCount([1, 2, 3, 4, 5, 6, 7], 8, 1) == 0)
    assert(s.maxCount([11], 7, 50) == 7)

if __name__ == "__main__":
    main()
