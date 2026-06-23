import unittest

from lc3699 import Solution

class Test_zigZagArrays(unittest.TestCase):
    def test(self):
        s = Solution()
        self.assertEqual(s.zigZagArrays(3, 4, 5), 2)
        self.assertEqual(s.zigZagArrays(3, 1, 3), 10)

if __name__ == "__main__":
    unittest.main()
