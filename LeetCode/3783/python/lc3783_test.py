import unittest

from lc3783 import Solution

class Test_reverse(unittest.TestCase):
    def test(self):
        s = Solution()
        self.assertEqual(s.reverse(12345), 54321)
        self.assertEqual(s.reverse(12321), 12321)
        self.assertEqual(s.reverse(10000), 1)
        self.assertEqual(s.reverse(1), 1)
        self.assertEqual(s.reverse(3001), 1003)

class Test_mirrorDistance(unittest.TestCase):
    def test(self):
        s = Solution()
        self.assertEqual(s.mirrorDistance(25), 27)
        self.assertEqual(s.mirrorDistance(10), 9)
        self.assertEqual(s.mirrorDistance(7), 0)

if __name__ == "__main__":
    unittest.main()
