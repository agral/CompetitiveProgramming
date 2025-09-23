import unittest

from lc0165 import Solution

class Test_compareVersion(unittest.TestCase):
    def test(self):
        s = Solution()
        self.assertEqual(s.compareVersion("1.2", "1.10"), -1) # 1.10 is greater than 1.2
        self.assertEqual(s.compareVersion("1.01", "1.00001"), 0) # version numbers are equal.
        self.assertEqual(s.compareVersion("1.0", "1.0.0.0"), 0) # version numbers are equal.
        self.assertEqual(s.compareVersion("1.2.3.4", "1.1.9.9"), 1) # 1.2 > 1.1.

if __name__ == "__main__":
    unittest.main()
