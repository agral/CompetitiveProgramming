import unittest

from lc1358 import Solution

class Test_numberOfSubstrings(unittest.TestCase):
    def test(self):
        s = Solution()
        self.assertEqual(s.numberOfSubstrings("abcabc"), 10)
        self.assertEqual(s.numberOfSubstrings("aaacb"), 3)
        self.assertEqual(s.numberOfSubstrings("abc"), 1)
        self.assertEqual(s.numberOfSubstrings("abb"), 0)

if __name__ == "__main__":
    unittest.main()
