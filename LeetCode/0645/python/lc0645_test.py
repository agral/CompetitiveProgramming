import unittest

from lc0645 import Solution

class Test_findErrorNums(unittest.TestCase):
    def test(self):
        s = Solution()
        self.assertEqual(s.findErrorNums([1, 2, 2, 4]), [2, 3])
        self.assertEqual(s.findErrorNums([1, 1]), [1, 2])

if __name__ == "__main__":
    unittest.main()
