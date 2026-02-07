import unittest

from lc0485 import Solution

class Test_findMaxConsecutiveOnes(unittest.TestCase):
    def test(self):
        s = Solution()
        self.assertEqual(s.findMaxConsecutiveOnes([1, 1, 0, 1, 1, 1]), 3)
        self.assertEqual(s.findMaxConsecutiveOnes([1, 0, 1, 1, 0, 1]), 2)
        self.assertEqual(s.findMaxConsecutiveOnes([0, 0, 0, 0, 0, 0]), 0)
        self.assertEqual(s.findMaxConsecutiveOnes([1, 1, 1, 1, 1, 1]), 6)

if __name__ == "__main__":
    unittest.main()
