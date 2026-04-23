import unittest

from lc2615 import Solution

class Test_distance(unittest.TestCase):
    def test(self):
        s = Solution()
        self.assertEqual(s.distance([1, 3, 1, 1, 2]), [5, 0, 3, 4, 0])
        self.assertEqual(s.distance([0, 5, 3]), [0, 0, 0])


if __name__ == "__main__":
    unittest.main()
