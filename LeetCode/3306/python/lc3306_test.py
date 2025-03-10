import unittest

from lc3306 import Solution

class Test_countOfSubstrings(unittest.TestCase):
    def test(self):
        s = Solution()
        self.assertEqual(s.countOfSubstrings("aeioqq", 1), 0)
        self.assertEqual(s.countOfSubstrings("aeiou", 0), 1)
        self.assertEqual(s.countOfSubstrings("ieaouqqieaouqq", 1), 3)

if __name__ == "__main__":
    unittest.main()
