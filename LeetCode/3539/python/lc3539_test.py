import unittest

from lc3539 import Solution

class Test_magicalSum(unittest.TestCase):
    def test(self):
        s = Solution()
        self.assertEqual(s.magicalSum(5, 5, [1, 10, 100, 10_000, 1_000_000]), 991600007)
        self.assertEqual(s.magicalSum(5, 5, [1, 10, 100, 10_000, 100_000]), 999160007)
        self.assertEqual(s.magicalSum(2, 2, [5, 4, 3, 2, 1]), 170)
        self.assertEqual(s.magicalSum(1, 1, [28]), 28)

if __name__ == "__main__":
    unittest.main()
