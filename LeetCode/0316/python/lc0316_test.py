import unittest

from lc0316 import Solution

class Test_removeDuplicateLetters(unittest.TestCase):
    def test(self):
        s = Solution()
        self.assertEqual(s.removeDuplicateLetters("bcabc"), "abc")
        self.assertEqual(s.removeDuplicateLetters("cbacdcbc"), "acdb")

if __name__ == "__main__":
    unittest.main()
