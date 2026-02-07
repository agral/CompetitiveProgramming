import unittest

from lc1653 import Solution

class Test_minimumDeletions(unittest.TestCase):
    def test(self):
        s = Solution()
        self.assertEqual(s.minimumDeletions("aababbab"), 2)
        self.assertEqual(s.minimumDeletions("bbaaaabb"), 2)

if __name__ == "__main__":
    unittest.main()
