import unittest

from lc1935 import Solution

class Test_canBeTypedWords(unittest.TestCase):
    def test(self):
        s = Solution()
        self.assertEqual(s.canBeTypedWords("hello world", "ad"), 1)
        self.assertEqual(s.canBeTypedWords("leet code", "lt"), 1)
        self.assertEqual(s.canBeTypedWords("leet code", "e"), 0)

if __name__ == "__main__":
    unittest.main()
