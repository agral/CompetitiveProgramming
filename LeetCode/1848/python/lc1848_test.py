import unittest

from lc1848 import Solution

class Test_getMinDistance(unittest.TestCase):
    def test(self):
        s = Solution()
        self.assertEqual(s.getMinDistance([1, 2, 3, 4, 5], 5, 3), 1)
        self.assertEqual(s.getMinDistance([1], 1, 0), 0)
        self.assertEqual(s.getMinDistance([1, 1, 1, 1, 1, 1, 1, 1, 1], 1, 0), 0)

if __name__ == "__main__":
    unittest.main()
