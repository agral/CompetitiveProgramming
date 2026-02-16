import unittest

from lc1668 import Solution

class Test_maxRepeating(unittest.TestCase):
    def test(self):
        s = Solution()
        self.assertEqual(s.maxRepeating("ababc", "ab"), 2)
        self.assertEqual(s.maxRepeating("ababc", "ba"), 1)
        self.assertEqual(s.maxRepeating("ababc", "ac"), 0)

if __name__ == "__main__":
    unittest.main()
