from typing import List

class Solution:
    def constructDistancedSequence(self, n: int) -> List[int]:
        ans = [0] * (2 * n - 1)

        def dfs(k: int, mask: int) -> bool:
            if k == len(ans):
                return True
            if ans[k] > 0:
                return dfs(k+1, mask)
            for num in range(n, 0, -1):
                if (mask >> num & 1) == 1:
                    continue
                if num == 1:
                    ans[k] = num
                    if dfs(k+1, mask | 1 << num):
                        return True
                    ans[k] = 0
                else: # so num in range {2, n}
                    if num + k >= len(ans) or ans[num + k] > 0:
                        continue
                    ans[k] = num
                    ans[k + num] = num
                    if dfs(k+1, mask | 1 << num):
                        return True
                    ans[k + num] = 0
                    ans[k] = 0
            return False

        dfs(0, 0)
        return ans


def main():
    s = Solution()
    assert(s.constructDistancedSequence(3) == [3, 1, 2, 3, 2])
    assert(s.constructDistancedSequence(5) == [5, 3, 1, 4, 3, 5, 2, 4, 2])

if __name__ == "__main__":
    main()
