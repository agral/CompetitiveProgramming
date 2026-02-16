import unittest

from lc0686 import Solution

class Test_repeatedStringMatch(unittest.TestCase):
    def test(self):
        s = Solution()
        self.assertEqual(s.repeatedStringMatch("abcd", "cdabcdab"), 3) # Example 1
        self.assertEqual(s.repeatedStringMatch("a", "aa"), 2) # Example 2

if __name__ == "__main__":
    unittest.main()
