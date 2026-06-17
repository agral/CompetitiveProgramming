import unittest

from lc3614 import Solution

class Test_processStr(unittest.TestCase):
    def test(self):
        s = Solution()
        self.assertEqual(s.processStr("a#b%*", 1), "a")
        self.assertEqual(s.processStr("cd%#*#", 3), "d")
        self.assertEqual(s.processStr("z*#", 0), ".")

if __name__ == "__main__":
    unittest.main()
