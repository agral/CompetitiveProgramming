from typing import List

class Solution:
    def minimizedMaximum(self, n: int, quantities: List[int]) -> int:
        def getStoresCount(mid: int) -> int:
            cnt = 0
            for q in quantities:
                cnt += ((q-1) // mid) + 1
            return cnt

        left_bound = 1
        right_bound = max(quantities)
        while left_bound < right_bound:
            mid = (left_bound + right_bound) // 2
            stores = getStoresCount(mid)
            if stores <= n:
                right_bound = mid
            else:
                left_bound = mid + 1

        return left_bound

def main():
    s = Solution()
    assert(s.minimizedMaximum(6, [11, 6]) == 3)
    assert(s.minimizedMaximum(7, [15, 10, 10]) == 5)
    assert(s.minimizedMaximum(1, [100000]) == 100000)

if __name__ == "__main__":
    main()
