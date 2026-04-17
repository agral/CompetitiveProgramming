import unittest

from lc3761 import Solution

class Test_reverse(unittest.TestCase):
    def test(self):
        s = Solution()
        self.assertEqual(s.reverse(12345), 54321)
        self.assertEqual(s.reverse(10000), 1)
        self.assertEqual(s.reverse(10001), 10001)

class Test_minMirrorPairDistance(unittest.TestCase):
    def test(self):
        s = Solution()
        self.assertEqual(s.minMirrorPairDistance([12, 21, 45, 33, 54]), 1)
        self.assertEqual(s.minMirrorPairDistance([120, 21]), 1)
        self.assertEqual(s.minMirrorPairDistance([21, 120]), -1)
        self.assertEqual(s.minMirrorPairDistance([12, 34, 46, 21, 12]), 1)


if __name__ == "__main__":
    unittest.main()
