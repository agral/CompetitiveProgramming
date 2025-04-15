from typing import List

class FenwickTree:
    @staticmethod
    def lsb(num: int) -> int:
        return num & -num

    def __init__(self, n: int):
        self.sums = [0] * (n + 1)

    def insert(self, num: int, delta: int) -> None:
        while num < len(self.sums):
            self.sums[num] += delta
            num += FenwickTree.lsb(num)

    def get(self, num: int) -> int:
        ans = 0
        while num > 0:
            ans += self.sums[num]
            num -= FenwickTree.lsb(num)
        return ans


# Runtime complexity: O(n*logn)
# Auxiliary space complexity: O(n); where n is the length of the input array.
class Solution:
    def goodTriplets(self, nums1: List[int], nums2: List[int]) -> int:
        length = len(nums1)
        numToIndex = {num: i for i, num in enumerate(nums1)}
        arr = [numToIndex[num] for num in nums2]

        # smaller: the count of arr[i] < arr[j], for i, j such that 0 <= i < j
        # larger: the count of arr[i] > arr[j], for i, j such that i < j < length
        smaller, larger = [0] * length, [0] * length
        fentree1 = FenwickTree(length)
        fentree2 = FenwickTree(length)

        for idx, num in enumerate(arr):
            smaller[idx] = fentree1.get(num)
            fentree1.insert(num+1, 1)
        for idx, num in reversed(list(enumerate(arr))):
            larger[idx] = fentree2.get(length) - fentree2.get(num)
            fentree2.insert(num+1, 1)

        return sum(l * r for l, r in zip(smaller, larger))
