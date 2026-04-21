import unittest

from lc1722 import Solution

class Test_minimumHammingDistance(unittest.TestCase):
    def test(self):
        s = Solution()
        self.assertEqual(s.minimumHammingDistance([1, 2, 3, 4], [2, 1, 4, 5], [[0, 1], [2, 3]]), 1)
        self.assertEqual(s.minimumHammingDistance([1, 2, 3, 4], [1, 3, 2, 4], []), 2)
        self.assertEqual(s.minimumHammingDistance([5, 1, 2, 4, 3], [1, 5, 4, 2, 3],
                                                  [[0, 4], [4, 2], [1, 3], [1, 4]]), 0)

if __name__ == "__main__":
    unittest.main()
