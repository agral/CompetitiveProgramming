import unittest

from lc0474 import Solution

class Test_findMaxForm(unittest.TestCase):
    def test(self):
        s = Solution()
        self.assertEqual(s.findMaxForm(["10", "0001", "111001", "1", "0"], 5, 3), 4)
        self.assertEqual(s.findMaxForm(["10", "0", "1"], 1, 1), 2)

if __name__ == "__main__":
    unittest.main()
