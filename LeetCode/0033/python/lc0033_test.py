import unittest

from lc0033 import Solution

class Test_search(unittest.TestCase):
    def test(self):
        s = Solution()
        self.assertEqual(s.search([4, 5, 6, 7, 0, 1, 2], 0), 4)
        self.assertEqual(s.search([4, 5, 6, 7, 0, 1, 2], 3), -1)
        self.assertEqual(s.search([1], 0), -1)

if __name__ == "__main__":
    unittest.main()
