import unittest

from lc1200 import Solution

class Test_minimumAbsDifference(unittest.TestCase):
    def test(self):
        s = Solution()
        self.assertEqual(s.minimumAbsDifference([4, 2, 1, 3]), [[1, 2], [2, 3], [3, 4]])
        self.assertEqual(s.minimumAbsDifference([1, 3, 6, 10, 15]), [[1, 3]])
        self.assertEqual(s.minimumAbsDifference([3, 8, -10, 23, 19, -4, -14, 27]),
                         [[-14, -10], [19, 23], [23, 27]])

if __name__ == "__main__":
    unittest.main()
