import unittest

from lc3003 import Solution

class Test_maxPartitionsAfterOperations(unittest.TestCase):
    def test(self):
        s = Solution()
        self.assertEqual(s.maxPartitionsAfterOperations("accca", 2), 3)
        self.assertEqual(s.maxPartitionsAfterOperations("aabaab", 3), 1)
        self.assertEqual(s.maxPartitionsAfterOperations("xxyz", 1), 4)

if __name__ == "__main__":
    unittest.main()
