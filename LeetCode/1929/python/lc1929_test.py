import unittest

from lc1929 import Solution

class Test_getConcatenation(unittest.TestCase):
    def test(self):
        s = Solution()
        self.assertEqual(s.getConcatenation([1, 2, 1]), [1, 2, 1, 1, 2, 1])
        self.assertEqual(s.getConcatenation([1, 3, 2, 1]), [1, 3, 2, 1, 1, 3, 2, 1])


if __name__ == "__main__":
    unittest.main()
