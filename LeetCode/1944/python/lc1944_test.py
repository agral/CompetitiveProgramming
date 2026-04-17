import unittest

from lc1944 import Solution

class Test_canSeePersonsCount(unittest.TestCase):
    def test(self):
        s = Solution()
        self.assertEqual(s.canSeePersonsCount([10, 6, 8, 5, 11, 9]), [3, 1, 2, 1, 1, 0])
        self.assertEqual(s.canSeePersonsCount([5, 1, 2, 3, 10]), [4, 1, 1, 1, 0])

if __name__ == "__main__":
    unittest.main()
