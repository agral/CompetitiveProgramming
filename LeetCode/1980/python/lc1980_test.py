import unittest

from lc1980 import Solution

class Test_findDifferentBinaryString(unittest.TestCase):
    def test(self):
        s = Solution()
        self.assertIn(s.findDifferentBinaryString(["01", "10"]), ["00", "11"])
        self.assertIn(s.findDifferentBinaryString(["00", "01"]), ["10", "11"])
        self.assertIn(s.findDifferentBinaryString(["111", "011", "001"]), ["000", "010", "100", "101", "110"])

if __name__ == "__main__":
    unittest.main()
