import unittest

from lc1470 import Solution

class Test_shuffle(unittest.TestCase):
    def test(self):
        s = Solution()
        self.assertEqual(s.shuffle([2, 5, 1, 3, 4, 7], 3), [2, 3, 5, 4, 1, 7])
        self.assertEqual(s.shuffle([1, 2, 3, 4, 4, 3, 2, 1], 4), [1, 4, 2, 3, 3, 2, 4, 1])
        self.assertEqual(s.shuffle([1, 1, 2, 2], 2), [1, 2, 1, 2])

if __name__ == "__main__":
    unittest.main()
