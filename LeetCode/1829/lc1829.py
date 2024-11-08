from typing import List
class Solution:
    def getMaximumXor(self, nums: List[int], maximumBit: int) -> List[int]:
        max_k = (1 << maximumBit) - 1
        xor = 0
        ans = []
        for n in nums:
            xor ^= n
            ans.append(xor ^ max_k)
        return ans[::-1]

def main():
    s = Solution();
    assert(s.getMaximumXor([0, 1, 1, 3], 2) == [0, 3, 2, 3])
    assert(s.getMaximumXor([2, 3, 4, 7], 3) == [5, 2, 6, 5])
    assert(s.getMaximumXor([0, 1, 2, 2, 5, 7], 3) == [4, 3, 6, 4, 6, 7])
    print("All test cases passed.")

if __name__ == "__main__":
    main()

