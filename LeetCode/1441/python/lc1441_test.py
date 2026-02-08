import unittest

from lc1441 import Solution

class Test_buildArray(unittest.TestCase):
    def test(self):
        s = Solution()
        self.assertEqual(s.buildArray([1, 3], 3), ["Push", "Push", "Pop", "Push"])
        self.assertEqual(s.buildArray([1, 2, 3], 3), ["Push", "Push", "Push"])
        self.assertEqual(s.buildArray([1, 2], 2), ["Push", "Push"])
        self.assertEqual(s.buildArray([4], 3), ["Push", "Pop", "Push", "Pop", "Push", "Pop", "Push"])

if __name__ == "__main__":
    unittest.main()
