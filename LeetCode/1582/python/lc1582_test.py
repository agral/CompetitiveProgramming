import unittest

from lc1582 import Solution

class Test_numSpecial(unittest.TestCase):
    def test(self):
        s = Solution()
        self.assertEqual(s.numSpecial([[1, 0, 0], [0, 0, 1], [1, 0, 0]]), 1)
        self.assertEqual(s.numSpecial([[1, 0, 0], [0, 1, 0], [0, 0, 1]]), 3)

if __name__ == "__main__":
    unittest.main()
