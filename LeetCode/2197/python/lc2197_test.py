import unittest

from lc2197 import Solution

class Test_replaceNonCoprimes(unittest.TestCase):
    def test(self):
        s = Solution()
        self.assertEqual(s.replaceNonCoprimes([6, 4, 3, 2, 7, 6, 2]), [12, 7, 6])
        self.assertEqual(s.replaceNonCoprimes([2, 2, 1, 1, 3, 3, 3]), [2, 1, 1, 3])

if __name__ == "__main__":
    unittest.main()
