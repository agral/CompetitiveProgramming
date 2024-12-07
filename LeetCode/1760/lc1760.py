from typing import List
import bisect

class Solution:
    def minimumSize(self, nums: List[int], maxOperations: int) -> int:
        def get_num_operations(m: int) -> bool:
            return sum( (num-1) // m for num in nums ) <= maxOperations
        return bisect.bisect_left(range(1, max(nums)), True, key=lambda m: get_num_operations(m)) + 1

def main():
    s = Solution()
    assert(s.minimumSize([9], 2) == 3)
    assert(s.minimumSize([2, 4, 8, 2], 4) == 2)

if __name__ == "__main__":
    main()
