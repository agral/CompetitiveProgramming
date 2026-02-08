import unittest

from lc0448 import Solution

class Test_findDisappearedNumbers(unittest.TestCase):
    def test(self):
        s = Solution()
        self.assertEqual(s.findDisappearedNumbers([4, 3, 2, 7, 8, 2, 3, 1]), [5, 6])
        self.assertEqual(s.findDisappearedNumbers([1, 1]), [2])

if __name__ == "__main__":
    unittest.main()
