import unittest

from lc3507 import Solution

class Test_isNonDecreasing(unittest.TestCase):
    def test(self):
        s = Solution()
        self.assertFalse(s.isNonDecreasing([5, 2, 3, 1]))
        self.assertTrue(s.isNonDecreasing([1, 2, 2]))

class Test_minimumPairRemoval(unittest.TestCase):
    def test(self):
        s = Solution()
        #self.assertEqual(s.minimumPairRemoval([5, 2, 3, 1]), 2)
        self.assertEqual(s.minimumPairRemoval([1, 2, 2]), 0)

if __name__ == "__main__":
    unittest.main()
