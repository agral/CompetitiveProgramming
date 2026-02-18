import unittest

from lc0001 import Solution

class Test_twoSum(unittest.TestCase):
    def test(self):
        s = Solution()
        self.assertEqual(s.twoSum([2, 7, 11, 15], 9), [0, 1])
        self.assertEqual(s.twoSum([3, 2, 4], 6), [1, 2])
        self.assertEqual(s.twoSum([3, 3], 6), [0, 1])

if __name__ == "__main__":
    unittest.main()
