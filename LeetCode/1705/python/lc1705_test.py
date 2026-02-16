import unittest

from lc1705 import Solution

class Test_eatenApples(unittest.TestCase):
    def test(self):
        s = Solution()
        self.assertEqual(s.eatenApples([1, 2, 3, 5, 2], [3, 2, 1, 4, 2]), 7)
        self.assertEqual(s.eatenApples([3, 0, 0, 0, 0, 2], [3, 0, 0, 0, 0, 2]), 5)

if __name__ == "__main__":
    unittest.main()
