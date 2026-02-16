import unittest

from lc0459 import Solution

class Test_repeatedSubstringPattern(unittest.TestCase):
    def test(self):
        s = Solution()
        self.assertEqual(s.repeatedSubstringPattern("abab"), True)
        self.assertEqual(s.repeatedSubstringPattern("aba"), False)
        self.assertEqual(s.repeatedSubstringPattern("abcabcabcabc"), True)

if __name__ == "__main__":
    unittest.main()
