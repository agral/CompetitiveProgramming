import unittest

from lc3005 import Solution

class Test_maxFrequencyElements(unittest.TestCase):
    def test(self):
        s = Solution()
        self.assertEqual(s.maxFrequencyElements([1, 2, 2, 3, 1, 4]), 4)
        self.assertEqual(s.maxFrequencyElements([1, 2, 3, 4, 5]), 5)

if __name__ == "__main__":
    unittest.main()
