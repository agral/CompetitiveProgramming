from typing import List

class Solution:
    def lexicographicallySmallestArray(self, nums: List[int], limit: int) -> List[int]:
        ans = [0] * len(nums)
        indexedSorted = sorted([(entry, num) for num, entry in enumerate(nums)])

        # group indexedSorted into groups where entries can be freely exchanged.
        groups = []
        for isorted in indexedSorted:
            if len(groups) == 0 or isorted[0] - groups[-1][-1][0] > limit:
                groups.append([isorted])
            else:
                groups[-1].append(isorted)

        for group in groups:
            # sorted is a copy of group values but without original indices.
            # maybe there's a better way to do this...?
            entries = [entry for entry, num in group]
            nums = sorted([num for entry, num in group])
            for entry, num in zip(entries, nums):
                ans[num] = entry
        return ans

def main():
    s = Solution()
    assert(s.lexicographicallySmallestArray([1, 5, 3, 9, 8], 2) == [1, 3, 5, 8, 9])
    assert(s.lexicographicallySmallestArray([1, 7, 6, 18, 2, 1], 3) == [1, 6, 7, 18, 1, 2])
    assert(s.lexicographicallySmallestArray([1, 7, 28, 19, 10], 3) == [1, 7, 28, 19, 10])
    print("All testcases passed successfully")

if __name__ == "__main__":
    main()
