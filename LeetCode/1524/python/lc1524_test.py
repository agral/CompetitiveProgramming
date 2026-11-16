import unittest

from lc1524 import Solution

class Test_numOfSubarrays(unittest.TestCase):
    def test(self):
        s = Solution()
        self.assertEqual(s.numOfSubarrays([1, 3, 5]), 4)
        self.assertEqual(s.numOfSubarrays([2, 4, 6]), 0)
        self.assertEqual(s.numOfSubarrays([1, 2, 3, 4, 5, 6, 7]), 16)

if __name__ == "__main__":
    unittest.main()
