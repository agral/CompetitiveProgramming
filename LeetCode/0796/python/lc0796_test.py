import unittest

from lc0796 import Solution

class Test_rotateString(unittest.TestCase):
    def test(self):
        s = Solution()
        self.assertEqual(s.rotateString("abcde", "cdeab"), True)
        self.assertEqual(s.rotateString("abcde", "abced"), False)
        self.assertEqual(s.rotateString("aa", "a"), False) # WA #01

if __name__ == "__main__":
    unittest.main()
