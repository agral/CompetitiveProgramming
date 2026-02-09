import unittest

from lc1700 import Solution

class Test_countStudents(unittest.TestCase):
    def test(self):
        s = Solution()
        self.assertEqual(s.countStudents([1, 1, 0, 0], [0, 1, 0, 1]), 0)
        self.assertEqual(s.countStudents([1, 1, 1, 0, 0, 1], [1, 0, 0, 0, 1, 1]), 3)

if __name__ == "__main__":
    unittest.main()
