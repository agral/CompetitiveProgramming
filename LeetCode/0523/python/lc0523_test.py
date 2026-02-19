import unittest

from lc0523 import Solution

class Test_checkSubarraySum(unittest.TestCase):
    def test(self):
        s = Solution()
        self.assertEqual(s.checkSubarraySum([23, 2, 4, 6, 7], 6), True)
        self.assertEqual(s.checkSubarraySum([23, 2, 6, 4, 7], 6), True)
        self.assertEqual(s.checkSubarraySum([23, 2, 6, 4, 7], 13), False)

if __name__ == "__main__":
    unittest.main()
