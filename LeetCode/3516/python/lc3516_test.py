import unittest

from lc3516 import Solution

class Test_findClosest(unittest.TestCase):
    def test(self):
        s = Solution()
        self.assertEqual(s.findClosest(2, 7, 4), 1)
        self.assertEqual(s.findClosest(2, 5, 6), 2)
        self.assertEqual(s.findClosest(1, 5, 3), 0)

if __name__ == "__main__":
    unittest.main()
